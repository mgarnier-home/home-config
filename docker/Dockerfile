FROM golang:1.22.3-alpine3.18 as builder

RUN apk add curl

WORKDIR /build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

COPY . .

RUN ls -la

RUN task build

FROM jamesread/olivetin as runtime

ENV ANSIBLE_VERSION 2.16.7

USER root

RUN microdnf install -y --nodocs --noplugins --setopt=keepcache=0 --setopt=install_weak_deps=0 ansible

COPY --from=builder /build/bin/home-cli /usr/local/bin/home-cli

RUN home-cli build-olivetin
