FROM mcr.microsoft.com/devcontainers/typescript-node:1-20-bullseye

# Install go
COPY --from=golang:1.22.2-alpine /usr/local/go/ /usr/local/go/
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
