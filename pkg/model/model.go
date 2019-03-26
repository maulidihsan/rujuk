package model

type Pasien struct {
	ID uint `gorm:"primary_key" json:"id"`
	NamaPasien string `gorm:"type:varchar(32);not null;" json:"nama_pasien"`
	Diagnosa string `gorm:"type:varchar(128);not null;" json:"diagnosa"`
}

type Room struct {
	ID uint `gorm:"primary_key" json:"id"`
	NamaRuangan string `gorm:"type:varchar(32);not null;" json:"nama_ruangan"`
	TipeRuangan string `gorm:"type:varchar(16);not null;" json:"tipe"`
	Jumlah int32 `gorm:"type:int(8);not null;" json:"jumlah"`
}

type Rumahsakit struct {
	ID uint `gorm:"primary_key" json:"id"`
	NamaRumahsakit string `gorm:"type:varchar(32);not null;" json:"nama_rs"`
	IP string `gorm:"type:varchar(16);not null;" json:"ip"`
}

type Request struct {
	Offset    string `form:"offset" json:"offset"`
	Limit string `form:"limit" json:"limit"`
}

type RequestRujuk struct {
	Id    string `form:"id" json:"id"`
	Pasien Pasien `form:"pasien" json:"pasien"`
}