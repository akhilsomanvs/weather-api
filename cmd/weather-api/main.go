package main

import "github.com/akhilsomanvs/weather-api/internal/config"

func main() {

	_ = config.MustLoad()
}
