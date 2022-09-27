package main

import (
	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver"
)

func main() {
	app := httpserver.CreateRouter()
	app.Run(":4000")
}
