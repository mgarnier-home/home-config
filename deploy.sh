#!/bin/bash

set -o allexport
source ./.env

docker stack config -c docker-compose.yml
docker stack deploy -c docker-compose.yml home

set +o allexport