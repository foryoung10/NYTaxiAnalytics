package database

import (
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/api/iterator"
)

// Migration script for DB
func MigrateDataFromBigQuery() {
	db := Connect()
	db.AutoMigrate(&Taxi{})

	const schema = "`bigquery-public-data.new_york_taxi_trips."
	client := BigQueryClient{}

	parameters := []bigquery.QueryParameter{}
	var colour = ""

	for _, table := range tripsTable {

		if strings.Contains(table, "yellow") {
			colour = "yellow"
		} else {
			colour = "green"
		}
	
		q := strings.Replace(query, "@table", schema+table, 1)
		rows, err := client.Query(q, parameters)

		fmt.Println(query)

		if err != nil {
		}

		for {
			var tmp Taxi
			err := rows.Next(&tmp)

			if err == iterator.Done {
				break
			}
			if err != nil {
			}

			rec := Taxi{
				Date:          tmp.Date,
				Total_trips:   tmp.Total_trips,
				Average_speed: tmp.Average_speed,
				Type:          colour,
			}
			db.Create(&rec)
		}
	}
}