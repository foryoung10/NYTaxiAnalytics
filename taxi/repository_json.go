// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

import (
	"encoding/json"

	"github.com/foryoung10/NYTaxiAnalytics/database"
)

// JsonRepository is a mock repo that returns static Json data from the data file
type JsonRepository struct {
	Client           database.BqClient // Set Client
	TripsData        string          // Set Trips data
	AverageSpeedData string          // Set Average speed data
	FaresData        string          // Set Fare and location data
}

// GetTotalTripsByStartEndDate: Gets trips data for a start date and end date from the database and converts to TotalTripsByDay array.
func (r JsonRepository) GetTotalTripsByStartEndDate(startDate string, endDate string) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	if err := json.Unmarshal([]byte(r.TripsData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

// GetAverageSpeedByDate: Gets average speed data for a date from the database and converts to AverageSpeedByDay array.
func (r JsonRepository) GetAverageSpeedByDate(date string) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	if err := json.Unmarshal([]byte(r.AverageSpeedData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

// GetFareLocationByDate: Get fares and location for a date from the database and converts to FarePickupByLocation array.
func (r JsonRepository) GetFareLocationByDate(date string, year int) ([]FarePickupByLocation, error) {
	var result []FarePickupByLocation

	if err := json.Unmarshal([]byte(r.FaresData), &result); err != nil {
		return nil, err
	}

	return result, nil
}
