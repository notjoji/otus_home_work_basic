FROM golang:1.22.0-alpine3.18 as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /otus_db

COPY docker.env /.env

FROM alpine:latest

COPY --from=builder /otus_db /otus_db

COPY --from=builder /.env .env

EXPOSE 8080

CMD ["/otus_db"]