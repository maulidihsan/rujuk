package model

type Pasien struct {
	ID uint `gorm:"primary_key" json:"id"`
	NamaPasien string `gorm:"type:varchar(32);not null;" json:"nama_pasien" binding:"required"`
	Diagnosa string `gorm:"type:varchar(128);not null;" json:"diagnosa" binding:"required"`
}

type Room struct {
	ID uint `gorm:"primary_key" json:"id"`
	NamaRuangan string `gorm:"type:varchar(32);not null;" json:"nama_ruangan" binding:"required"`
	TipeRuangan string `gorm:"type:varchar(16);not null;" json:"tipe" binding:"required"`
	Jumlah int32 `gorm:"type:int(8);not null;" json:"jumlah" binding:"required"`
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

type TransferRujuk struct {
	Id    uint `form:"id" json:"id" binding:"required"`
	Ip string `json:"ip"`
	Pasien Pasien `form:"pasien" json:"pasien" binding:"required"`
}

type RequestRujuk struct {
	Id    uint `form:"id" json:"id" binding:"required"`
	Pasien Pasien `form:"pasien" json:"pasien" binding:"required"`
}