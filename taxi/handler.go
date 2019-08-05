// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data
package taxi

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler handles and validates requests from the api.
// FetchTotalTrips: Handles request to get total trips
// FetchAverageSpeed: Handles request to get average speed
// FetchAverageFaresS2id: Handles request to get average fare for s2id.
type IHandler interface {
	FetchTotalTrips() func(c *gin.Context)
	FetchAverageSpeed() func(c *gin.Context)
	FetchAverageFareS2id() func(c *gin.Context)
}

const maxYear = 2017        // Max year of available data
const minYear = 2014        // Min year of available data
const format = "2006-01-02" // Format of date string: "YYYY-MM-DD"

// Handler handles and validates requests from the api.
// Uses service interface
type Handler struct {
	Svc IService
}

// Fetch total trips on each day for a start date and end date.
// Validation: Check that start date must be before the end date.
// Validation: Data is only available from minYear to maxYear.
// Validation: Date must be in the correct format.
func (hand Handler) FetchTotalTrips() func(c *gin.Context) {
	return func(c *gin.Context) {
		startDate := c.Query("start-date")
		endDate := c.Query("end-date")

		year, err := getYearValidateStartEndDate(startDate, endDate)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		if year > maxYear || year < minYear {
			c.JSON(http.StatusNoContent, "")
			return
		}

		res, err := hand.Svc.GetTotalTripsByStartEndDate(startDate, endDate, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}
}

// Fetch average speed for the last 24 hours (current date -1).
// Validation: Data is only available from minYear to maxYear.
// Validation: Date must be in the correct format.
func (hand Handler) FetchAverageSpeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Query("date")

		// Get previous date and year
		prevDate, year, err := getPreviousDateAndYearValidateDate(date)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		if year > maxYear || year < minYear {
			c.JSON(http.StatusNoContent, "")
			return
		}

		res, err := hand.Svc.GetAverageSpeedByDate(prevDate, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}
}

// Fetch average fare amount of a location(s2id) for a date.
// Validation: Data is only available from minYear to maxYear.
// Validation: Date must be in the correct format.
func (hand Handler) FetchAverageFareS2id() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Query("date")
		const level = 16

		year, err := getYearValidateDate(date)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		if year > maxYear || year < minYear {
			c.JSON(http.StatusNoContent, "")
			return
		}

		res, err := hand.Svc.GetAverageFarePickUpByLocation(date, year, 16)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}
}

// Validator: Get the previous date and year from the input date.
// Date should be in correct format
func getPreviousDateAndYearValidateDate(date string) (string, int, error) {
	time, err := time.Parse(format, date)

	if err != nil {
		return "", 0, errors.New("Date should be in YYYY-MM-DD format")
	}

	previousDate := time.AddDate(0, 0, -1)
	year := previousDate.Year()

	return previousDate.Format(format), year, nil
}

// Validator: Get the year from the input date.
// Date should be in correct format
func getYearValidateDate(date string) (int, error) {
	time, err := time.Parse(format, date)
	if err != nil {
		return 0, errors.New("Date should be in YYYY-MM-DD format")
	}

	year := time.Year()

	return year, nil
}

// Validator: Get the year from the input start date and end date.
// Start date should be before end date.
// Start date and end date must be in the same year to reduce size of data being queried.
func getYearValidateStartEndDate(startDate string, endDate string) (int, error) {
	start, err := time.Parse(format, startDate)
	if err != nil {
		return 0, errors.New("Start date should be in YYYY-MM-DD format")
	}

	end, err := time.Parse(format, endDate)
	if err != nil {
		return 0, errors.New("End date should be in YYYY-MM-DD format")
	}

	checkStartEnd := start.Before(end)

	if checkStartEnd == false {
		return 0, errors.New("End date is before Start Date")
	}

	startYear := start.Year()
	endYear := end.Year()

	if startYear != endYear {
		return 0, errors.New("Start date and End date must be in the same year")
	}

	return startYear, nil
}
