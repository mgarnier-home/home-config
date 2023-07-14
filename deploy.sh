#!/bin/bash

set -o allexport
source ./.env

# docker stack config -c home.yml
docker stack deploy -c home.yml home

# docker stack config -c samba.yml
docker stack deploy -c samba.yml samba

set +o allexport