// package data contains sample data obtained from google big query
// dataset: bigquery-public-data.new_york.tlc_yellow_trips_2015
// development to be done using sample data before calling the big query api
package taxi

// all pickup locations lat, long and fare amount
var faresData = `
[   
    {
      "pickup_longitude": "-51.1408982",
      "pickup_latitude" : "-30.0441054",
      "fare_amount": "4.32"
    },
    {
        "pickup_longitude": "-73.91737365722656",
        "pickup_latitude": "40.61402130126953",
        "fare_amount": "27.5"
    },
    {
        "pickup_longitude": "-73.9610824584961",
        "pickup_latitude": "40.7188835144043",
        "fare_amount": "24.5"
    },
    {
        "pickup_longitude": "-73.97857666015625",
        "pickup_latitude": "40.670406341552734",
        "fare_amount": "26.5"
    },
    {
        "pickup_longitude": "-73.85911560058594",
        "pickup_latitude": "40.73643112182617",
        "fare_amount": "40.0"
    },
    {
        "pickup_longitude": "-73.994873046875",
        "pickup_latitude": "40.703060150146484",
        "fare_amount": "27.0"
    },
    {
        "pickup_longitude": "-73.92514038085938",
        "pickup_latitude": "40.807518005371094",
        "fare_amount": "0.0"
    },
    {
      "pickup_longitude": "-51.1395916",
      "pickup_latitude": "-30.0436491",
      "fare_amount": "10.0"
    },
    {
      "pickup_longitude": "-51.1905683",
      "pickup_latitude": "-30.0320944",
      "fare_amount": "10.0"
    },
    {
      "pickup_longitude": "-51.1905683",
      "pickup_latitude": "-30.0331944",
      "fare_amount": "10.0"
    },
    {
      "pickup_longitude": "-74.005106",
      "pickup_latitude": "40.710977",
      "fare_amount": "10.0"
    },
    {
      "pickup_longitude": "-74.005433",
      "pickup_latitude": "40.711306",
      "fare_amount": "15.0"
    },
    {
      "pickup_longitude": "-74.009666",
      "pickup_latitude": "40.714963",
      "fare_amount": "10.0"
    },
    {
      "pickup_longitude": "-74.010205",
      "pickup_latitude": "40.715170",
      "fare_amount": "15.0"
    },
    {
      "pickup_longitude": "-74.009881",
      "pickup_latitude": "40.714709",
      "fare_amount": "30.0"
    }
]`

// the average speed of all trips over the last 24 hours
var averageSpeedData = `
        [
            {
                "average_speed": 14.1
            }
        ]
`

// total number of trips on each day
var totalTripsData = `
[
  {
    "date": "2015-01-01",
    "total_trips": "382014"
  },
  {
    "date": "2015-01-02",
    "total_trips": "345296"
  },
  {
    "date": "2015-01-03",
    "total_trips": "406769"
  },
  {
    "date": "2015-01-04",
    "total_trips": "328848"
  },
  {
    "date": "2015-01-05",
    "total_trips": "363454"
  },
  {
    "date": "2015-01-06",
    "total_trips": "384324"
  },
  {
    "date": "2015-01-07",
    "total_trips": "429653"
  },
  {
    "date": "2015-01-08",
    "total_trips": "450920"
  },
  {
    "date": "2015-01-09",
    "total_trips": "447947"
  },
  {
    "date": "2015-01-10",
    "total_trips": "515540"
  },
  {
    "date": "2015-01-11",
    "total_trips": "419629"
  },
  {
    "date": "2015-01-12",
    "total_trips": "396367"
  }
]`
