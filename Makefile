include .env
export

check-micro:
	ffmpeg -f avfoundation -list_devices true -i ""

run-jarvis-fast:
	@go run -ldflags="-extldflags=-Wl,-no_warn_duplicate_libraries" main.go

build-jarvis:
	@go build -ldflags="-extldflags=-Wl,-no_warn_duplicate_libraries" -o jarvis ./cmd/main.go

run-jarvis:
	./jarvis

docker-build:
	@docker compose build

docker-up:
	@docker compose up

docker-down:
	@docker compose down

env-up:
	@docker compose up -d jarvis-postgres

env-stop:
	@docker compose stop jarvis-postgres

env-down:
	@docker compose down jarvis-postgres

env-cleanup:
	@read -p "Are you sure you want to cleanup the environment? (y/n): " ans;\
	if [ "$$ans" = "y" ]; then\
		docker compose down jarvis-postgres && \
		sudo rm -rf out/pgdata && \
		echo "Environment cleaned up successfully";\
	else \
		echo "Environment cleanup cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
			echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq={seq}"; \
			exit 1; \
		fi; \
		docker compose run --rm jarvis-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
			echo "Отсутствует необходимый параметр action. Пример: make migrate-action action={action}"; \
			exit 1; \
		fi; \

		docker compose run --rm jarvis-postgres-migrate \
       	-path /migrations \
       	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@jarvis-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
        "$(action)"
