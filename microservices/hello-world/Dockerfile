# Simple Go Hello World microservice
FROM docker.io/library/golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/main"]
