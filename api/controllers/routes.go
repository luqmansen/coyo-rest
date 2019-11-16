package controllers

import "github.com/luqmansen/coyo-rest/api/middlewares"

func (s *Server) initializeRoutes() {

	//KTA routes
	s.Router.HandleFunc("/kta/all", middlewares.SetMiddlewareJSON(s.GetKTAs)).Methods("GET")
	s.Router.HandleFunc("/kta/add", middlewares.SetMiddlewareJSON(s.AddEntryKTA)).Methods("POST")

	//Pengajuan Routes
	s.Router.HandleFunc("/kta/pengajuan/history", middlewares.SetMiddlewareJSON(s.GetHistory)).Methods("GET")
	s.Router.HandleFunc("/kta/pengajuan/add", middlewares.SetMiddlewareJSON(s.AddEntryPengajuan)).Methods("POST")

}
