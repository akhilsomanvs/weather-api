package routes

import (
	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/handlers"
	"github.com/akhilsomanvs/weather-api/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, cfg *config.Config, database *db.Database) {
	server.GET("/weather/:location", handlers.GetWeatherData(cfg, database))
}
