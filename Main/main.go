package main

import (
	"challenge/Server"
	"challenge/Storage/database"
	"challenge/Storage/logs"
	"log"
	"net/http"
)
// main function that calls the server and the database connection
func main () {
	_ = logs.InitLogger()

	_ = database.NewRoachClient()
	r := Server.NewServer()
	log.Fatal(http.ListenAndServe(":3000", r))
}
