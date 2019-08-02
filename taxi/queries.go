package taxi

// This file contains the queries for big query

// List of tables to query
var tables = "`bigquery-public-data.new_york.tlc_yellow_trips_2015` \n" +
	"`bigquery-public-data.new_york.tlc_green_trips_2015` \n"

// Query for total trips
var totalTripsQ = `
SELECT
	DATE(pickup_datetime) date,
	COUNT(*) total_trips
FROM ` + tables +
	`
WHERE
	pickup_datetime between @startDate and @endDate
GROUP BY
	date
ORDER BY
	date
LIMIT 100`

// Query for average speed
var averageSpeedQ = `
SELECT
	ROUND(AVG(trip_distance/TIMESTAMP_DIFF(TIMESTAMP(dropoff_datetime), TIMESTAMP(pickup_datetime), SECOND)) * 3600, 1) average_speed
FROM` +
	tables +
	` 
WHERE
	trip_distance > 0
AND
	DATE(dropoff_datetime) = @date
AND
	dropoff_datetime > pickup_datetime 
`

// Query for average fares
var averageFareQ = `
SELECT
  pickup_longitude , pickup_latitude,  fare_amount
FROM ` +
	tables + ` 
WHERE 
  date(pickup_datetime) = @date
`
