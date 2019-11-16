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

func (server *Server) AddEntry(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.KTA{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data.Prepare()
	err = data.Validate()
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
	dataCreated, err := data.AddEntry(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, dataCreated.ID))
	responses.JSON(w, http.StatusCreated, dataCreated)
}

func (server *Server) GetKTAs(w http.ResponseWriter, r *http.Request) {

	kta := models.KTA{}

	posts, err := kta.FindAllKTAs(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	data := make(map[string]interface{})
	data["data"] = posts
	responses.JSON(w, http.StatusOK, data)
}

//func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	pid, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	post := models.Post{}
//
//	postReceived, err := post.FindPostByID(server.DB, pid)
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	responses.JSON(w, http.StatusOK, postReceived)
//}

//func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	// Check if the post id is valid
//	pid, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	//CHeck if the auth token is valid and  get the user id from it
//	uid, err := auth.ExtractTokenID(r)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//
//	// Check if the post exist
//	post := models.Post{}
//	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
//	if err != nil {
//		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
//		return
//	}
//
//	// If a user attempt to update a post not belonging to him
//	if uid != post.AuthorID {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	// Read the data posted
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	// Start processing the request data
//	postUpdate := models.Post{}
//	err = json.Unmarshal(body, &postUpdate)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	//Also check if the request user id is equal to the one gotten from token
//	if uid != postUpdate.AuthorID {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//
//	postUpdate.Prepare()
//	err = postUpdate.Validate()
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	postUpdate.ID = post.ID //this is important to tell the model the post id to update, the other update field are set above
//
//	postUpdated, err := postUpdate.UpdateAPost(server.DB)
//
//	if err != nil {
//		formattedError := formaterror.FormatError(err.Error())
//		responses.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	responses.JSON(w, http.StatusOK, postUpdated)
//}
//
//func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	// Is a valid post id given to us?
//	pid, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	// Is this user authenticated?
//	uid, err := auth.ExtractTokenID(r)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//
//	// Check if the post exist
//	post := models.Post{}
//	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
//	if err != nil {
//		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
//		return
//	}
//
//	// Is the authenticated user, the owner of this post?
//	if uid != post.AuthorID {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	_, err = post.DeleteAPost(server.DB, pid, uid)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
//	responses.JSON(w, http.StatusNoContent, "")
//}
