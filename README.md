# NY Taxi Analytics api

## Setting up the application

For Big Query api  
Add Google application credentials and project name to config.json file.
For more info go to <https://cloud.google.com/docs/authentication/getting-started>

### config.json

```json
{
    "ApplicationCredentialsPath": "",
    "ProjectName": ""
}
```

## Running the application

```cmd
go run main.go
```

## Building the application

```cmd
go build
```

## Api Endpoints  

### Total trips per day

```http
http://localhost:8080/total_trips?start=<start-date>&end=<end-date>
```

### Fare heatmap

```http
http://localhost:8080/average_fare_heatmap?date=<date>>
```

### Average speed in the past 24 hours

```http
http://localhost:8080/average_speed_24hrs?date=<date>>
```

## Design of Api

The main design of the api is to let Big query do the heavy lifting whenever possible.  
Big query can process large volumes of datasets very quickly.  
Thus for 2 of the endpoints **Average speed in the past 24 hours** and **Total trips per day is a direct load** loads from Big query processing.  
For the **Fare heatmap** data some processing is needed from the bigquery data.

There are 2 packages  
**taxi**: Package for the taxi entity contains handler, model, queries, repositories, service  
**database**: Package for connecting and querying databases  

### taxi package

This is how a request to the api is processed.  

**handler** -> **service** -> **repository** -> **client (database)**  

The application is structured this way so that the layers are interchangeable and can be replaced or reused.

#### handler

Handler handles and validates requests from the api.  
Validates the query string for date input is in the correct format, start date before end date  
Handle the return of the api response

#### service

Service is the business logic layer to transform dataset to the output.
For the fares heatmap the s2 library is used to transform lat, long to s2id level 16 using the region coverer algorithm  
See <https://godoc.org/github.com/golang/geo/s2#RegionCoverer.Covering>  
After the s2id is obtained calculates the average fare.

#### repository

Repository handles data transfer between application and database
The repository generates the tables to query, the query and the query parameters.
The tables to query is generated based on the year as the bigquery data set is large there is no point to query all of the tables.

#### client (database)

The client connects and queries the database.
Setup and connects to the Big query client with the query and query parameters.
