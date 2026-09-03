include .env
export

run-jarvis-fast:
	go run main.go

build-jarvis:
	go build -o jarvis ./cmd/main.go

run-jarvis:
	./jarvis

docker-build:
	docker compose build

docker-up:
	docker compose up

docker-down:
	docker compose down
