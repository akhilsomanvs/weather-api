package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/akhilsomanvs/weather-api/internal/models"
)

const weatherApiUrl = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

func BuildWeatherUrl(location string, apiKey string) (string, error) {
	apiUrl := fmt.Sprintf("%s/%s/today?key=%s", weatherApiUrl, location, apiKey)

	urlBuilder, err := url.Parse(apiUrl)
	if err != nil {
		return "", errors.New("Could not parse URL : " + err.Error())
	}

	return urlBuilder.String(), nil
}

func GetWeatherData(url string) (*models.WeatherApiResponseModel, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Error : " + err.Error())
	}

	queryMap := request.URL.Query()
	queryMap.Add("unitGroup", "metric")
	request.URL.RawQuery = queryMap.Encode()

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("Error : " + err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status code : %v", response.StatusCode)
	}

	var responseModel models.WeatherApiResponseModel
	if err = json.NewDecoder(response.Body).Decode(&responseModel); err != nil {
		return nil, fmt.Errorf("Cannot decode data into model : %s", err.Error())
	}

	return &responseModel, nil
}
