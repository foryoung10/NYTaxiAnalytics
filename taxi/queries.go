package taxi

// This file contains the queries for big query

// Query for total trips
var totalTripsQ = `
SELECT
	CAST(DATE(pickup_datetime) as string) date,
	COUNT(*) total_trips
FROM 
	@tables 
WHERE
	pickup_datetime between @startDate and @endDate
GROUP BY
	date
ORDER BY
	date
`

// Query for average speed

var averageSpeedQ = `
SELECT
	ROUND(AVG(trip_distance / TIMESTAMP_DIFF(dropoff_datetime, 
		pickup_datetime, 
		SECOND)) * 3600, 1) average_speed
FROM 
	@tables 
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
FROM 
	@tables 
WHERE 
  date(pickup_datetime) = @date
`
