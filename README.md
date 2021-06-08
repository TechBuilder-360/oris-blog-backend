# oris-blog-backend
Oris Blog

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

## API Docs
1. To set up go-swagger, run `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
1. Run `swagger generate spec -o ./swagger.json --scan-models`
1. Run `swagger serve -F=swagger swagger.json`

To convert generated json to html
1. Run `npm install -g redoc-cli`
1. Run `redoc-cli bundle -o templates/index.html swagger.json`

### Run Test Suite

Run `go test -v ./...`
