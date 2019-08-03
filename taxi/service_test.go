package taxi

import (
	"NYTaxiAnalytics/database"
	"testing"
)

func TestGetTotalTripsByStartEndDate(t *testing.T) {
	const startDate string = "2015-01-01"
	const endDate string = "2015-02-02"

	t.Log("Running TestGetTotalTripsByStartEndDate")

	client := database.TestClient{}

	for i := 0; i < len(totalTripsData); i++ {

		t.Logf("Testcase - %v", i+1)

		var r TaxiRepo = TaxiJsonRepo{
			Client:    client,
			tripsData: totalTripsData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetTotalTripsByStartEndDate(startDate, endDate)
		expectedResult := totalTripsResult[i]

		resLength := len(res)
		expectedLength := len(expectedResult)

		t.Log("Comparing length of result and expected")
		if resLength != expectedLength {
			t.Errorf("Error with length of received(%v) and expected(%v) object", resLength, expectedLength)
		}

		t.Log("Comparing values of result and expected")
		for j := 0; j < resLength; j++ {
			if res[j] != expectedResult[j] {
				t.Errorf("Error with received %v and expected %v", res[j], expectedResult[j])
			}
		}
	}
}

func TestGetAverageSpeedByDate(t *testing.T) {
	const date string = "2015-01-01"

	t.Log("Running TestGetAverageSpeedByDate")

	client := database.TestClient{}

	for i := 0; i < len(averageSpeedData); i++ {

		t.Logf("Testcase - %v", i+1)

		var r TaxiRepo = TaxiJsonRepo{
			Client:           client,
			averageSpeedData: averageSpeedData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetAverageSpeedByDate(date)
		expectedResult := averageSpeedResult[i]

		resLength := len(res)
		expectedLength := len(expectedResult)

		t.Log("Comparing length of result and expected")
		if resLength != expectedLength {
			t.Errorf("Error with length of received(%v) and expected(%v) object", resLength, expectedLength)
		}

		t.Log("Comparing values of result and expected")
		for j := 0; j < resLength; j++ {
			if res[j] != expectedResult[j] {
				t.Errorf("Error with received %v and expected %v", res[j], expectedResult[j])
			}
		}
	}
}

func TestGetAverageFarePickUpByLocation(t *testing.T) {

	const date string = "2015-01-01"
	const level int = 16

	t.Log("Running TestGetAverageFarePickUpByLocation")

	client := database.TestClient{}

	for i := 0; i < len(faresData); i++ {

		t.Logf("Testcase - %v", i+1)

		var r TaxiRepo = TaxiJsonRepo{
			Client:    client,
			faresData: faresData[i]}

		var s IService = Service{Repo: r}

		res, _ := s.GetAverageFarePickUpByLocation(date, level)
		expectedResult := averageFaresLocationResult[i]

		// convert expectedResult to map

		expectedMap := make(map[string]float64)

		for z := 0; z < len(expectedResult); z++ {
			expectedMap[expectedResult[z].s2id] = expectedResult[z].fare
		}

		t.Log("Comparing values of result and expected")
		for j := 0; j < len(res); j++ {
			id := res[j].s2id
			if res[j].fare != expectedMap[id] {
				t.Errorf("Error with received %v and expected %v", res[j].fare, expectedMap[id])
			}
		}
	}
}
