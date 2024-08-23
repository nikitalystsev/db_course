run: build
	./app

build:
	docker compose up ss-postgres ss-redis
	go build -o app cmd/app/main.go

migrate-up:
	./migrate.sh up

migrate-down:
	./migrate.sh down