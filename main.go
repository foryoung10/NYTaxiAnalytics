package main

import (
	"NYTaxiAnalytics/database"
	"NYTaxiAnalytics/taxi"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(s taxi.IService) *gin.Engine {

	router := gin.New()

	// Taxi api routes
	router.GET("/total_trips", func(c *gin.Context) {
		startdate := c.Query("start-date")
		enddate := c.Query("end-date")

		c.String(http.StatusOK, "Date %s %s", startdate, enddate)
	})

	router.GET("/average_speed", func(c *gin.Context) {
		date := c.Query("date")

		c.String(http.StatusOK, "Date %s", date)
	})

	router.GET("/average_fare_heatmap", func(c *gin.Context) {
		date := c.Query("date")

		c.String(http.StatusOK, "Date %s", date)

	})

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
		Client: client,
	}
	var s taxi.IService = taxi.Service{Repo: r}

	// Setup router
	router := setupRouter(s)

	// Configure port
	router.Run(":8080")
}
