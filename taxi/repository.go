package taxi

import (
	"NYTaxiAnalytics/database"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type TaxiRepo interface {
	GetTotalTripsByStartEndDate(string, string) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string) ([]AverageSpeedByDay, error)
	GetAverageFareByLocation(string) ([]FarePickupByLocation, error)
}

type TaxiBQRepo struct {
	client database.Client
}

func (r TaxiBQRepo) GetTotalTripsByStartEndDate(startDate string, endDate string) ([]TotalTripsByDay, error) {

	parameters := []bigquery.QueryParameter{
		{
			Name:  "startDate",
			Value: startDate,
		},
		{
			Name:  "endDate",
			Value: endDate,
		},
	}

	var res []TotalTripsByDay

	rows, _ := r.client.Query(totalTripsQ, parameters)

	for {
		err := rows.Next(&res)

		if err == iterator.Done {
			break
		}
		if err != nil {

		}
	}
	return res, nil
}

func (r TaxiBQRepo) GetAverageSpeedByDate(date string) ([]AverageSpeedByDay, error) {
	parameters := []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: date,
		},
	}
	rows, _ := r.client.Query(averageSpeedQ, parameters)
	var res []AverageSpeedByDay

	for {
		err := rows.Next(&res)

		if err == iterator.Done {
			break
		}
		if err != nil {

		}
	}
	return res, nil
}

func (r TaxiBQRepo) GetAverageFareByLocation(date string) ([]FarePickupByLocation, error) {
	parameters := []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: date,
		},
	}
	rows, _ := r.client.Query(averageFareQ, parameters)
	var res []FarePickupByLocation

	for {
		err := rows.Next(&res)

		if err == iterator.Done {
			break
		}
		if err != nil {

		}
	}
	return res, nil
}
