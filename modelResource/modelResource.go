package modelResource

import (
	"encoding/json"
	"fmt"
	"os"
)

// The model
type Resource struct {
	ID       int     `json:"ID"`
	Vorname  string  `json:"Vorname"`
	Nachname string  `json:"Nachname"`
	Alter    int     `json:"Alter"`
	Score    float64 `json:"Score"`
	Present  bool    `json:"Present"`
}

var TestResource = Resource{1000,
	"Testuser", "Testname", 105, 57.8, true}

var MMyResources = make(map[int]Resource)

// Deklarationen
func Init() {
	// some data for test:
	MMyResources[0] = Resource{}
	MMyResources[1] = Resource{
		ID:       1,
		Vorname:  "Erika",
		Nachname: "Musterfrau",
		Alter:    35,
		Score:    1.75,
		Present:  true,
	}
	MMyResources[2] = Resource{2,
		"Hans", "Mustermann", 56, 3.5, true}
	MMyResources[3] = Resource{3,
		"Kim", "Muster", 45, 1.66E-25, true}

	fmt.Println("MMyResources initialised!")

	writer := os.Stdout
	enc := json.NewEncoder(writer)
	enc.Encode(MMyResources)

	/*
		fmt.Println("MyResources initialised!")
		enc.Encode(&MMyResources)

		//	part for maps
		var MResourceA = MResource{
			"ID":       "1",
			"Vorname":  "Max",
			"Nachname": "Mustermann",
			"Alter":    "34",
			"Score":    "1.45",
		}
		var MResourceB = MResource{
			"ID":       "2",
			"Vorname":  "Erika",
			"Nachname": "Musterfrau",
			"Alter":    "34",
			"Score":    "1.45",
		}

		var MResourceC = MResource{
			"ID":       "3",
			"Vorname":  "Kim",
			"Nachname": "Muster",
			"Alter":    "45",
			"Score":    "1.65",
		}

		var MResourceEmtpy = MResource{
			"ID":       "",
			"Vorname":  "",
			"Nachname": "",
			"Alter":    "",
			"Score":    "",
		}
	*/
	/*
		MMyResources = make(map[int]Resource)
		MMyResources[0] = TestResource
		MMyResources[1] = MyResources[1]
		MMyResources[2] = MyResources[2]
		MMyResources[3] = MyResources[3]
	*/
	/*
	   for i := 0; i < len(MMyResources); i++ {

	   		fmt.Printf("%s %2d %s %s %s %s %s %s\n",
	   			"Nr.", i, ":",
	   			MMyResources[i]["ID"],
	   			MMyResources[i]["Vorname"],
	   			MMyResources[i]["Nachname"],
	   			MMyResources[i]["Alter"],
	   			MMyResources[i]["Score"], )
	   	}
	*/
}
