package database

import (
	"challenge/Storage/logs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Db struct
type RoachClient struct {
	*sql.DB
}
// DB Connection
func NewRoachClient () *RoachClient {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	if err != nil {
		logs.Log().Errorf("error connecting to the database: %s ", err.Error())
		panic(err)
	}
	fmt.Println("Database open")
	// ping the DB in case of no return, return error
	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to cockroach")
	}
// create the initial table into the data base if not exists
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS players_tb(id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name STRING, score STRING)"); err != nil {
		fmt.Println("cannot create table", err)
	}

	return &RoachClient{db}
}
