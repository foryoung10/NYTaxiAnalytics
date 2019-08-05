// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

// This file contains the queries for big query.

// Query for total trips for a start date and end date.
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

// Query for average speed for a date.
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

// Query for location and fares on a date.
var fareLocationQ = `
SELECT
  pickup_longitude , pickup_latitude,  fare_amount
FROM 
	@tables 
WHERE 
  date(pickup_datetime) = @date
`
