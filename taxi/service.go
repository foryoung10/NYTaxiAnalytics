// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

import (
	"github.com/golang/geo/s2"
)

// Service interface to apply business logic to dataset and transforming to the output.
// GetTotalTripsByStartEndDate: Returns the TotalTripsByDay data from the repository.
// GetAverageSpeedByDate: Returns the AverageSpeedByDay data from the repository.
// GetAverageFarePickUpByLocation: Uses the geo s2 library to transform data to the S2idFare result.
type IService interface {
	GetTotalTripsByStartEndDate(string, string, int) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string, int) ([]AverageSpeedByDay, error)
	GetAverageFarePickUpByLocation(string, int, int) ([]S2idFare, error)
}

// Service to apply business logic to dataset and transforming to the output.
type Service struct {
	Repo Repository // Set Repository
}

// GetTotalTripsByStartEndDate: Returns the TotalTripsByDay data from the repository.
// No business logic is applied
func (s Service) GetTotalTripsByStartEndDate(startDate string, endDate string, year int) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	result, _ = s.Repo.GetTotalTripsByStartEndDate(startDate, endDate)

	return result, nil
}

// GetAverageSpeedByDate: Returns the AverageSpeedByDay data from the repository.
// No business logic is applied
func (s Service) GetAverageSpeedByDate(date string, year int) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	result, _ = s.Repo.GetAverageSpeedByDate(date)

	return result, nil
}

// GetAverageFarePickUpByLocation: Uses the geo s2 library to transform data to the S2idFare result.
// Uses s2 library and region coverer to get s2id at level 16 for each location
// Aggregates the location and returns the average fare
func (s Service) GetAverageFarePickUpByLocation(date string, year int, level int) ([]S2idFare, error) {

	var data []FarePickupByLocation
	var fareByLocation []S2idFare

	data, err := s.Repo.GetFareLocationByDate(date, year)

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

			// Return list of s2id and fare
			for j := 0; j < len(cellUnion); j++ {
				fareByLocation = append(fareByLocation, S2idFare{S2id: cellUnion[j].ToToken(), Fare: data[i].Fare})
			}
		}
	}

	// Create fares map to sum the fare for a s2id location.
	fares := make(map[string]float64)
	// Creates count map to count the number of s2id location.
	count := make(map[string]int)

	// Sum fare amount and count by the s2id location.
	for i := 0; i < len(fareByLocation); i++ {
		s2id := fareByLocation[i].S2id
		fare := fareByLocation[i].Fare

		fares[s2id] += fare
		count[s2id] += 1
	}

	var result []S2idFare

	// To get average get sum/count by s2id location
	for k, v := range fares {
		result = append(result, S2idFare{S2id: k, Fare: v / float64(count[k])})
	}

	return result, nil
}
