package funcResource

import (
	"encoding/json"
	"log"
	"net/http"
	"tech-bricks/http/json/ServerCRUDmap/modelResource"
)

func GetResource(ID int) (result modelResource.Resource, httpStatus int) {
	if CheckResource(ID) {
		return modelResource.MMyResources[ID], http.StatusOK
	} else {
		log.Println("** resource", ID, "does not exist!")
		return modelResource.Resource{}, http.StatusNotFound
	}
}

func GetAllResources() (result string, httpStatus int) {
	var jsonData []byte
	//	persistResource.SelectFrom(modelResource.Resource())
	var err error
	jsonData, err = json.Marshal(modelResource.MMyResources)
	if err != nil {
		return "", http.StatusUnauthorized
	} else {
		return string(jsonData), http.StatusOK
	}
}
