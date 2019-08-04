package main

import (
	"NYTaxiAnalytics/database"
	"NYTaxiAnalytics/taxi"

	"github.com/gin-gonic/gin"
)

func setupRouter(hand taxi.IHandler) *gin.Engine {

	router := gin.New()

	// Taxi api routes
	router.GET("/total_trips", hand.FetchTotalTrips())
	router.GET("/average_speed", hand.FetchAverageSpeed())
	router.GET("/average_fare_heatmap", hand.FetchAverageFareS2id())

	return router
}

func main() {
	// Setup big query client
	database.BigQueryClientSetup()

	// Setup client, taxi repo, service
	// Use mock repo for development
	// TO DO: switch to bigquery client when done
	client := database.TestClient{}
	var r taxi.TaxiRepo = taxi.TaxiJsonRepo{
		Client:           client,
		FaresData:        taxi.FaresData[0],
		TripsData:        taxi.TotalTripsData[0],
		AverageSpeedData: taxi.AverageSpeedData[0],
	}
	var s = taxi.Service{Repo: r}

	var hand = taxi.Handler{Svc: s}

	// Setup router
	router := setupRouter(hand)

	// Configure port
	router.Run(":8080")

}
