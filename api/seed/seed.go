package seed

import (
	"github.com/luqmansen/coyo-rest/api/models"
	"log"
	"github.com/jinzhu/gorm"
	"time"
)

var posts = []models.KTA{
	models.KTA{
		ID:                 0,
		Title:              "CTBC Bank KTA Dana Cinta",
		Icon:               "https://kreditgogo.com/img/logo/bank-dbs.png",
		ShortDescription:   "Butuh dana pinjaman? KTA Dana Cinta dari Bank CTBC memberikan anda pinjaman dengan proses cepat, angsuran tetap, dan suku bunga rendah. Dengan nilai pinjaman hingga Rp200 juta, KTA Dana cinta cocok untuk keperluan bisnis anda.",
		Pembiayaan:         "Rp 200 Juta",
		Tenor:              "5 Tahun",
		Pencairan:          "3 Hari",
		SukuBunga:          "15,48% sampai 19,08%",
		BiayaProses:        "3% Dari Jumlah Pinjaman",
		BiayaAdmin:         "Sesuai syarat dan ketentuan yang berlaku",
		BiayaAsuransi:      "Tidak Ada",
		BiayaProvisi:       "Tidak Ada",
		BeaMaterai:         "Rp.6000",
		BiayaPelunasan:     "9% sisa cicilan pokok , Biaya pelunasan yang dipercepat",
		BiayaKeterlambatan: "5% jumlah cicilan",
		MiniumPedapatan:    2.5,
		UsiaMinimum:        21,
		UsiaMaksimum:       55,
		Pendaftar:          "Karyawan, Wiraswasta,Profesional",
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.KTA{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.KTA{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	//err = db.Debug().Model(&models.KTA{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	//if err != nil {
	//	log.Fatalf("attaching foreign key error: %v", err)
	//}

	err = db.Debug().Model(&models.KTA{}).Create(&posts[0]).Error
	if err != nil {
		log.Fatalf("cannot seed posts table: %v", err)
	}

}

