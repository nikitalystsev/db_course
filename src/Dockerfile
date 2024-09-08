FROM golang:latest AS builder

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./go.work", "./go.work.sum", "./"]

COPY ./internal/repositories/go.mod ./internal/repositories/go.mod
COPY ./internal/repositories/go.sum ./internal/repositories/go.sum
COPY ./internal/services/go.mod ./internal/services/go.mod
COPY ./internal/services/go.sum ./internal/services/go.sum
COPY ./internal/api/go.mod ./internal/api/go.mod
COPY ./internal/api/go.sum ./internal/api/go.sum

RUN go mod download

COPY ./internal/repositories ./internal/repositories
COPY ./internal/services ./internal/services
COPY ./internal/api ./internal/api

COPY ./cmd/app ./cmd/app
COPY ./internal/app ./internal/app
COPY ./internal/config ./internal/config

RUN go build -o ./bin/app cmd/app/main.go

FROM golang:latest AS runner

WORKDIR /usr/local/src

COPY --from=builder /usr/local/src/bin/app ./
COPY configs/config.yml ./configs/config.yml
COPY .env ./.env

CMD ["./app"]