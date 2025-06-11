package routes

import (
	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, cfg *config.Config) {
	server.GET("/weather/:location", handlers.GetWeatherData(cfg))
}
