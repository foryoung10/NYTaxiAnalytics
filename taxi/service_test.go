package taxi

import (
	"testing"

	"github.com/foryoung10/NYTaxiAnalytics/database"
)

func TestGetTotalTripsByStartEndDate(t *testing.T) {
	const startDate string = "2015-01-01"
	const endDate string = "2015-02-02"
	const year int = 2015

	client := database.TestClient{}

	for i := 0; i < len(TotalTripsData); i++ {

		var r TaxiRepo = TaxiJsonRepo{
			Client:    client,
			TripsData: TotalTripsData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetTotalTripsByStartEndDate(startDate, endDate, year)
		expectedResult := totalTripsResult[i]

		resLength := len(res)
		expectedLength := len(expectedResult)

		if resLength != expectedLength {
			t.Errorf("Error with length of received(%v) and expected(%v) object", resLength, expectedLength)
		}

		for j := 0; j < resLength; j++ {
			if res[j] != expectedResult[j] {
				t.Errorf("Error with received %v and expected %v", res[j], expectedResult[j])
			}
		}
	}
}

func TestGetAverageSpeedByDate(t *testing.T) {
	const date string = "2015-01-01"
	const year int = 2015

	client := database.TestClient{}

	for i := 0; i < len(AverageSpeedData); i++ {

		var r TaxiRepo = TaxiJsonRepo{
			Client:           client,
			AverageSpeedData: AverageSpeedData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetAverageSpeedByDate(date, year)
		expectedResult := averageSpeedResult[i]

		resLength := len(res)
		expectedLength := len(expectedResult)

		if resLength != expectedLength {
			t.Errorf("Error with length of received(%v) and expected(%v) object", resLength, expectedLength)
		}

		for j := 0; j < resLength; j++ {
			if res[j] != expectedResult[j] {
				t.Errorf("Error with received %v and expected %v", res[j], expectedResult[j])
			}
		}
	}
}

func TestGetAverageFarePickUpByLocation(t *testing.T) {

	const date string = "2015-01-01"
	const year int = 2015
	const level int = 16

	client := database.TestClient{}

	for i := 0; i < len(FaresData); i++ {

		var r TaxiRepo = TaxiJsonRepo{
			Client:    client,
			FaresData: FaresData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetAverageFarePickUpByLocation(date, year, level)
		expectedResult := averageFaresLocationResult[i]

		// convert expectedResult to map
		expectedMap := make(map[string]float64)

		for z := 0; z < len(expectedResult); z++ {
			expectedMap[expectedResult[z].S2id] = expectedResult[z].Fare
		}

		// use map for comparison
		for j := 0; j < len(res); j++ {
			id := res[j].S2id
			if res[j].Fare != expectedMap[id] {
				t.Errorf("Error with received %v and expected %v", res[j].Fare, expectedMap[id])
			}
		}
	}
}
