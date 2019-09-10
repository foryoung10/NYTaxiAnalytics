package database

import (
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/api/iterator"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "TaxiData"
)

func MigrateDataFromBigQuery() {
	fmt.Println("Migrating from bigquery")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Taxi{})

	if err != nil {
		panic(err)
	}

	//	getQuery()

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
		fmt.Println(colour)

		q := strings.Replace(query, "@table", schema+table, 1)
		rows, err := client.Query(q, parameters)

		fmt.Println(query)

		log.Println(rows.TotalRows)

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

func QueryDb() {
	var t Taxi

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := gorm.Open("postgres", psqlInfo)

	db.Select("date, sum(total_trips) as total_trips").Group("date").Find(&t)

	fmt.Println(t)

}
