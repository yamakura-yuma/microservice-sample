# Hello World Microservice

This is a simple Go microservice that returns "Hello, world!" on HTTP GET /

## How to run (locally)

```
go run main.go
```

## How to build and run with Podman

```
podman build -t hello-service .
podman run --rm -p 8080:8080 hello-service
```

Then access http://localhost:8080/ and you will see:

```
Hello, world!
```
