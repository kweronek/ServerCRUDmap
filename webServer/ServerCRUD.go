
package webServer
/*
import (
	"fmt"
	"net/http"
	"tech-bricks/http/json/ServerCRUD/controllerResource"
	"tech-bricks/http/json/ServerCRUD/modelResource"
)

func main() {

	//	next line just generates some test data records:
	modelResource.Init()

	// next define the endpoints and the functions to be called
	// this fuctionality is the mux (multiplexer)
	// there are mux frameworks (e.g. Gorilla).
	// Usualley they are not necessary for microservices
	// However, they are very useful for multiple endpoint APIs
	//
	http.HandleFunc("/resources", controllerResource.ResourcesHandleFunc)
	http.HandleFunc("/resources/", controllerResource.ResourceHandleFunc)
	http.HandleFunc("/", controllerResource.Home)
	http.HandleFunc("/about", about)


	// configuration parameters for webserver
	Port := ":8080"

	// log server start and start server
	fmt.Println("Server started and listening on port " + Port)
	http.ListenAndServe(Port, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	// response for version information
	fmt.Fprint(w, "KWer's CRUD-REST-API for resources: v0.0.7, 08/06/2019 11:56 MEST.")
}
*/