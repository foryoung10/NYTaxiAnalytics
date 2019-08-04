package taxi

import (
	"NYTaxiAnalytics/database"
	"strings"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const tablePlaceholder string = "@tables"

type TaxiRepo interface {
	GetTotalTripsByStartEndDate(string, string, int) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string, int) ([]AverageSpeedByDay, error)
	GetAverageFareByLocation(string, int) ([]FarePickupByLocation, error)
}

type TaxiBQRepo struct {
	Client database.Client
}

func (r TaxiBQRepo) GetTotalTripsByStartEndDate(startDate string, endDate string, year int) ([]TotalTripsByDay, error) {

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
	tables := getTaxiTables(year)
	query := strings.Replace(totalTripsQ, tablePlaceholder, tables, 1)

	rows, err := r.Client.Query(query, parameters)

	if err != nil {
		return nil, err
	}

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

func (r TaxiBQRepo) GetAverageSpeedByDate(date string, year int) ([]AverageSpeedByDay, error) {
	parameters := []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: date,
		},
	}

	tables := getTaxiTables(year)
	query := strings.Replace(averageSpeedQ, tablePlaceholder, tables, 1)

	rows, err := r.Client.Query(query, parameters)
	var res []AverageSpeedByDay

	if err != nil {
		return nil, err
	}

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

func (r TaxiBQRepo) GetAverageFareByLocation(date string, year int) ([]FarePickupByLocation, error) {
	parameters := []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: date,
		},
	}

	tables := getTaxiTables(year)
	query := strings.Replace(averageFareQ, tablePlaceholder, tables, 1)

	rows, err := r.Client.Query(query, parameters)
	var res []FarePickupByLocation

	if err != nil {
		return nil, err
	}

	if rows.TotalRows > 0 {
		for {
			err := rows.Next(&res)

			if err == iterator.Done {
				break
			}
		}
	}
	return res, nil

}

// returns 2 bigquery taxi tables based on year
func getTaxiTables(year int) string {
	const schema = "`bigquery-public-data.new_york."

	var yellowTripTables = make(map[int]string)
	var greenTripTables = make(map[int]string)

	yellowTripTables[2015] = "tlc_yellow_trips_2015`"
	yellowTripTables[2016] = "tlc_yellow_trips_2016`"
	yellowTripTables[2017] = "tlc_yellow_trips_2017`"

	greenTripTables[2014] = "tlc_green_trips_2014`"
	greenTripTables[2015] = "tlc_green_trips_2015`"
	greenTripTables[2016] = "tlc_green_trips_2016`"
	greenTripTables[2017] = "tlc_green_trips_2017`"

	_, okYellow := yellowTripTables[year]
	_, okGreen := greenTripTables[year]
	tables := ""
	if okYellow {
		tables = schema + yellowTripTables[year] + "\n"
	}
	if okGreen {
		tables = tables + schema + greenTripTables[year]
	}

	return tables
}
