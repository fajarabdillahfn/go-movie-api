# movie-service

# pull the official docker image
FROM golang:1.17

WORKDIR /movie-app

# install dependencies
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o .

CMD [ "./go-movie-api" ]