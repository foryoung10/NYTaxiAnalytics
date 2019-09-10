package database

type Taxi struct {
	// gorm.model
	Date          string  `bigquery:"date"`
	Total_trips   int     `bigquery:"total_trips"`
	Average_speed float64 `bigquery:"average_speed"`
	Type          string
}

var tripsTable = []string{
	"tlc_yellow_trips_2015`",
	"tlc_yellow_trips_2016`",
	"tlc_yellow_trips_2017`",
	"tlc_green_trips_2014`",
	"tlc_green_trips_2015`",
	"tlc_green_trips_2016`",
	"tlc_green_trips_2017`",
}

var query = `SELECT CAST(DATE(pickup_datetime) as string) date, count(*) total_trips,        
AVG(trip_distance / TIMESTAMP_DIFF(timestamp(cast(dropoff_datetime as string)), timestamp(cast(pickup_datetime as string)), SECOND)) * 3600 average_speed
FROM @table
WHERE trip_distance is not null and trip_distance > 0 and dropoff_datetime is not null and pickup_datetime is not null
and dropoff_datetime > pickup_datetime
GROUP BY date
order by date
`

/*
var query = `  SELECT FORMAT_TIMESTAMP("%Y-%m-%dT%X%Ez", cast(pickup_datetime as timestamp)) as pickup_datetime,
FORMAT_TIMESTAMP("%Y-%m-%dT%X%Ez", cast(dropoff_datetime as timestamp)) as dropoff_datetime,
 pickup_latitude, pickup_longitude, fare_amount, trip_distance
 FROM @table
where dropoff_datetime is not null
and pickup_datetime is not null and pickup_latitude is not null and pickup_longitude is not null
and trip_distance is not null and fare_amount is not null and fare_amount > 0
and trip_distance > 0 and pickup_latitude > 0 and pickup_longitude > 0
`
*/
