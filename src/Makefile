run: build-ui run-app
	./ui

build-ui:
	go build -o ui cmd/ui/main.go

run-app: build-app
	docker compose up -d ss-app ss-postgres ss-redis

build-app:
	docker build -t smartshopper:local .

migrate-up:
	./migrate.sh up

migrate-down:
	./migrate.sh down

clean:
	rm ./ui