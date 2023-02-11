################################################################################

ARG GO_VERSION="1.20"
ARG BUILD_TARGET="consumer"
ARG BUILD_TYPE="debug"

################################################################################

FROM docker.io/library/golang:${GO_VERSION} AS base-consumer
WORKDIR /usr/src

COPY examples/consumer/go.mod examples/consumer/go.sum ./
RUN go mod download && go mod verify

COPY examples/consumer/ .
RUN go build -v -o /usr/local/bin/app

################################################################################

FROM docker.io/library/golang:${GO_VERSION} AS base-producer
WORKDIR /usr/src

COPY examples/producer/go.mod examples/producer/go.sum ./
RUN go mod download && go mod verify

COPY examples/producer/ .
RUN go build -v -o /usr/local/bin/app

################################################################################

FROM base-${BUILD_TARGET} AS base

FROM docker.io/library/debian:11 AS app-release
COPY --from=base /usr/local/bin/app /usr/local/bin/app

FROM base AS app-debug
RUN set -eux; \
    DEBIAN_FRONTEND=noninteractive apt-get update; \
    DEBIAN_FRONTEND=noninteractive apt-get install -y -q tmux vim telnet; \
    rm -rf /var/cache/apt;

################################################################################

FROM app-${BUILD_TYPE} AS app
WORKDIR /
CMD ["app"]