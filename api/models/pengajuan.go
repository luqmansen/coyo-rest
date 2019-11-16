package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Pengajuan struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Nama           string    `gorm:"size:255;not null;" json:"nama"`
	Telepon        string    `gorm:"size:255;not null;" json:"telepon"`
	TempatLahir    string    `gorm:"size:255;not null;" json:"tempat_lahir"`
	TanggalLahir   string    `gorm:"size:255;not null;" json:"tanggal_lahir"`
	KotaDomisili   string    `gorm:"size:255;not null;" json:"kota_domisili"`
	AjuanKTA       string    `gorm:"foreign_key" json:"ajuan_kta"`
	JumlahPinjaman string    `gorm:"foreign_key" json:"jumlah_pinjaman"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Pengajuan) PreparePengajuan() {
	p.ID = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
func (k *Pengajuan) ValidatePengajuan() error {

	if k.Nama == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (p *Pengajuan) AddEntryPengajuan(db *gorm.DB) (interface{}, error) {
	var err error

	//err = db.Debug().Model(&User{}).Create(&p.Nasabah).Error
	//if err != nil {
	//	return &User{}, err
	//}
	err = db.Debug().Model(&Pengajuan{}).Create(&p).Error
	if err != nil {
		return &Pengajuan{}, err
	}
	return p, nil
}
