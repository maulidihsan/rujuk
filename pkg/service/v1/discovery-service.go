package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/jinzhu/gorm"

	"github.com/maulidihsan/rujuk/pkg/api/v1"
	"github.com/maulidihsan/rujuk/pkg/model"
)

type service struct {
	db *gorm.DB
}

func NewDiscoveryServiceServer(db *gorm.DB) v1.DiscoveryServiceServer {
	if !db.HasTable(&model.Rumahsakit{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf-8").CreateTable(&model.Rumahsakit{}).Error; err != nil {
			panic(err)
		}
	}
	return &service{db: db}
}

func (s *service) transformRumahsakitToRPC(in *model.Rumahsakit) *v1.Rumahsakit {
	if in == nil {
		return nil
	}
	res := &v1.Rumahsakit{
		ID: in.ID,
		Nama: in.NamaRumahsakit,
		Ip: in.IP,
	}
	return res
}

func (s *service) GetAllRumahsakit(req *v1.FetchRequest, stream v1.RujukService_GetAllRoomServer) error {
	offset := 0
	limit := 10
	if req != nil {
		offset = req.Offset
		limit = req.Limit
	}
	var rumahsakit []model.Rumahsakit
	err := s.db.Model(&model.Rumahsakit{}).Offset(offset).Limit(limit).Find(&rumahsakit).Error;
	if err != nil {
		return err
	}
	for _, rs := range rumahsakit {
		r := s.transformRumahsakitToRPC(rs)
		if err := stream.Send(r); err != nil {
			log.Println(err.Error())
			return err
		}
	}
}

func (s *service) Register(req *v1.Rumahsakit) (*v1.Response, error) {
	if req == nil {
		//TODO: mekanisme return error kalau kosong
	}
	rumahsakit = model.Rumahsakit{
		NamaRumahsakit: req.Nama,
		IP: req.Ip
	}
	if err := s.db.Where(model.Rumahsakit{IP: req.Ip}).FirstOrCreate(&rumahsakit).Error; err != nil {
		return nil, err
	}

	return &v1.Response{
		Kode: 200,
		Pesan: "Registrasi berhasil"
	}, nil
}