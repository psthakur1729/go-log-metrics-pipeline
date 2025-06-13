APP_NAME = go-log-metrics-pipeline

run:
	go run ./cmd/main.go

prometheus-up:
	docker-compose -f deploy/docker-compose.yml up -d

prometheus-down:
	docker-compose -f deploy/docker-compose.yml down

logs:
	docker-compose -f deploy/docker-compose.yml logs -f

build:
	go build -o $(APP_NAME) ./cmd/main.go
