// taxi analytics
package taxi

import (
	"encoding/json"
	"fmt"

	"github.com/golang/geo/s2"
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
// Using s2 library and region coverer to get s2id at level 16

func AverageFareLevel(date string, level int) {

	var data []farePickupLocation
	var fareByLocation []s2idFare

	if err := json.Unmarshal([]byte(faresData), &data); err != nil {
		panic(err)
	}

	for i := 0; i < len(data); i++ {

		// Get latlng and convert to point
		latlng := s2.LatLngFromDegrees(data[i].Lat, data[i].Lng).Normalized()
		point := s2.PointFromLatLng(latlng)

		// Create cap from point
		// Region from cap
		cap := s2.CapFromPoint(point)
		region := s2.Region(cap)

		// Using region coverer set the level
		// Use covering to get list of cellids
		rc := &s2.RegionCoverer{MaxLevel: level, MinLevel: level}
		cellUnion := rc.Covering(region)

		// Return list of CellId and fare
		for j := 0; j < len(cellUnion); j++ {
			fareByLocation = append(fareByLocation, s2idFare{s2id: cellUnion[j].ToToken(), fare: data[i].Fare})
		}

		fmt.Println(fareByLocation)

	}

}
