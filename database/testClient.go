// Package for connecting and querying databases
package database

import "cloud.google.com/go/bigquery"

// For testing of client
type TestClient struct{}

// For implementing Client interface
func (c TestClient) Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error) {

	return nil, nil
}
