// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

import (
	"log"
	"strings"
	
	"github.com/foryoung10/NYTaxiAnalytics/database"
	
	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const tablePlaceholder string = "@tables"
const tableName string = "taxis"

// Repository handles data transfer between application and database.
// GetTotalTripsByStartEndDate: Gets trips data for a start date and end date from the database and converts to TotalTripsByDay array.
// GetAverageSpeedByDate: Gets average speed data for a date from the database and converts to AverageSpeedByDay array.
// GetFareLocationByDate: Get fares and location for a date from the database and converts to FarePickupByLocation array.
type Repository interface {
	GetTotalTripsByStartEndDate(string, string) ([]TotalTripsByDay, error)
	GetAverageSpeedByDate(string) ([]AverageSpeedByDay, error)
	GetFareLocationByDate(string, int) ([]FarePickupByLocation, error)
}

// BQRepo queries and retrieves data from big query database.
// The repo handles the setting of Big query parameters
// The repo handles the query generation from the query file
// The repo coverts the raw data to a struct
type DataRepository struct {
	BqClient database.BqClient
	DbClient database.TaxiConn
}


// GetTotalTripsByStartEndDate: Gets trips data for a start date and end date from the database and converts to TotalTripsByDay array.
func (r DataRepository) GetTotalTripsByStartEndDate(startDate string, endDate string) ([]TotalTripsByDay, error) {

	var trips []TotalTripsByDay
	db := r.DbClient.Db
	db.Table(tableName).Select("date, sum(total_trips) as total_trips").Where("date BETWEEN ? AND ?", startDate, endDate).Group("date").Scan(&trips)

	return trips, nil
}

// GetAverageSpeedByDate: Gets average speed data for a date from the database and converts to AverageSpeedByDay array.
func (r DataRepository) GetAverageSpeedByDate(date string) ([]AverageSpeedByDay, error) {

	var avgSpeed []AverageSpeedByDay
	db := r.DbClient.Db
	db.Table(tableName).Select("round(avg(average_speed),1) as average_speed").Where("date = ?", date).Scan(&avgSpeed)
	
	return avgSpeed, nil
}

// GetFareLocationByDate: Get fares and location for a date from the database and converts to FarePickupByLocation array.
func (r DataRepository) GetFareLocationByDate(date string, year int) ([]FarePickupByLocation, error) {
	// Set big query parameters
	parameters := []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: date,
		},
	}

	// Generating query
	tables := getTaxiTables(year)
	query := strings.Replace(fareLocationQ, tablePlaceholder, tables, 1)

	rows, err := r.BqClient.Query(query, parameters)
	var res []FarePickupByLocation

	log.Println(rows.TotalRows)

	if err != nil {
		return nil, err
	}

	// Read and converts to FarePickupByLocation array
	for {
		var tmp FarePickupByLocation
		err := rows.Next(&tmp)

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		res = append(res, tmp)
	}

	return res, nil

}

// Gets the table names based on the year
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
