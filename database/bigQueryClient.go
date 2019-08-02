package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

type BigQueryClient struct{}

type Configuration struct {
	ApplicationCredentialsPath string
	ProjectName                string
}

var config = Configuration{}

// to do initialize in main method
// Initializing Big Query handler, reading config file for
// Google application credentials and project name
/*
func init() {
	fmt.Println("Initializing Big Query handler")

	file, err := os.Open("../config.json")
	if err != nil {
		fmt.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.ApplicationCredentialsPath)
	proj := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if proj == "" {
		fmt.Println("GOOGLE_APPLICATION_CREDENTIALS environment must be set")
		os.Exit(1)
	}

	fmt.Println(config.ApplicationCredentialsPath, config.ProjectName)
}
*/
// Exceute a query using the big query client
func (c BigQueryClient) Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error) {
	fmt.Println("running BigQueryClient")
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, config.ProjectName)
	if err != nil {
		fmt.Println(err)
		//	return nil, err
	}

	query := client.Query(q)
	query.Parameters = parameters
	query.DryRun = true

	fmt.Println(query.Q)

	return query.Read(ctx)
}
