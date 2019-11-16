package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/luqmansen/coyo-rest/api/models"
	"github.com/luqmansen/coyo-rest/api/responses"
	"github.com/luqmansen/coyo-rest/api/utils/formaterror"
	"io/ioutil"
	"net/http"
)

func (server *Server) AddEntryPengajuan(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Pengajuan{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data.PreparePengajuan()
	err = data.ValidatePengajuan()
	if err != nil {
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
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, dataCreated))
	responses.JSON(w, http.StatusCreated, dataCreated)
}
