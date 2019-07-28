// taxi analytics
package taxi

import (
	"encoding/json"
	"fmt"
)

// For start and end date, return total number of trips
func TotalTrips(startDate string, endDate string) {
	var result []totalTripsDay

	if err := json.Unmarshal([]byte(totalTripsData), &result); err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// For date, return the average speed
func AverageSpeed(date string) {
	var result []averageSpeedDay

	if err := json.Unmarshal([]byte(averageSpeedData), &result); err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// For date, return average fare level of a location's s2id
func AverageFareLevel(date string, level int) {

	var result []farePickupLocation

	if err := json.Unmarshal([]byte(faresData), &result); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
