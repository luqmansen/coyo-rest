package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type KTA struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name             string `gorm:"size:255;not null;" json:"title"`
	Icon             string `gorm:"size:255;not null;" json:"icon"`
	ShortDescription string `gorm:"size:255;not null;" json:"short_description"`
	Pembiayaan       string `gorm:"size:255;not null;" json:"pembiayaan"`
	Tenor            string `gorm:"size:255;not null;" json:"tenor"`
	Pencairan        string `gorm:"size:255;not null;" json:"pencairan"`
	SukuBunga        string `gorm:"size:255;not null;" json:"suku_bunga"`
	BiayaProses      string `gorm:"size:255;not null;" json:"biaya_proses"`
	BiayaAdmin       string `gorm:"size:255;not null;" json:"biaya_admin"`
	BiayaAsuransi    string `gorm:"size:255;not null;" json:"biaya_asuransi"`
	BiayaProvisi     string `gorm:"size:255;not null;" json:"biaya_provisi"`
	BeaMaterai         string    `gorm:"size:255;not null;" json:"bea_materai"`
	BiayaPelunasan     string    `gorm:"size:255;not null;" json:"biaya_pelunasan"`
	BiayaKeterlambatan string    `gorm:"size:255;not null;" json:"biaya_keterlambatan"`
	MiniumPedapatan    float32       `gorm:"size:255;not null;" json:"minimum_pedapatan"`
	UsiaMinimum        int       `gorm:"size:255;not null;" json:"usia_minimum"`
	UsiaMaksimum       int       `gorm:"size:255;not null;" json:"usia_maksimum"`
	Pendaftar          string    `gorm:"size:255;not null;" json:"pendaftar"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (k *KTA) Prepare() {
	k.ID = 0
	k.Name = html.EscapeString(strings.TrimSpace(k.Name))
	k.CreatedAt = time.Now()
	k.UpdatedAt = time.Now()
}

func (k *KTA) Validate() error {

	if k.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (k *KTA) AddEntry(db *gorm.DB) (*KTA, error) {
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
//
//func (p *KTA) FindPostByID(db *gorm.DB, pid uint64) (*KTA, error) {
//	var err error
//	err = db.Debug().Model(&KTA{}).Where("id = ?", pid).Take(&p).Error
//	if err != nil {
//		return &KTA{}, err
//	}
//	if p.ID != 0 {
//		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
//		if err != nil {
//			return &KTA{}, err
//		}
//	}
//	return p, nil
//}
//
//func (p *KTA) UpdateAPost(db *gorm.DB) (*KTA, error) {
//
//	var err error
//	// db = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&Post{}).UpdateColumns(
//	// 	map[string]interface{}{
//	// 		"title":      p.Name,
//	// 		"content":    p.Description,
//	// 		"updated_at": time.Now(),
//	// 	},
//	// )
//	// err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
//	// if err != nil {
//	// 	return &Post{}, err
//	// }
//	// if p.ID != 0 {
//	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
//	// 	if err != nil {
//	// 		return &Post{}, err
//	// 	}
//	// }
//	err = db.Debug().Model(&KTA{}).Where("id = ?", p.ID).Updates(KTA{Name: p.Name, Description: p.Description, UpdatedAt: time.Now()}).Error
//	if err != nil {
//		return &KTA{}, err
//	}
//	if p.ID != 0 {
//		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
//		if err != nil {
//			return &KTA{}, err
//		}
//	}
//	return p, nil
//}
//
//func (p *KTA) DeleteAPost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {
//
//	db = db.Debug().Model(&KTA{}).Where("id = ? and author_id = ?", pid, uid).Take(&KTA{}).Delete(&KTA{})
//
//	if db.Error != nil {
//		if gorm.IsRecordNotFoundError(db.Error) {
//			return 0, errors.New("Post not found")
//		}
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}
//
