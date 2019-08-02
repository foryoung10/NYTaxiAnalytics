package database

import "cloud.google.com/go/bigquery"

type TestClient struct{}

func (c TestClient) Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error) {

	return nil, nil
}
