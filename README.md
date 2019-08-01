# NY Taxi Analytics api

## Setting up the application

#### config.json
Replace [GOOGLE_APPLICATION_CREDENTIALS PATH] and [PROJECT NAME]  
Go to https://cloud.google.com/bigquery/docs/reference/libraries#client-libraries-install-go for more info


```json
{
    "ApplicationCredentialsPath": "[GOOGLE_APPLICATION_CREDENTIALS PATH]",
    "ProjectName": "[PROJECT_NAME]"
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
