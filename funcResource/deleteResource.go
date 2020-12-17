package funcResource

import "tech-bricks/http/json/ServerCRUDmap/modelResource"

func DeleteResource(ID int) {
	delete(modelResource.MMyResources, ID)
}
