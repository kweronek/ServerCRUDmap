package funcResource

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tech-bricks/http/json/ServerCRUDmap/modelResource"
	"tech-bricks/http/json/ServerCRUDmap/serviceGlobals"
)

func PostResource(body string) (Result string, HttpStatus int) {

	var toBeCreated = make([]modelResource.Resource, 1, 1)
	var indexCreated = make([]int, 0, 0)

	// JSON single object
	if body[0] == '{' {
		err := json.Unmarshal([]byte(body), &toBeCreated[0])
		if err != nil {
			log.Println("** could not create new resource")
			return "{\"ID\": -1}", http.StatusPartialContent
		} else {
			NewID := serviceGlobals.RecCnt.NextValue()
			modelResource.MMyResources[NewID] = toBeCreated[0]
			return "{\"ID\": " + strconv.Itoa(NewID) + "}", http.StatusCreated
		}
	}

	// JSON array:
	if body[0] == '[' {
		err := json.Unmarshal([]byte(body), &toBeCreated)
		if err != nil {
			log.Println("** could not create new resources")
			return "{\"ID\": -1}", http.StatusPartialContent
		} else {
			for i := 0; i < len(toBeCreated); i++ {
				NewID := serviceGlobals.RecCnt.NextValue()
				modelResource.MMyResources[NewID] = toBeCreated[i]
				indexCreated = append(indexCreated, NewID)
			}
			jsonData, err := json.Marshal(indexCreated)
			if err != nil {
				log.Println("** could not create new resources")
				return "{\"ID\": -1}", http.StatusInternalServerError
			}
			return "{\"ID\": " + string(jsonData) + "}", http.StatusCreated
		}
	}
	return "{\"ID\": -1}", http.StatusBadRequest
}
