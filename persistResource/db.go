package persistResource

import (
	"database/sql"
	"fmt"
	_	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDatabase() {
	fmt.Println("Database connected!\n")
	db, err = sql.Open("mysql", "root:root@/test")
}
