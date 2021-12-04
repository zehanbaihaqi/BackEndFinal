package main

import (
	"sim/configs"
	"sim/routes"
)

func main() {
	configs.ConnectDB()
	e := routes.RoutePemberitahuan()
	e.Start(":8000")
}
