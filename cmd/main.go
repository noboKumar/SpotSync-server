package main

import (
	"github.com/noboKumar/SpotSync-server/config"
	"github.com/noboKumar/SpotSync-server/server"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabase(cfg)
	server.Start(cfg, db)
}
