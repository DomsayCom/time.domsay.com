package lib

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:fak@(localhost:3306)/time"
)

var db *sql.DB

var connectionError error

func StartDBConn() {

	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}

}

func init() {

	StartDBConn()

}
