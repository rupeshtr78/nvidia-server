# https://hub.docker.com/_/golang/tags
FROM golang:1.22.2-bullseye AS builder  
ARG TARGETOS
ARG TARGETARCH

# Install required packages
RUN set -ex; \
    apt-get update; \
    apt-get install -y --no-install-recommends curl ca-certificates gcc; \
    update-ca-certificates; \
    rm -rf /var/lib/apt/lists/*;

RUN mkdir /app
WORKDIR /app

# RUN go env -w GOPRIVATE=*.

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Add content
ADD . .

# Run tests
# RUN go test -v ./...

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=1 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -o nvidia-server ./cmd/main.go

# Image to run the service in production
FROM nvidia/cuda:12.4.1-cudnn-runtime-ubuntu22.04

# add git commit label
ARG GIT_COMMIT=unspecified
LABEL git_commit=$GIT_COMMIT

# add required trusted root CA

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections
RUN echo "keyboard-configuration keyboard-configuration/layout select US" | debconf-set-selections 
RUN echo "keyboard-configuration keyboard-configuration/variant select English (US)" | debconf-set-selections 
RUN echo "keyboard-configuration keyboard-configuration/modelcode select pc105" | debconf-set-selections

# Install required packages
RUN set -ex; \
    apt-get update; \
    apt-get -y install curl daemontools ca-certificates software-properties-common; \
    add-apt-repository ppa:graphics-drivers/ppa; \
    apt-get update; \
    apt-get -y install nvidia-driver-390;\
    update-ca-certificates; \
    apt-get install -y --no-install-recommends \
    build-essential \
    freeglut3-dev \
    libx11-dev \
    libxmu-dev \
    libxi-dev \
    libnuma-dev \
    libbz2-1.0 \
    libbz2-dev \
    libbz2-ocaml \
    libbz2-ocaml-dev \
    libssl-dev \
    openssl \
    librdmacm-dev \
    libibverbs-dev \
    libgtk2.0-dev; \
    apt-get clean; \
    apt-get autoremove --purge; \
    rm -rf /var/lib/apt/lists/* /usr/share/man/* /usr/share/doc/*;



# Install tini https://github.com/krallin/tini as signal handler
# tini handles the reaping of zombie processes and to forward signals correctly to subprocesses.
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

# Create a non-root user and group with a writeable home directory
# These UID/GID values should match who we run the container as
WORKDIR /app
RUN mkdir /config \
    && mkdir /logs
COPY config/metrics.yaml /config/metrics.yaml

# add non-root user
ARG USER_NAME=user
ARG USER_UID=999
ARG USER_GID=$USER_UID

# setup permissions to prevent root access
RUN groupadd --gid $USER_GID $USER_NAME \
    && useradd --uid $USER_UID --gid $USER_GID --shell /bin/bash -m $USER_NAME \
    && mkdir -p /etc/sudoers.d \
    && echo "$USER_NAME ALL=(ALL:ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER_NAME \
    && chmod 0440 /etc/sudoers.d/$USER_NAME

RUN chown -R $USER_NAME: /app && \
    chown -R $USER_NAME: /config && \
    chown -R $USER_NAME: /logs

# set user
USER $USER_NAME

COPY --chown=$USER_UID:$USER_GID --from=builder /app/nvidia-metrics /app/nvidia-metrics


# Set environment variables with default values
ENV LOG_LEVEL=info
ENV PORT=9500
ENV HOST=0.0.0.0
ENV LOG_FILE_PATH=/logs/nvidia-metrics.log
ENV LOG_TO_FILE=false

EXPOSE $PORT
# use tini to perform correct signal handling
ENTRYPOINT ["/tini", "--"]

# server should be the default command
CMD ["/app/nvidia-server"]





