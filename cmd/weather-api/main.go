package main

import (
	"log"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/routes"
	"github.com/akhilsomanvs/weather-api/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.MustLoad()
	database, err := db.InitCache(cfg)
	if err != nil {
		log.Printf("DB : %s", err.Error())
	}

	//Server setup
	server := gin.Default()
	routes.RegisterRoutes(server, cfg, database)
	server.Run(cfg.GetHostAddress())
}
