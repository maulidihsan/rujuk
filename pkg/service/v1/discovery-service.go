package v1

import (
	"context"
	"github.com/jinzhu/gorm"

	"github.com/maulidihsan/rujuk/pkg/api/v1"
	"github.com/maulidihsan/rujuk/pkg/model"
)

type service struct {
	db *gorm.DB
}

func NewDiscoveryServiceServer(db *gorm.DB) v1.DiscoveryServiceServer {
	if !db.HasTable(&model.Rumahsakit{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Rumahsakit{}).Error; err != nil {
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
		Id: uint32(in.ID),
		Nama: in.NamaRumahsakit,
		Ip: in.IP,
	}
	return res
}

func (s *service) GetAllRumahsakit(ctx context.Context, req *v1.FetchRequest) (*v1.Rumahsakits, error) {
	offset := int32(0)
	limit := int32(10)
	if req != nil {
		offset = req.Offset
		limit = req.Limit
	}
	var rumahsakits []model.Rumahsakit
	if err := s.db.Model(&model.Rumahsakit{}).Offset(offset).Limit(limit).Find(&rumahsakits).Error; err != nil {
		return nil, err
	}
	payload := make([]*v1.Rumahsakit, len(rumahsakits))
	for i, rs := range rumahsakits {
		r := s.transformRumahsakitToRPC(&rs)
		payload[i] = r
	}

	reply := &v1.Rumahsakits{
		Rumahsakits: payload,
	}
	return reply, nil
}

func (s *service) Register(ctx context.Context, req *v1.Rumahsakit) (*v1.Response, error) {
	if req == nil {
		//TODO: mekanisme return error kalau kosong
	}
	rumahsakit := model.Rumahsakit{
		NamaRumahsakit: req.Nama,
		IP: req.Ip,
	}
	if err := s.db.Where(model.Rumahsakit{IP: req.Ip}).FirstOrCreate(&rumahsakit).Error; err != nil {
		return nil, err
	}

	return &v1.Response{
		Kode: 200,
		Pesan: "Registrasi berhasil",
	}, nil
}