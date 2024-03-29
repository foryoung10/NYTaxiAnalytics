// Package for connecting and querying databases
package database

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

// Creates a Big Query Client connection
type BigQueryClient struct{}

// Configuration for Big Query setup
type Configuration struct {
	ApplicationCredentialsPath string
	ProjectName                string
}

var config = Configuration{}

// Set to true to use dry run.
const dryRun = false

// Initializing Big Query Client, reading config file for Google application credentials and project name and setting configuration
func BigQueryClientSetup() {
	log.Println("Initializing Big Query handler")

	file, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
	}

	if config.ApplicationCredentialsPath == "" {
		log.Println("GOOGLE_APPLICATION_CREDENTIALS environment must be set")
		os.Exit(1)
	}
}

// Using the Big Query Client, queries Big Query dataset using the Big Query api and return raw data.
// Pass in the query and query parameters
func (c BigQueryClient) Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error) {
	log.Println("Running BigQueryClient")

	// If config is not set
	if config.ProjectName == "" || config.ApplicationCredentialsPath == "" {
		BigQueryClientSetup()
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, config.ProjectName, option.WithCredentialsFile(config.ApplicationCredentialsPath))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := client.Query(q)
	// Set parameters
	query.Parameters = parameters
	// Set location
	query.Location = "US"
	// Set dry run
	query.DryRun = dryRun

	// Run query using job
	job, err := query.Run(ctx)
	if err != nil {
		return nil, err
	}

	// Wait for job to complete
	status, err := job.Wait(ctx)
	if err != nil {
		return nil, err
	}
	if err := status.Err(); err != nil {
		return nil, err
	}

	log.Println(job.Config())
	log.Println(job.LastStatus())

	//Read data from bigquery
	res, err := job.Read(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
