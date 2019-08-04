// taxi analytics
package taxi

import (
	"github.com/golang/geo/s2"
)

type IService interface {
	GetTotalTripsByStartEndDate(string, string, int) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string, int) ([]AverageSpeedByDay, error)
	GetAverageFarePickUpByLocation(string, int, int) ([]S2idFare, error)
}

type Service struct {
	Repo TaxiRepo
}

// For start and end date, return total number of trips
func (s Service) GetTotalTripsByStartEndDate(startDate string, endDate string, year int) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	result, _ = s.Repo.GetTotalTripsByStartEndDate(startDate, endDate)

	return result, nil
}

// For date, return the average speed
func (s Service) GetAverageSpeedByDate(date string, year int) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	result, _ = s.Repo.GetAverageSpeedByDate(date)

	return result, nil
}

// For date, return average fare level of a location's s2id
// Using s2 library and region coverer to get s2id at level 16
func (s Service) GetAverageFarePickUpByLocation(date string, year int, level int) ([]S2idFare, error) {

	var data []FarePickupByLocation
	var fareByLocation []S2idFare

	data, err := s.Repo.GetAverageFareByLocation(date)

	if data != nil && err == nil {
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
				fareByLocation = append(fareByLocation, S2idFare{S2id: cellUnion[j].ToToken(), Fare: data[i].Fare})
			}
		}
	}

	// to do: improve average code
	fares := make(map[string]float64)
	count := make(map[string]int)

	for i := 0; i < len(fareByLocation); i++ {
		s2id := fareByLocation[i].S2id
		fare := fareByLocation[i].Fare

		fares[s2id] += fare
		count[s2id] += 1
	}

	var result []S2idFare

	for k, v := range fares {
		result = append(result, S2idFare{S2id: k, Fare: v / float64(count[k])})
	}

	return result, nil
}
