// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data.
package taxi

// This file contains the queries for big query.

// Query for location and fares on a date.
var fareLocationQ = `
SELECT
  pickup_longitude , pickup_latitude,  fare_amount
FROM 
	@tables 
WHERE 
  date(pickup_datetime) = @date
`
