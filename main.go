// Main package to setup dependencies and run the server.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/foryoung10/NYTaxiAnalytics/database"
	"github.com/foryoung10/NYTaxiAnalytics/taxi"
	"github.com/gin-gonic/gin"
)

// Setup logger, gin router and api routes.
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

	// Setup gin routers
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.LstdFlags)

	// Setup api routes
	router.GET("/total_trips", hand.FetchTotalTrips())
	router.GET("/average_speed", hand.FetchAverageSpeed())
	router.GET("/average_fare_heatmap", hand.FetchAverageFareS2id())

	return router
}

// Dev router uses json stored in data file.
func useDevRepo() taxi.Repository {
	client := database.TestClient{}
	var r taxi.Repository = taxi.JsonRepository{
		Client:           client,
		FaresData:        taxi.FaresData[0],
		AverageSpeedData: taxi.AverageSpeedData[0],
		TripsData:        taxi.TotalTripsData[0],
	}

	return r
}

// BQ repo connects and pulls data from Bigquery api.
func useBQRepo() taxi.Repository {
	// Setup client, taxi repo, service
	client := database.BigQueryClient{}
	var r taxi.Repository = taxi.BqRepository{
		Client: client,
	}
	return r
}

func main() {
	// Setup big query client
	// database.MigrateDataFromBigQuery()

	database.QueryDb()

	database.BigQueryClientSetup()

	time1, _ := time.Parse(time.RFC3339, "2014-06-01 05:22:37 UTC")

	fmt.Println(time1)

	// Setup client, taxi repo, service

	// For dev
	// var r = useDevRepo()

	// For connecting to Bigquery
	var r = useBQRepo()

	var s = taxi.Service{Repo: r}
	var hand = taxi.Handler{Svc: s}

	// Setup router
	router := setupRouter(hand)

	// Configure port
	router.Run(":8080")
}
