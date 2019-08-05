// Package taxi contains library for the taxi entity.
// Handler, model, query, repository, service, data
package taxi

import (
	"testing"
)

// Test for all date validators.

func TestValidateDate(t *testing.T) {
	dates := []struct {
		input    string
		expected int
	}{
		{"2015-01-01", 2015},
		{"2019-01-10", 2019},
		{"2016-06-31", 0},
		{"10-10-2009", 0},
	}

	for _, date := range dates {
		res, _ := getYearValidateDate(date.input)
		if res != date.expected {
			t.Errorf("Incorrect result (%v) -Input (%v) Expected (%v)", res, date.input, date.expected)
		}
	}
}

func TestValidateStartEndDate(t *testing.T) {
	dates := []struct {
		startDateInput string
		endDateInput   string
		expected       int
	}{
		{"2015-01-01", "2015-01-31", 2015},
		{"2015-31-10", "2015-01-31", 0},
		{"2015-01-10", "2015-31-31", 0},
		{"2015-31-10", "2015-31-31", 0},
		{"2015-01-31", "2015-01-01", 0},
		{"2013-12-01", "2014-01-01", 0},
	}

	for _, date := range dates {
		res, _ := getYearValidateStartEndDate(date.startDateInput, date.endDateInput)
		if res != date.expected {
			t.Errorf("Incorrect result (%v) -Input (%v) (%v) Expected (%v)", res, date.startDateInput, date.endDateInput, date.expected)
		}
	}
}

func TestGetPreviousDateYear(t *testing.T) {
	dates := []struct {
		input        string
		expectedDate string
		expectedYear int
	}{
		{"2015-01-02", "2015-01-01", 2015},
		{"2016-01-01", "2015-12-31", 2015},
		{"2016-06-31", "", 0},
		{"10-10-2009", "", 0},
	}

	for _, date := range dates {
		prevDate, year, _ := getPreviousDateAndYearValidateDate(date.input)
		if prevDate != date.expectedDate {
			t.Errorf("Incorrect previous date (%v) -Input (%v) Expected (%v)", prevDate, date.input, date.expectedDate)
		}
		if year != date.expectedYear {
			t.Errorf("Incorrect year (%v) -Input (%v) Expected (%v)", year, date.input, date.expectedYear)
		}
	}
}
