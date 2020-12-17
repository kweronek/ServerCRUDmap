package persistResource

import (
	"tech-bricks/http/json/ServerCRUD/modelResource"
)

// SELECT * FROM RESOURCE
func SelectFrom(table modelResource.Resource)(result ResourceTable){
	return nil
}

// SELECT * FROM RESOURCE where ID=idVal
func SelectFromWhere(table ResourceTable, idVal int) (resultRecord modelResource.Resource) {
	return
}

// INSERT INTO table_name VALUES (value1, value2, value3, ...);

// INSERT RESOURCE
func Insert(newRecord modelResource.Resource) (newID int){
	return
}

// DELETE * FROM RESOURCE WHERE ID=idVal
func Delete(table modelResource.Resource, idVal int) (err error){
	return nil
}

// IF EXISTS (SELECT TOP 1 1 FROM Resource WHERE ID=IDval)
func Exists(IDval int) (err error){
	return
}
