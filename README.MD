# Weather API

Weather API created with Golang and uses Redis for Caching

- The project create based on the [requirements](https://roadmap.sh/projects/weather-api-wrapper-service)
- The picture below explains the basic requirements of the project.
  <p align="center" width="100%" >
    <img width="100%" src="https://assets.roadmap.sh/guest/weather-api-f8i1q.png" > 
</p>


## Instructions to run the project

- This project uses [Visual Crossing](https://www.visualcrossing.com/) for fetching weather data
    - Sign up and get an API key from Visual crossing.
    - Copy and paste the API key in [local.yaml](config/local.yaml) file against "API_KEY" field
- Install [Golang](https://go.dev/doc/install)
- Install [Redis Open Source](https://redis.io/downloads/)
- Clone the repository
- Start the Redis Server by running the following command

```console
redis-server
```

- Go to your project folder
- run the command

```console
go run cmd/weather-api/main.go --config config/local.yaml
```
