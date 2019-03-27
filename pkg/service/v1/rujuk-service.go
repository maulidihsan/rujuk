package v1;

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/maulidihsan/rujuk/pkg/api/v1"
	"github.com/maulidihsan/rujuk/pkg/model"
)

type rujuk struct {
	db *gorm.DB
}

func NewRujukServiceServer(db *gorm.DB) v1.RujukServiceServer {
	if !db.HasTable(&model.Room{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Room{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&model.Pasien{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Pasien{}).Error; err != nil {
			panic(err)
		}
	}
	return &rujuk{db: db}
}

func (s *rujuk) transformRoomToRPC(in *model.Room) *v1.Room {
	if in == nil {
		return nil
	}
	res := &v1.Room{
		Id: uint32(in.ID),
		Nama: in.NamaRuangan,
		Type: in.TipeRuangan,
		Jumlah: in.Jumlah,
	}
	return res
}

func (s *rujuk) GetAllRoom(ctx context.Context, req *v1.FetchRequest) (*v1.Rooms, error) {
	offset := int32(0)
	limit := int32(10)
	if req != nil {
		offset = req.Offset
		limit = req.Limit
	}
	var rooms []model.Room
	err := s.db.Model(&model.Room{}).Offset(offset).Limit(limit).Find(&rooms).Error;
	if err != nil {
		return nil,err
	}
	payload := make([]*v1.Room, len(rooms))
	for i, room := range rooms {
		r := s.transformRoomToRPC(&room)
		payload[i] = r
	}
	reply := &v1.Rooms{
		Rooms: payload,
	}
	return reply, nil
}

func (s *rujuk) RequestRoom(ctx context.Context, req *v1.RequestRujuk) (*v1.Response, error) {
	if req == nil {
		//TODO: mekanisme return error kalau kosong
	}
	fmt.Println(req)
	var ruangan model.Room
	if err := s.db.Where("id = ?", uint32(req.Id)).First(&ruangan).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if (ruangan.Jumlah > 0) {
		tx := s.db.Begin()
		newPasien := model.Pasien{
			NamaPasien: req.Pasien.Nama,
			Diagnosa: req.Pasien.Diagnosa,
		}

		if err := tx.Create(&newPasien).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		ruangan.Jumlah -= 1;
		if err := tx.Save(&ruangan).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}

		return &v1.Response{
			Kode: 200,
			Pesan: "Ruangan berhasil dipesan",
		}, nil
	}
	return &v1.Response{
		Kode: 400,
		Pesan: "Ruangan sudah penuh",
	}, nil
}