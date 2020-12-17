package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

func reqBody(r *http.Request) (rB string, httpStatus int) {
	if r == nil {
		return "", http.StatusBadRequest
	} else {
		defer r.Body.Close()
		var reqBody, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			return "", http.StatusBadRequest
		}
		return string(reqBody), http.StatusOK
	}
}