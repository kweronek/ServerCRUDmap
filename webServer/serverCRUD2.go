package webServer

import (
	"fmt"
	"log"
	"net/http"
	"tech-bricks/http/json/ServerCRUDmap/controller"
	"tech-bricks/http/json/ServerCRUDmap/serviceGlobals"
	"time"
)

func StartWebserver() {

	controller.SetRoutesResource()

	// configuration parameters for webserver
	s := &http.Server{
		Addr:           serviceGlobals.SvcGlob.Port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("\nWebserver is started and listening on port", serviceGlobals.SvcGlob.Port,"\n")
	log.Fatal(s.ListenAndServe())
}