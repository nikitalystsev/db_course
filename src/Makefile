run: build-ui run-app
	./ui

build-ui:
	go build -o ui cmd/ui/main.go

run-app: build-app
	docker compose up -d ss-app ss-postgres ss-redis

build-app:
	docker build -t smartshopper:local .

sql-tests:
	cd internal/repositories/testing/ && docker build -f ./Dockerfile . -t pgtap-tester:latest
	docker-compose up -d ss-postgres
	cd internal/repositories/testing/ && ./test_runner.sh

migrate-up:
	./migrate.sh up

migrate-down:
	./migrate.sh down

clean:
	rm ./ui