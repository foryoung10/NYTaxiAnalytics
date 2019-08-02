package taxi

import (
	"NYTaxiAnalytics/handlers"
	"encoding/json"
	"fmt"
)

type TaxiJsonRepo struct {
	Client           handlers.Client
	tripsData        string
	averageSpeedData string
	faresData        string
}

func (r TaxiJsonRepo) GetTotalTripsByStartEndDate(startDate string, endDate string) ([]TotalTripsByDay, error) {
	var result []TotalTripsByDay

	if err := json.Unmarshal([]byte(r.tripsData), &result); err != nil {
		return nil, err
	}

	return result, nil

}

func (r TaxiJsonRepo) GetAverageSpeedByDate(date string) ([]AverageSpeedByDay, error) {
	var result []AverageSpeedByDay

	if err := json.Unmarshal([]byte(r.averageSpeedData), &result); err != nil {
		return nil, err
	}

	fmt.Println(result)

	return result, nil

}

func (r TaxiJsonRepo) GetAverageFareByLocation(date string) ([]FarePickupByLocation, error) {
	var result []FarePickupByLocation

	if err := json.Unmarshal([]byte(r.faresData), &result); err != nil {
		return nil, err
	}

	fmt.Println(result)

	return result, nil
}
