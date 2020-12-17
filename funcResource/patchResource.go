package funcResource

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"tech-bricks/http/json/ServerCRUDmap/modelResource"
)

//func PatchResource(pID int, patchJSON string) (httpStatus int) {
func PatchResource(pID int, patchJSON string) (httpStatus int) {
	tbpRes, httpStatus := GetResource(pID)
	if httpStatus != 200 {
		return (http.StatusNotFound)
	}
	var rVOtbpRes = reflect.ValueOf(tbpRes)

	var patchMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(patchJSON), &patchMap)
	if err != nil {
		return http.StatusBadRequest
	}

	for k, v := range patchMap	 {
//		log.Println(k, v, rVOtbpRes.FieldByName(k).Kind())
		if rVOtbpRes.FieldByName(k).IsValid() {
			switch rVOtbpRes.FieldByName(k).Kind() {
			case reflect.String:
				ins, _ := v.(string)
				reflect.ValueOf(&tbpRes).Elem().FieldByName(k).SetString(ins)
			case reflect.Int:
				ins := int64(v.(float64))
				reflect.ValueOf(&tbpRes).Elem().FieldByName(k).SetInt(ins)
			case reflect.Float64:
				ins, _ := v.(float64)
				reflect.ValueOf(&tbpRes).Elem().FieldByName(k).SetFloat(ins)
			case reflect.Bool:
				ins, _ := v.(bool)
				reflect.ValueOf(&tbpRes).Elem().FieldByName(k).SetBool(ins)
			default:
				log.Println("** type (", rVOtbpRes.FieldByName(k).Kind(), ") to set not implemented yet")
			}
		}
	}
	modelResource.MMyResources[pID] = tbpRes
	return http.StatusOK
}
