FROM golang:1.22.3-alpine3.18 as builder

RUN apk add curl

WORKDIR /build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

COPY . .


RUN ls -la

RUN task build

FROM jamesread/olivetin as runtime

COPY --from=builder /build/bin/home-cli /usr/local/bin/home-cli

RUN home-cli build-olivetin
