package main

import (
	"go-auth/db"
	webserver "go-auth/web-server"
)

func main() {

	db.Process() // insert users.json to redis cache

	webserver.Start("8081") //start the web server
}
