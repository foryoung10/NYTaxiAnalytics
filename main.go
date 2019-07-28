package main

import (
	"NYTaxiAnalytics/taxi"
)

func main() {
	// for testing of functions
	taxi.TotalTrips("2015-01-01", "2015-01-02")

	taxi.AverageSpeed("2015-01-01")
	taxi.AverageFareLevel("2015-01-02", 16)
}
