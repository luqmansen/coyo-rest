package controllers

import "github.com/luqmansen/coyo-rest/api/middlewares"

func (s *Server) initializeRoutes() {


	//Posts routes
	s.Router.HandleFunc("/kta/all", middlewares.SetMiddlewareJSON(s.GetKTAs)).Methods("GET")
	s.Router.HandleFunc("/kta/add", middlewares.SetMiddlewareJSON(s.AddEntry)).Methods("POST")
	//s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetKTAs)).Methods("GET")
	//s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")


}

