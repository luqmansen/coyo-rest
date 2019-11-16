package controllers

import (
	"fmt"
	"github.com/luqmansen/coyo-rest/api/models"
	"github.com/luqmansen/coyo-rest/api/responses"
	"github.com/luqmansen/coyo-rest/api/utils/formaterror"
	"log"
	"net/http"
)

func (server *Server) AddEntryPengajuan(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	for key, value := range r.Form {
		fmt.Printf("%s - %s \n", key, value)
	}
	//var c int
	data := models.Pengajuan{}
	//for data.ValidatePengajuan() != nil {
	data.Nama = r.Form.Get("nama")
	data.TempatLahir = r.Form.Get("tempat_lahir")
	data.TanggalLahir = r.Form.Get("tanggal_lahir")
	data.KotaDomisili = r.Form.Get("kota_domisili")
	data.Telepon = r.Form.Get("telepon")
	data.AjuanKTA = r.Form.Get("ajuan_kta")
	data.JumlahPenghasilan = r.Form.Get("penghasilan")
	data.JumlahPinjaman = r.Form.Get("jumlah_pinjaman")
	//fmt.Print(c)
	//c++
	//}

	//decoder := schema.NewDecoder()
	//err = decoder.Decode(&data, r.PostForm)

	data.PreparePengajuan()
	err = data.ValidatePengajuan()
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//uid, err := auth.ExtractTokenID(r)
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	//	return
	//}
	//if uid != data.AuthorID {
	//	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	//	return
	//}
	dataCreated, err := data.AddEntryPengajuan(server.DB)
	if err != nil {
		fmt.Println(err)
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, dataCreated))
	responses.JSON(w, http.StatusCreated, dataCreated)
}
