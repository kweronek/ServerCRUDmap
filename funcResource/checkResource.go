package funcResource

import (
	"tech-bricks/http/json/ServerCRUDmap/modelResource"
)

func CheckResource(ID int) (exists bool) {
	_, exists = modelResource.MMyResources[ID]
	return exists
}
