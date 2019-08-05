package main

import (
	"io"
	"log"
	"os"

	"github.com/foryoung10/NYTaxiAnalytics/database"
	"github.com/foryoung10/NYTaxiAnalytics/taxi"
	"github.com/gin-gonic/gin"
)

func setupRouter(hand taxi.IHandler) *gin.Engine {

	router := gin.New()

	// Setup logger
	router.Use(gin.Logger())

	filePath := "log/gin.log"
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.LstdFlags)

	// Taxi api routes
	router.GET("/total_trips", hand.FetchTotalTrips())
	router.GET("/average_speed", hand.FetchAverageSpeed())
	router.GET("/average_fare_heatmap", hand.FetchAverageFareS2id())

	return router
}

// dev router uses json stored in data file
func useDevRepo() taxi.TaxiRepo {
	client := database.TestClient{}
	var r taxi.TaxiRepo = taxi.TaxiJsonRepo{
		Client:           client,
		FaresData:        taxi.FaresData[0],
		AverageSpeedData: taxi.AverageSpeedData[0],
		TripsData:        taxi.TotalTripsData[0],
	}

	return r
}

// BQ repo connects and pulls data from BQ
func useBQRepo() taxi.TaxiRepo {
	// Setup client, taxi repo, service
	client := database.BigQueryClient{}
	var r taxi.TaxiRepo = taxi.TaxiBQRepo{
		Client: client,
	}
	return r
}

func main() {
	// Initialise logger
	// Setup big query client
	database.BigQueryClientSetup()

	// Setup client, taxi repo, service

	// for dev
	// var r = useDevRepo()

	// for actual use
	var r = useBQRepo()

	var s = taxi.Service{Repo: r}
	var hand = taxi.Handler{Svc: s}

	// Setup router
	router := setupRouter(hand)

	// Configure port
	router.Run(":8080")
}
