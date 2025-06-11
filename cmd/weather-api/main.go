package main

import (
	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.MustLoad()

	//Server setup
	server := gin.Default()
	routes.RegisterRoutes(server, cfg)
	server.Run(cfg.GetHostAddress())
}
