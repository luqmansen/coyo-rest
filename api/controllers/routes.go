package controllers

import "github.com/luqmansen/coyo-rest/api/middlewares"

func (s *Server) initializeRoutes() {

	//KTA routes
	s.Router.HandleFunc("/kta/all", middlewares.SetMiddlewareJSON(s.GetKTAs)).Methods("GET")
	s.Router.HandleFunc("/kta/add", middlewares.SetMiddlewareJSON(s.AddEntryKTA)).Methods("POST")

	//Pengajuan Routes
	//s.Router.HandleFunc("/pengajuan/all", middlewares.SetMiddlewareJSON(s.GetPengajuan)).Methods("GET")
	s.Router.HandleFunc("/pengajuan/add", middlewares.SetMiddlewareJSON(s.AddEntryPengajuan)).Methods("POST")

}
