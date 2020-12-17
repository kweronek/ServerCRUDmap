package main

import (
	"tech-bricks/http/json/ServerCRUDmap/modelResource"
	"tech-bricks/http/json/ServerCRUDmap/persistResource"
	"tech-bricks/http/json/ServerCRUDmap/webServer"
)

func main() {
	// connect database
	persistResource.ConnectDatabase()

	//	next line just generates some test data records:
	modelResource.Init()

	// start webserver
	webServer.StartWebserver()
}
