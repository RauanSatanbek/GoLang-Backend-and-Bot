package db

import (
	"database/sql"
	"fmt"
	"log"
	"makebex-backend/server/config"
)

//-* Connect to DataBase *-//
func DB() (*sql.DB) {
	db, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil { log.Fatal(err) }

	err = db.Ping()
	if err != nil { log.Fatal(err) }

	fmt.Println("Successfully connected to DB")

	return db
}