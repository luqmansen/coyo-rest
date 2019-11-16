package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/luqmansen/coyo-rest/api/models"
	"log"
	"time"
)

var ktas = []models.KTA{
	models.KTA{
		ID:                 1,
		Name:               "CTBC Bank KTA Dana Cinta",
		Icon:               "https://kreditgogo.com/img/logo/ctbc-bank.png",
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
	models.KTA{
		ID:                 2,
		Name:               "Bank DBS Dana Bantuan Sahabat",
		Icon:               "https://kreditgogo.com/img/logo/bank-dbs.png",
		ShortDescription:   "Bank DBS menghadirkan produk Dana Bantuan Sahabat yang memberi anda fasilitas pinjaman dana tunai tanpa agunan dengan suku bunga rendah, proses cepat dan plafon pinjaman sampai dengan 200 juta.",
		Pembiayaan:         "Rp 300 juta",
		Tenor:              "3 Tahun",
		Pencairan:          "3 Hari",
		SukuBunga:          "10,56% sampai 17,88%",
		BiayaProses:        "Sesuai syarat dan ketentuan yang berlaku",
		BiayaAdmin:         "Rp 399.000",
		BiayaAsuransi:      "1% jumlah cicilan",
		BiayaProvisi:       "1,75% dari jumlah pinjaman hanya untuk tahun pertama",
		BeaMaterai:         "Rp.6000",
		BiayaPelunasan:     "8% sisa cicilan pokok",
		BiayaKeterlambatan: "Rp 250.000 ditambah 6% jumlah cicilan",
		MiniumPedapatan:    3,
		UsiaMinimum:        21,
		UsiaMaksimum:       55,
		Pendaftar:          "Warga negara Indonesia atau pemegang KITAS, Karyawan, Profesional, Wiraswasta",
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
	for i, _ := range ktas {
		err = db.Debug().Model(&models.KTA{}).Create(&ktas[i]).Error
		if err != nil {
			log.Fatalf("cannot seed ktas table: %v", err)
		}
	}

}
