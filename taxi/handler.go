package taxi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	FetchTotalTrips() func(c *gin.Context)
	FetchAverageSpeed() func(c *gin.Context)
	FetchAverageFareS2id() func(c *gin.Context)
}

type Handler struct {
	Svc IService
}

func (hand Handler) FetchTotalTrips() func(c *gin.Context) {
	return func(c *gin.Context) {
		startDate := c.Query("start-date")
		endDate := c.Query("end-date")

		res, err := hand.Svc.GetTotalTripsByStartEndDate(startDate, endDate)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		jsonData, err := json.Marshal(res)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, jsonData)
	}
}

func (hand Handler) FetchAverageSpeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Query("date")

		res, err := hand.Svc.GetAverageSpeedByDate(date)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		jsonData, err := json.Marshal(res)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, jsonData)
	}
}

func (hand Handler) FetchAverageFareS2id() func(c *gin.Context) {
	return func(c *gin.Context) {
		date := c.Query("date")

		const level = 16
		res, err := hand.Svc.GetAverageFarePickUpByLocation(date, 16)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			c.JSON(http.StatusNoContent, "")
			return
		}

		jsonData, err := json.Marshal(res)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, jsonData)
	}
}

// Returns the previous date and year
func getPreviousDateYear(date string) (string, int, error) {
	const format = "2006-01-02"
	time, err := time.Parse(format, date)

	if err != nil {
		return "", 0, errors.New("Date should be in YYYY-MM-DD format")
	}

	previousDate := time.AddDate(0, 0, -1)

	year := previousDate.Year()

	return previousDate.Format(format), year, nil
}

// Returns the year
func getYearValidateDate(date string) (int, error) {
	const format = "2006-01-02"
	time, err := time.Parse(format, date)
	if err != nil {
		return 0, errors.New("Date should be in YYYY-MM-DD format")
	}
	year := time.Year()

	return year, nil
}

// Returns the year
func getYearValidateStartEndDate(startDate string, endDate string) (int, error) {
	const format = "2006-01-02"

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