package persistResource

import (
//	"encoding/json"
	"fmt"
	"tech-bricks/http/json/ServerCRUD/modelResource"
//	"image"
//	"os"
)

type ResourceTable []modelResource.Resource
var MyResourcesTable = make(ResourceTable, 4, 10)

var EmptyResource modelResource.Resource = modelResource.Resource{}
var TestResource modelResource.Resource = modelResource.Resource{1000, "Testuser", "Testname", 105, 57.8, true}

type MResource map[string]string

var MMyResources = make([]MResource, 5, 10)

// Deklarationen
func persistInit() {

	MyResourcesTable[0] = EmptyResource
	MyResourcesTable[1] = modelResource.Resource{
		ID:       1,
		Vorname:  "Erika",
		Nachname: "Musterfrau",
		Alter:    35,
		Score:    1.75,
		Present:  true,
	}
	// some data for test:
	MyResourcesTable[2] = modelResource.Resource{2, "Hans", "Mustermann", 56, 3.5, true}
	MyResourcesTable[3] = modelResource.Resource{3, "Kim", "Muster", 45, 1.66E-25,true}
	fmt.Println("MyResources initialised")
}