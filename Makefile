test:
	go vet ./...
	go test  -v -coverpkg ./src/internal/... -coverprofile coverage.out -covermode count ./src/internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

server:
	go run ./src/cmd/main.go

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

seed:
	go run ./src/. seed