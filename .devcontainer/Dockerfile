FROM mcr.microsoft.com/devcontainers/typescript-node:1-20-bullseye

# Install go
COPY --from=golang:1.22.2-alpine /usr/local/go/ /usr/local/go/
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# common apt install
RUN <<EOF
apt-get update -y
apt-get install -y --no-install-recommends \
    software-properties-common
apt-get clean
rm -rf /var/lib/apt/lists/*
EOF

# Install podman
RUN <<EOF
apt-get update -y
apt-get install -y --no-install-recommends \
    podman
apt-get clean
rm -rf /var/lib/apt/lists/*

podman completion bash -f /etc/bash_completion.d/podman
EOF
