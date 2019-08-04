package database

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
)

type BigQueryClient struct{}

type Configuration struct {
	ApplicationCredentialsPath string
	ProjectName                string
}

var config = Configuration{}

const dryRun = false

// to do initialize in main method
// Initializing Big Query handler, reading config file for
// Google application credentials and project name

func BigQueryClientSetup() {
	log.Println("Initializing Big Query handler")

	file, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.ApplicationCredentialsPath)

	if config.ApplicationCredentialsPath == "" {
		log.Println("GOOGLE_APPLICATION_CREDENTIALS environment must be set")
		os.Exit(1)
	}

}

// Exceute a query using the big query client
func (c BigQueryClient) Query(q string, parameters []bigquery.QueryParameter) (*bigquery.RowIterator, error) {
	log.Println("Running BigQueryClient")

	if config.ProjectName == "" || config.ApplicationCredentialsPath == "" {
		BigQueryClientSetup()
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, config.ProjectName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := client.Query(q)
	query.Parameters = parameters
	query.Location = "US"
	query.DryRun = dryRun

	job, err := query.Run(ctx)
	if err != nil {
		return nil, err
	}

	status, err := job.Wait(ctx)
	if err != nil {
		return nil, err
	}
	if err := status.Err(); err != nil {
		return nil, err
	}

	log.Println(job.Config())
	log.Println(job.LastStatus())

	res, err := job.Read(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
