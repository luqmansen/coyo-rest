package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Pengajuan struct {
	ID                uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Nama              string    `gorm:"size:255;not null;" json:"nama"`
	Telepon           string    `gorm:"size:255;not null;" json:"telepon"`
	TempatLahir       string    `gorm:"size:255;not null;" json:"tempat_lahir"`
	TanggalLahir      string    `gorm:"size:255;not null;" json:"tanggal_lahir"`
	KotaDomisili      string    `gorm:"size:255;not null;" json:"kota_domisili"`
	AjuanKTA          string    `gorm:"size:255;not null" json:"ajuan_kta"`
	JumlahPenghasilan string    `gorm:"size:255;not null" json:"penghasilan"`
	JumlahPinjaman    string    `gorm:"size:255;not null" json:"jumlah_pinjaman"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Pengajuan) PreparePengajuan() {
	p.ID = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
func (p *Pengajuan) ValidatePengajuan() error {

	if p.Nama == "" {
		return errors.New("Required Name")
	}
	if p.Telepon == "" {
		return errors.New("Required Telepon")
	}
	if p.TempatLahir == "" {
		return errors.New("Required ttl")
	}
	if p.TanggalLahir == "" {
		return errors.New("Required ttl")
	}
	if p.KotaDomisili == "" {
		return errors.New("Required domisili")
	}
	if p.AjuanKTA == "" {
		return errors.New("Required ajuankta")
	}
	if p.JumlahPinjaman == "" {
		return errors.New("Required jumlahpinjam")
	}
	if p.JumlahPenghasilan == "" {
		return errors.New("Required penghasilan")
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
		log.Printf("error decoding response: %v", err)
		return &Pengajuan{}, err
	}

	return p, nil
}
