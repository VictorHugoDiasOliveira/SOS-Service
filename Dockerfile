FROM golang:1.22.4-alpine3.20

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /auth-service

EXPOSE 8080

CMD ["/auth-service"]