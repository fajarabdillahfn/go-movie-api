# go-movie-api
This is a REST API to manage movies data.

## Pre-requisite
### Create .env that contains
* POSTGRES_USER
* POSTGRES_PASSWORD
* POSTGRES_DB
* SQL_HOST
* SQL_PORT
* APP_PORT

## Run the app
    docker-compose up -d

# REST API
The REST API to the example app is described below.

## Insert a New Movie
`POST /movies`

### Request Data (application/json)
example:
```
{
    "title": "Pengabdi Setan 2",
    "description": "adalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017",
    "rating": 9.0,
    "image": ""
}
```

## Get a List of Movies
`GET /movies`

## Get a Movie by Id
`GET /movies/:id`

## Update a Movie by Id
`PATCH /movies/:id`

### Request Data (application/json)
example:
```
{
    "title": "Pengabdi Setan 2",
    "description": "adalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017",
    "rating": 9.0,
    "image": ""
}
```

## Delete a Movie by Id
`DELETE /movies/:id`
