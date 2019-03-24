package model
import "github.com/jinzhu/gorm"

type Pasien struct {
	gorm.Model
	NamaPasien string `gorm:"type:varchar(32);not null;"`
	Diagnosa string `gorm:"type:varchar(128);not null;"`
}

type Room struct {
	gorm.Model
	NamaRuangan string `gorm:"type:varchar(32);not null;"`
	TipeRuangan string `gorm:"type:varchar(16);not null;"`
	Jumlah int32 `gorm:"type:int(8);not null;"`
}

type Rumahsakit struct {
	gorm.Model
	NamaRumahsakit string `gorm:"type:varchar(32);not null;"`
	IP string `gorm:"type:varchar(16);not null;"`
}