package controller

import (
	"encoding/json"
	"fmt"
	"helper"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"tech-bricks/http/json/ServerCRUDmap/funcResource"
	"tech-bricks/http/json/ServerCRUDmap/serviceGlobals"
	"tech-bricks/http/json/ServerCRUDmap/viewModelResource"
)

// home page
func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.Method, r.Host+r.URL.Path, r.Proto, r.Header["Content-Type"])

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("** HTTP-method " + r.Method + " not supported here. Use GET!")
	}

	switch r.Header.Get("Accept") {

	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var t, err = template.ParseFiles("viewResource/home.html")
		if err != nil {
			log.Print("Unable to load template")
			return
		}
		err = t.Execute(w, viewModelResource.Home)
		if err != nil {
			log.Println("t.Execute not possible")
		}

	case "text/plain":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w,
			"KWer's CRUD-REST-API for resources: ",
			serviceGlobals.SvcGlob.Version,
			" 08/06/2019 11:56 MEST.")
	}
}

// ********************************************************************************************************************
// handler for resources without ID
// ********************************************************************************************************************
func ResourcesHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.Method, r.Host+r.URL.Path, r.Proto, r.Header["Content-Type"])

	switch r.Method {

	case "GET":
		result, httpStatus := funcResource.GetAllResources()
		if httpStatus != 200 {
			w.WriteHeader(http.StatusBadRequest)
			log.Print(httpStatus)
		} else {
			w.Header().Set("Content-Type", "application/json")
			var _, err = fmt.Fprintf(w, result)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
			}
		}
	case "POST":
		requestBody, httpStatus := reqBody(r)
		if httpStatus != 200 {
			w.WriteHeader(http.StatusBadRequest)
			log.Print(httpStatus)
		} else {
			result, httpStatus := funcResource.PostResource(requestBody)
			w.WriteHeader(httpStatus)
			fmt.Fprint(w, result)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Print("** method ", r.Method, " not supported without ID")
	}
}

// ********************************************************************************************************************
// Handler for resources with ID specified
// ********************************************************************************************************************
func ResourceHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.Method, r.Host+r.URL.Path, r.Proto, r.Header["Content-Type"])
	var pp = utils.ParsePath(r.URL.Path)

	ID, err := strconv.Atoi(pp[1])
	if err != nil {
		if pp[1] == "" {
			log.Print(string(pp[1]) + "** no resource specified, redirected to -->")
			ResourcesHandleFunc(w, r)
		} else {
			log.Print(string("** \""+pp[1]) + "\" is not a number")
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	switch r.Method {
	case "GET": // read a single resource
		resourceBuf, httpStatus := funcResource.GetResource(ID)
		if httpStatus != 200 {
			w.WriteHeader(http.StatusNotFound)
		} else {

			switch r.Header.Get("Accept") {

			case "": // this order leads to better performance
				fallthrough
			case "application/json":
				fallthrough
			default:
				w.Header().Set("Content-Type", "application/json")
				jsonData, err := json.Marshal(resourceBuf)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				fmt.Fprint(w, string(jsonData))

			case "text/html":
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				t, err := template.ParseFiles("viewResource/resource.html")
				if err != nil {
					log.Print(w, "Unable to load template")
				}
				err = t.Execute(w, resourceBuf)
				if err != nil {
					log.Print(err)
				}

			case "text/plain":
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprint(w, resourceBuf)
			}
		}

	case "POST": // create a new Resource
		var pp []string = utils.ParsePath(r.URL.Path)
		if pp[1] != "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("** bad request")
			return
		}

		var requestBody, httpStatus = reqBody(r)
		if httpStatus != 200 {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("** httpStatus: ", httpStatus)
		} else {
			switch r.Header.Get("Content-Type") {
			case "":
				fallthrough
			case "application/json":
				newID, httpStatus := funcResource.PostResource(string(requestBody))
				w.WriteHeader(httpStatus)
				fmt.Println(newID)
				fmt.Fprint(w, "{\nnew ID:\t", newID, "\n}")
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				log.Print("** unsupported media type")
			}
		}

	case "PUT": // replace a Resource completely
		var pp = utils.ParsePath(r.URL.Path)
		pID, err := strconv.Atoi(pp[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Header.Get("Content-Type") {
		case "", "application/json":
			requestBody, httpStatus := reqBody(r)
			if httpStatus != 200 {
				w.WriteHeader(httpStatus)
				return
			} else {
				httpStatus = funcResource.PutResource(pID, requestBody)
				w.WriteHeader(httpStatus)
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			log.Print("** unsupported media type")
		}

	case "PATCH": // change parts of a Resource
		var pp []string = utils.ParsePath(r.URL.Path)
		pID, err := strconv.Atoi(pp[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		requestBody, httpStatus := reqBody(r)
		if httpStatus != 200 {
			w.WriteHeader(httpStatus)
			return
		} else {
			switch r.Header.Get("Content-Type") {
			case "", "application/json":
				httpStatus := funcResource.PatchResource(pID, string(requestBody))
				w.WriteHeader(httpStatus)
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				log.Print("** unsupported media type")
			}
		}

	case "DELETE": // delete a resource
		if funcResource.CheckResource(ID) {
			funcResource.DeleteResource(ID)
		} else {
			log.Println("** resource ", ID, "does not exist!")
		}

	default: // all unsupported methods
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("** method: " + r.Method + "not supported for this URI!\n")
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.Method, r.Host+r.URL.Path, r.Proto, r.Header["Content-Type"])
	// response for version information
	switch r.Header.Get("Content-Type") {

	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(serviceGlobals.SvcGlob)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			log.Print("** internal Server Error")
		} else {
			fmt.Fprint(w, string(jsonData))
		}

	case "text/plain":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w,
			serviceGlobals.SvcGlob.APIname,
			serviceGlobals.SvcGlob.Version,
			serviceGlobals.SvcGlob.ReleaseDate,
		)

	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, err := template.ParseFiles("viewResource/about.html")
		if err != nil {
			log.Print("** unable to load template about.html !")
		} else {
			err = t.Execute(w, serviceGlobals.SvcGlob)
			if err != nil {
				log.Print(err)
			}
		}
	}
}
