// taxi models for data returned from big query
package taxi

// average speed of all trips on a day
type AverageSpeedByDay struct {
	AverageSpeed float64 `json:"average_speed"`
}

// total number trips on a day
type TotalTripsByDay struct {
	Date       string
	TotalTrips int `json:"total_trips,string"`
}

// pick up location and fare price
type FarePickupByLocation struct {
	Lng  float64 `json:"pickup_longitude,string"`
	Lat  float64 `json:"pickup_latitude,string"`
	Fare float64 `json:"fare_amount,string"`
}

type s2idFare struct {
	s2id string
	fare float64
}
