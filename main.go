package main

import (
	"mail-sender/database"
	"mail-sender/server"
)

func main() {
	database.StartDB()

	s := server.NewServer()
	s.Run()
}