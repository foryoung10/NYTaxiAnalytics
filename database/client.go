// Package for connecting and querying databases
package database

import "cloud.google.com/go/bigquery"

// Client is a connection to the database.
// Query uses the client to query the database.
type BqClient interface {
	Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error)
}

type DbClient interface {
}
