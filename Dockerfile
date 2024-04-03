# Use the official Golang image as a base
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o serve/go-avatar .

FROM alpine:latest  

RUN mkdir /app

COPY --from=builder /app/serve/go-avatar /app/serve/go-avatar

EXPOSE 8055

CMD ["/app/serve/go-avatar"]
