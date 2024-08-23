FROM golang:latest AS builder

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./go.work", "./go.work.sum", "./"]

COPY ./internal/repositories ./internal/repositories
COPY ./internal/services ./internal/services
COPY ./internal/api ./internal/api

RUN go mod download

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