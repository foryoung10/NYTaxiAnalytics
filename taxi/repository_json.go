// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

import (
	"encoding/json"

	"github.com/foryoung10/NYTaxiAnalytics/database"
)

// JsonRepository is a mock repo that returns static Json data from the data file
type JsonRepository struct {
	Client           database.Client
	TripsData        string
	AverageSpeedData string
	FaresData        string
}

func (r JsonRepository) GetTotalTripsByStartEndDate(startDate string, endDate string, year int) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	if err := json.Unmarshal([]byte(r.TripsData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

func (r JsonRepository) GetAverageSpeedByDate(date string, year int) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	if err := json.Unmarshal([]byte(r.AverageSpeedData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

func (r JsonRepository) GetFareLocationByDate(date string, year int) ([]FarePickupByLocation, error) {
	var result []FarePickupByLocation

	if err := json.Unmarshal([]byte(r.FaresData), &result); err != nil {
		return nil, err
	}

	return result, nil
}
