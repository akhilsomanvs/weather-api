package handlers

import (
	"log"
	"net/http"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/controllers"
	"github.com/akhilsomanvs/weather-api/internal/models"
	"github.com/akhilsomanvs/weather-api/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func GetWeatherData(cfg *config.Config, database *db.Database) func(cotnext *gin.Context) {
	return func(context *gin.Context) {

		location := context.Param("location")
		if location == "" {
			location = "kerala"
		}

		weatherUrl, err := controllers.BuildWeatherUrl(location, cfg.ApiKey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Could not fetch data : "+err.Error()))
			return
		}
		weatherData, err := database.GetWeatherData(weatherUrl)
		if err != nil {
			//Fetch data from remote API
			log.Printf("Data does not exist in cache: %s\n", err.Error())

			response, err := controllers.GetWeatherData(weatherUrl)

			if err != nil {
				context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Could not fetch data : "+err.Error()))
				return
			}

			//Save the data to cache
			database.SetWeatherData(weatherUrl, *response)

			context.JSON(http.StatusOK, models.NewApiResponseModel("Success", response))
		} else {
			//Return data from cache
			log.Printf("Data fetched from cache")
			context.JSON(http.StatusOK, models.NewApiResponseModel("Success", weatherData))
		}

	}
}
