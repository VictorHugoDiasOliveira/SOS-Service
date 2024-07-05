FROM golang:1.22.4 AS builder

WORKDIR /app

COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
GOOS=linux go build -o sos-service .


FROM golang:1.22.4 AS runner

RUN useradd -ms /bin/bash victor

COPY --from=builder /app/sos-service /app/sos-service

RUN chown -R victor:victor /app
RUN chmod +x /app/sos-service

EXPOSE 8080

USER victor

CMD ["/app/sos-service"]