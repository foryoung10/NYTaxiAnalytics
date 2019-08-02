package database

import "cloud.google.com/go/bigquery"

type Client interface {
	Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error)
}
