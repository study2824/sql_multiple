package main

import (
	"sql_multiple1/db"
	"sql_multiple1/server"
)

func main() {
	db.Init()
	server.Init()
}