build:
	 docker-compose build
	 docker-compose up

run:
	go run main.go

mod:
	go mod tidy
	go mod vendor

