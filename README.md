# Marks_RestAPI
## Installation

Install Docker dan Docker Compose first, then run this command :

```bash
$ docker compose up -d --build
```

## Build Application

If you have update the source codes then want to make changes to the executable application, first enter to inside docker container with this command :

```bash
$ docker exec -it student-api sh
```

Then run this command : 
```bash
$ go build -o student-api ./cmd
```

## Curl commands
Migration (Drop existing student data and generate new student marks value)
```bash
curl --request GET --url localhost:8080/migrations
```
Fetch Student data with parameter StartDate, EndDate, min Marks, and max Marks
```bash
curl --request POST --url http://localhost:8080/records --header 'Content-Type:application/json' --data '{
    "startDate": "2016-01-26", "endDate": "2018-02-02", "minCount": 100, "maxCount": 300}'
```