# Build image
FROM golang:1.20.4-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN go build -o /app/lib/pgdb/migration/main /app/lib/pgdb/migration/main.go

# Run in minimized image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/lib/pgdb/migration/main ./lib/pgdb/migration/
COPY ./lib/pgdb/migration/migrations ./lib/pgdb/migration/migrations
COPY .env .
COPY start.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]