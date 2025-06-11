package handlers

import (
	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func GetWeatherData(cfg *config.Config) func(cotnext *gin.Context) {
	return func(context *gin.Context) {

		location := context.Param("location")
		if location == "" {
			location = "kerala"
		}
		_, err := controllers.GetWeatherData(location, cfg.ApiKey)

		if err != nil {

		}
	}
}
