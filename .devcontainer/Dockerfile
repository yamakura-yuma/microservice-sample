FROM mcr.microsoft.com/devcontainers/base:bookworm

############################################################
# Common apt install
#   - software-properties-common: for add-apt-repository and other utilities
############################################################
RUN <<EOF
apt-get update -y
apt-get install -y --no-install-recommends \
    software-properties-common
apt-get clean
rm -rf /var/lib/apt/lists/*
EOF

############################################################
# Podman & rootless Podman support
#   - podman: main package
#   - uidmap: for rootless podman user namespace
#   - slirp4netns: for rootless podman network isolation
############################################################
RUN <<EOF
apt-get update -y
apt-get install -y --no-install-recommends \
    podman \
    podman-compose \
    uidmap \
    slirp4netns
apt-get clean
rm -rf /var/lib/apt/lists/*
EOF
