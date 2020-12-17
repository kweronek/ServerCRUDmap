package controller

import (
	"net/http"
)

func  SetRoutesResource() {
//	s.router.HandleFunc("/api/", s.handleAPI())
//	s.router.HandleFunc("/about", s.handleAbout())
//	s.router.HandleFunc("/", s.handleIndex())

	http.HandleFunc("/resources", ResourcesHandleFunc)
	http.HandleFunc("/resources/", ResourceHandleFunc)
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/About", About)
}