# build bin
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /backend-app ./cmd/main.go

# run bin
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /backend-app /app/backend-app

EXPOSE 8080

CMD ["./backend-app"]