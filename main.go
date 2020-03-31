package main

import (
	"golang-crud-api/app"
	"golang-crud-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
