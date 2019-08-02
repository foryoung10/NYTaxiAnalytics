// taxi analytics
package taxi

import (
	"github.com/golang/geo/s2"
)

type IService interface {
	GetTotalTripsByStartEndDate(string, string) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string) ([]AverageSpeedByDay, error)
	GetAverageFarePickUpByLocation(string, int) ([]s2idFare, error)
}

type Service struct {
	Repo TaxiRepo
}

// For start and end date, return total number of trips
func (s Service) GetTotalTripsByStartEndDate(startDate string, endDate string) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	result, _ = s.Repo.GetTotalTripsByStartEndDate(startDate, endDate)

	return result, nil
}

// For date, return the average speed
func (s Service) GetAverageSpeedByDate(date string) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	result, _ = s.Repo.GetAverageSpeedByDate(date)

	return result, nil
}

// For date, return average fare level of a location's s2id
// Using s2 library and region coverer to get s2id at level 16

func (s Service) GetAverageFarePickUpByLocation(date string, level int) ([]s2idFare, error) {

	// client.Query("averageFareQuery")
	var data []FarePickupByLocation
	var fareByLocation []s2idFare

	data, _ = s.Repo.GetAverageFareByLocation(date)

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

	}
	return fareByLocation, nil
}
