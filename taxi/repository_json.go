package taxi

import (
	"encoding/json"

	"github.com/foryoung10/NYTaxiAnalytics/database"
)

// TaxiJsonRepo is a mock repo that returns static Json data from the data file
type TaxiJsonRepo struct {
	Client           database.Client
	TripsData        string
	AverageSpeedData string
	FaresData        string
}

func (r TaxiJsonRepo) GetTotalTripsByStartEndDate(startDate string, endDate string, year int) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	if err := json.Unmarshal([]byte(r.TripsData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

func (r TaxiJsonRepo) GetAverageSpeedByDate(date string, year int) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	if err := json.Unmarshal([]byte(r.AverageSpeedData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

func (r TaxiJsonRepo) GetAverageFareByLocation(date string, year int) ([]FarePickupByLocation, error) {
	var result []FarePickupByLocation

	if err := json.Unmarshal([]byte(r.FaresData), &result); err != nil {
		return nil, err
	}

	return result, nil
}
