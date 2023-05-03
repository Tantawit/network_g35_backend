# Socket Chat Server

## Stacks
- golang
- websocket
- gRPC
- cassandra
- redis

## Getting Start
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- golang 1.19 or [later](https://go.dev)
- docker
- makefile

### Installing
1. Clone the project from this repository
2. Import project
3. Setup config
    1. Copy `app.example.yaml` in `app` and paste it in the same location then remove `.example` from its name.
    1. Copy `database.example.yaml` in `database` and paste it in the same location then remove `.example` from its name.
4. Download dependencies by `go mod download`

### Testing
1. Run `go test -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./...` or `make test`

### Running
1. Run `docker-compose up -d` or `make compose-up`
2. Run `go run ./src/.` or `make server`

### Compile proto file
1. Run `make proto`
