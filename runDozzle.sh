#!/bin/bash

set -o allexport
source ./.env
source ./ports.env

docker rm -f dozzle
docker pull amir20/dozzle:latest
docker run \
  --name dozzle \
  -p "${APOLLON_DOZZLE_PORT}:8080" \
  -d \
  amir20/dozzle:latest \
  --remote-host "tcp://${HERMES_IP}:${HERMES_SOCKET_PROXY_PORT}|Hermes" \
  --remote-host "tcp://${APOLLON_IP}:${APOLLON_SOCKET_PROXY_PORT}|Apollon" \
  --remote-host "tcp://${ATHENA_IP}:${ATHENA_SOCKET_PROXY_PORT}|Athena" \
  --remote-host "tcp://${ARTEMIS_IP}:${ARTEMIS_SOCKET_PROXY_PORT}|Artemis" \
  

set +o allexport