package handlers

import (
	"net/http"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/controllers"
	"github.com/akhilsomanvs/weather-api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetWeatherData(cfg *config.Config) func(cotnext *gin.Context) {
	return func(context *gin.Context) {

		location := context.Param("location")
		if location == "" {
			location = "kerala"
		}
		response, err := controllers.GetWeatherData(location, cfg.ApiKey)

		if err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Could not fetch data : "+err.Error()))
			return
		}
		context.JSON(http.StatusOK, models.NewApiResponseModel("Success", response))

	}
}
