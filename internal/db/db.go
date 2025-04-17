package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenDBConnection(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
