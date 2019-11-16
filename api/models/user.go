package models

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Nama         string    `gorm:"size:255;not null;" json:"nama"`
	Telepon      string    `gorm:"size:255;not null;" json:"telepon"`
	TempatLahir  string    `gorm:"size:255;not null;" json:"tempat_lahir"`
	TanggalLahir string    `gorm:"size:255;not null;" json:"tanggal_lahir"`
	KotaDomisili string    `gorm:"size:255;not null;" json:"kota_domisili"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
