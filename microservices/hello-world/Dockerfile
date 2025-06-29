# Simple Go Hello World microservice
FROM docker.io/library/golang:bookworm AS build
ENV GOTOOLCHAIN=auto

WORKDIR /app

ENV GO111MODULE="on"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# 完全な静的バイナリでビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/http/main ./cmd/http/main.go



FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build /app/build/http /app/build/http
EXPOSE 8081
CMD ["/app/build/http/main"]
