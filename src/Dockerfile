FROM golang:latest AS builder

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./go.work", "./go.work.sum", "./"]
COPY ./internal ./internal

RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./bin/app cmd/app/main.go

FROM golang:latest AS runner

WORKDIR /usr/local/src

COPY --from=builder /usr/local/src/bin/app ./
COPY configs/config.yml ./configs/config.yml
COPY .env ./.env

CMD ["./app"]