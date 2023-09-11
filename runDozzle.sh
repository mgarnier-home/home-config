#!/bin/bash

set -o allexport
source ./.env

docker rm -f dozzle
docker pull amir20/dozzle:latest
docker run \
  --name dozzle \
  -p 10002:8080 \
  -d \
  amir20/dozzle:latest \
  --remote-host "tcp://${HERMES_IP}:10001|Hermes" \
  --remote-host "tcp://${APOLLON_IP}:10001|Apollon" \
  --remote-host "tcp://${ATHENA_IP}:10001|Athena" \
  --remote-host "tcp://${ARTEMIS_IP}:4201|Artemis"
  

set +o allexport