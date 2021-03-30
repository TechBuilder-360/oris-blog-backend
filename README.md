# oris-blog-backend
Oris Blog


npm i ckeditor5

## Environment Setup

Download and install [Go](https://golang.org/doc/install)

### Install MongoDB

Ensure that system has Docker running. Otherwise visit [Docker Website](https://www.docker.com/products/docker-desktop) to install

Run `docker pull mongo` in terminal

Run `docker run --name mongo-db -p 27017:27017 -d mongo:latest`



### Environment variables

DATABASE_NAME=<specify_database_name>

BLOG_COLLECTION=<specify_collection_name>

PORT=:<specify_port> ':' is compulsory e.g :8000

### Run project

Run `go get -u ./...` to install dependencies

Run `go build` to build app or Run `go run main.go` to run app

Run `./blog.[extension]` to run the built app. 


### Run Test Suite

Run `go test -v ./...`
