package main

import (
	"fmt"
	"task/config"
	"task/database"
	app "task/src"
)

func main() {
	config.Init_config()
	database.Init_db()
	server := app.Init_app()
	address := fmt.Sprintf("%v:%v", config.ENV.API_HOST, config.ENV.API_PORT)
	server.Run(address)
}
