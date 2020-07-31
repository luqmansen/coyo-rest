package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type KTA struct {
	ID                 uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name               string    `gorm:"size:255;not null;" json:"title"`
	Icon               string    `gorm:"size:255;not null;" json:"icon"`
	ShortDescription   string    `gorm:"size:255;not null;" json:"short_description"`
	Pembiayaan         string    `gorm:"size:255;not null;" json:"pembiayaan"`
	Tenor              string    `gorm:"size:255;not null;" json:"tenor"`
	Pencairan          string    `gorm:"size:255;not null;" json:"pencairan"`
	SukuBunga          string    `gorm:"size:255;not null;" json:"suku_bunga"`
	BiayaProses        string    `gorm:"size:255;not null;" json:"biaya_proses"`
	BiayaAdmin         string    `gorm:"size:255;not null;" json:"biaya_admin"`
	BiayaAsuransi      string    `gorm:"size:255;not null;" json:"biaya_asuransi"`
	BiayaProvisi       string    `gorm:"size:255;not null;" json:"biaya_provisi"`
	BeaMaterai         string    `gorm:"size:255;not null;" json:"bea_materai"`
	BiayaPelunasan     string    `gorm:"size:255;not null;" json:"biaya_pelunasan"`
	BiayaKeterlambatan string    `gorm:"size:255;not null;" json:"biaya_keterlambatan"`
	MiniumPedapatan    float32   `gorm:"size:255;not null;" json:"minimum_pedapatan"`
	UsiaMinimum        int       `gorm:"size:255;not null;" json:"usia_minimum"`
	UsiaMaksimum       int       `gorm:"size:255;not null;" json:"usia_maksimum"`
	Pendaftar          string    `gorm:"size:255;not null;" json:"pendaftar"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (k *KTA) PrepareKTA() {
	k.ID = 0
	k.Name = html.EscapeString(strings.TrimSpace(k.Name))
	k.CreatedAt = time.Now()
	k.UpdatedAt = time.Now()
}

func (k *KTA) ValidateKTA() error {

	if k.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (k *KTA) AddEntryKTA(db *gorm.DB) (*KTA, error) {
	var err error

	err = db.Debug().Model(&KTA{}).Create(&k).Error
	if err != nil {
		return &KTA{}, err
	}
	return k, nil
}

func (k *KTA) FindAllKTAs(db *gorm.DB) (*[]KTA, error) {
	var err error
	ktas := []KTA{}
	err = db.Debug().Model(&KTA{}).Limit(100).Find(&ktas).Error
	if err != nil {
		return &[]KTA{}, err
	}
	//if len(ktas) > 0 {
	//	for i, _ := range ktas {
	//		err := db.Debug().Model(&User{}).Where("id = ?", ktas[i].AuthorID).Take(&ktas[i].Author).Error
	//		if err != nil {
	//			return &[]KTA{}, err
	//		}
	//	}
	//}
	return &ktas, nil
}
