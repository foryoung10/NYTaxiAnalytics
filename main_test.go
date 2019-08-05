package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foryoung10/NYTaxiAnalytics/database"
	"github.com/foryoung10/NYTaxiAnalytics/taxi"

	"github.com/magiconair/properties/assert"
)

func performRequest(r http.Handler, method, path string, param map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)

	q := req.URL.Query()
	for k, v := range param {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func setUpTestRouterDeps() taxi.IHandler {
	client := database.TestClient{}
	var r taxi.TaxiRepo = taxi.TaxiJsonRepo{
		Client:           client,
		FaresData:        taxi.FaresData[0],
		TripsData:        taxi.TotalTripsData[0],
		AverageSpeedData: taxi.AverageSpeedData[0],
	}
	var s = taxi.Service{Repo: r}

	var handler = taxi.Handler{Svc: s}

	return handler
}

func TestFetchTotalTrips(t *testing.T) {

	// setup test cases
	dates := []struct {
		inputStartDate  string
		inputEndDate    string
		expectedResCode int
	}{
		{"2015-01-01", "2015-01-10", http.StatusOK},
		{"2019-01-10", "2019-05-01", http.StatusNoContent},
		{"2016-06-31", "2016-05-01", http.StatusNotFound},
	}
	// setup Router for test
	var handler = setUpTestRouterDeps()

	for _, date := range dates {

		router := setupRouter(handler)
		var params = make(map[string]string)

		params["start-date"] = date.inputStartDate
		params["end-date"] = date.inputEndDate

		w := performRequest(router, "GET", "/total_trips", params)

		assert.Equal(t, w.Code, date.expectedResCode)

		// test the result output correctly in json format
		if w.Code == http.StatusOK {

			expectedResult := []taxi.TotalTripsByDay{
				{
					Date:       "2015-01-01",
					TotalTrips: 382014,
				},
				{
					Date:       "2015-01-02",
					TotalTrips: 345296,
				},
				{
					Date:       "2015-01-03",
					TotalTrips: 406769,
				},
			}

			var response []taxi.TotalTripsByDay
			json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Equal(t, response, expectedResult)

		}
	}
}

func TestFetchAverageSpeed(t *testing.T) {

	// setup test cases
	dates := []struct {
		inputDate       string
		expectedResCode int
	}{
		{"2015-01-01", http.StatusOK},
		{"2019-01-10", http.StatusNoContent},
		{"10-10-2015", http.StatusNotFound},
	}
	// setup Router for test
	var handler = setUpTestRouterDeps()

	for _, date := range dates {

		router := setupRouter(handler)
		var params = make(map[string]string)

		params["date"] = date.inputDate

		w := performRequest(router, "GET", "/average_speed", params)

		assert.Equal(t, w.Code, date.expectedResCode)

		// test the result output correctly in json format
		if w.Code == http.StatusOK {

			expectedResult := []taxi.AverageSpeedByDay{
				{
					AverageSpeed: 14.1,
				},
			}

			var response []taxi.AverageSpeedByDay
			json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Equal(t, response, expectedResult)
		}
	}
}

func TestFetchFetchAverageFareS2id(t *testing.T) {

	// setup test cases
	dates := []struct {
		inputDate       string
		expectedResCode int
	}{
		{"2015-01-01", http.StatusOK},
		{"2019-01-10", http.StatusNoContent},
		{"10-10-2015", http.StatusNotFound},
	}
	// setup Router for test
	var handler = setUpTestRouterDeps()

	for _, date := range dates {

		router := setupRouter(handler)
		var params = make(map[string]string)

		params["date"] = date.inputDate

		w := performRequest(router, "GET", "/average_fare_heatmap", params)

		assert.Equal(t, w.Code, date.expectedResCode)

		// test the result output correctly in json format
		if w.Code == http.StatusOK {

			expectedResult := []taxi.S2idFare{
				{
					S2id: "89c25a3a1",
					Fare: 27.0,
				},
				{
					S2id: "89c2f5dd3",
					Fare: 0.0,
				},
				{
					S2id: "951977d37",
					Fare: 10.0,
				},
				{
					S2id: "951978321",
					Fare: 10.0,
				},
			}

			// make map for testing
			expectedMap := make(map[string]float64)
			for z := 0; z < len(expectedResult); z++ {
				expectedMap[expectedResult[z].S2id] = expectedResult[z].Fare
			}

			// use map for comparison
			var response []taxi.S2idFare
			json.Unmarshal([]byte(w.Body.String()), &response)

			for j := 0; j < len(response); j++ {
				id := response[j].S2id
				assert.Equal(t, response[j].Fare, expectedMap[id])
			}
		}
	}
}
