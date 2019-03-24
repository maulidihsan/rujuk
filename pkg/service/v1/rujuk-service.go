package v1;

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

func NewRujukServiceServer(db *gorm.DB) v1.RujukServiceServer {
	if !db.HasTable(&model.Room{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf-8").CreateTable(&model.Room{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&model.Pasien{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT=utf-8").CreateTable(&model.Pasien{}).Error; err != nil {
			panic(err)
		}
	}
	return &service{db: db}
}

func (s *service) transformRoomToRPC(in *model.Room) *v1.Room {
	if in == nil {
		return nil
	}
	res := &v1.Room{
		ID: in.ID,
		Nama: in.NamaRuangan,
		Tipe: in.TipeRuangan,
		Jumlah: in.Jumlah,
	}
	return res
}

func (s *service) GetAllRoom(req *v1.FetchRequest, stream v1.RujukService_GetAllRoomServer) error {
	offset := 0
	limit := 10
	if req != nil {
		offset = req.Offset
		limit = req.Limit
	}
	var rooms []model.Room
	err := s.db.Model(&model.Room{}).Offset(offset).Limit(limit).Find(&rooms).Error;
	if err != nil {
		return err
	}
	for _, room := range rooms {
		r := s.transformRoomToRPC(room)
		if err := stream.Send(r); err != nil {
			log.Println(err.Error())
			return err
		}
	}
}

func (s *service) RequestRoom(req *v1.RequestRujuk) (*v1.Response, error) {
	if req == nil {
		//TODO: mekanisme return error kalau kosong
	}
	var ruangan = model.Room
	if err := s.db.Where(model.Room{ID: req.Id}).First(&ruangan).Error; err != nil {
		return nil, err
	}

	if (ruangan.Jumlah > 0) {
		tx := s.db.Begin()
		newPasien := model.Pasien{
			NamaPasien: req.Pasien.Nama,
			Diagnosa: req.Pasien.Diagnosa
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
			Pesan: "Ruangan berhasil dipesan"
		}, nil
	}
}