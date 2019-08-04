// taxi models for data returned from big query
package taxi

// average speed of all trips on a day
type AverageSpeedByDay struct {
	AverageSpeed float64 `bigquery:"average_speed"`
}

// total number trips on a day
type TotalTripsByDay struct {
	Date       string `json:"date"`
	TotalTrips int    `bigquery:"total_trips"`
}

//Date       string `bigquery:"date"`
//TotalTrips int64  `bigquery:"total_trips"`

// pick up location and fare price
type FarePickupByLocation struct {
	Lng  float64 `bigquery:"pickup_longitude"`
	Lat  float64 `bigquery:"pickup_latitude"`
	Fare float64 `bigquery:"fare_amount"`
}

type S2idFare struct {
	S2id string  `json:"s2id"`
	Fare float64 `json:"fare"`
}
