package routers

import (
	"net/http"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/maulidihsan/rujuk/pkg/model"
	"github.com/maulidihsan/rujuk/pkg/api/v1"
)

func (s *Service) transformDataRS(in *v1.Rumahsakit) model.Rumahsakit {
	out := model.Rumahsakit{
		ID: uint(in.Id),
		NamaRumahsakit: in.Nama,
		IP: in.Ip,
	}
	return out
}

func (s *Service) transformDataRoom(in *v1.Room) model.Room {
	out := model.Room{
		ID: uint(in.Id),
		NamaRuangan: in.Nama,
		TipeRuangan: in.Type,
		Jumlah: in.Jumlah,
	}
	return out
}

type Service struct {
	db *gorm.DB
	r *grpc.ClientConn
	d *grpc.ClientConn
}

func Init(db *gorm.DB, discoveryUri string) (*Service, error) {
	fmt.Println(discoveryUri)
	if !db.HasTable(&model.Room{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Room{}).Error; err != nil {
			return nil, err
		}
	}
	if !db.HasTable(&model.Pasien{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Pasien{}).Error; err != nil {
			return nil, err
		}
	}
	discoveryConn, err := grpc.Dial(fmt.Sprintf("%s", discoveryUri), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Service{db: db, d: discoveryConn}, nil
}

func (s *Service) GetMyPasiens(c *gin.Context) {
	var pasiens []model.Pasien
	err := s.db.Model(&model.Pasien{}).Find(&pasiens).Error;
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pasiens})
}

func (s *Service) GetMyRooms(c *gin.Context) {
	var rooms []model.Room
	err := s.db.Model(&model.Room{}).Find(&rooms).Error;
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

func (s *Service) ListOtherRS(c *gin.Context) {
	var query model.Request
	c.Bind(&query)
	offset := 0
	limit := 1000

	client := v1.NewDiscoveryServiceClient(s.d)

	req := v1.FetchRequest{
		Offset: int32(offset),
		Limit: int32(limit),
	}

	data, err := client.GetAllRumahsakit(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	list := make([]model.Rumahsakit, len(data.GetRumahsakits()))
	for i, rs := range data.GetRumahsakits(){
		item := s.transformDataRS(rs)
		list[i] = item
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (s *Service) Register(c *gin.Context) {
	var rs model.Rumahsakit
	c.Bind(&rs)
	client := v1.NewDiscoveryServiceClient(s.d)
	req := v1.Rumahsakit{
		Nama: rs.NamaRumahsakit,
		Ip: rs.IP,
	}
	data, err := client.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"kode": data.GetKode(), "msg": data.GetPesan()})
}

func (s *Service) ConnectToRs(ip string) error {
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		return err
	}
	s.r = conn
	return nil
}

func (s *Service) GetRoomList(c *gin.Context) {
	var rs model.Rumahsakit
	c.Bind(&rs)
	if err := s.ConnectToRs(rs.IP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	offset := 0
	limit := 1000
	client := v1.NewRujukServiceClient(s.r)
	req := v1.FetchRequest{
		Offset: int32(offset),
		Limit: int32(limit),
	}

	data, err := client.GetAllRoom(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	list := make([]model.Room, len(data.GetRooms()))
	for i, room := range data.GetRooms(){
		item := s.transformDataRoom(room)
		list[i] = item
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (s *Service) TransferPasien(c *gin.Context) {
	var request model.TransferRujuk
	c.Bind(&request)
	if err := s.ConnectToRs(request.Ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	client := v1.NewRujukServiceClient(s.r)
	req := v1.RequestRujuk{
		Id: uint32(request.Id),
		Pasien: &v1.Pasien{
			Nama: request.Pasien.NamaPasien,
			Diagnosa: request.Pasien.Diagnosa,
		},
	}
	fmt.Printf("%+v\n", request)
	data, err := client.RequestRoom(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := s.db.Delete(&request.Pasien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"kode": data.GetKode(), "msg": data.GetPesan()})
}

func (s *Service) AddPasien(c *gin.Context) {
	var pasien model.Pasien
	c.Bind(&pasien)
	fmt.Println(&pasien)
	if err := s.db.Create(&pasien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (s *Service) AddRoom(c *gin.Context) {
	var room model.Room
	c.Bind(&room)
	fmt.Println(&room)
	if err := s.db.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}