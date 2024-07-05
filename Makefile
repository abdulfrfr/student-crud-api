build:
	go build -o student-crud-api

run:
	go run main.go

test:
	go test ./...

migrate:
	go run main.go migrate
