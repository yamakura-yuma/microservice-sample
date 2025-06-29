# Simple Go Hello World microservice
FROM docker.io/library/golang:bookworm AS build
ENV GOTOOLCHAIN=auto

WORKDIR /app

ENV GO111MODULE="on"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./build/http/main ./cmd/http/main.go



FROM docker.io/library/alpine:3.19
WORKDIR /app
COPY --from=build /app/build/http /app/build/http
EXPOSE 8080
CMD ["/app/build/http/main"]
