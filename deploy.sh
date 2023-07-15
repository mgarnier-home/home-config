#!/bin/bash

set -o allexport
source ./.env

# docker stack config -c home.yml
docker stack deploy -c home.yml home

# docker stack config -c samba.yml
docker stack deploy -c samba.yml samba

# docker stack config -c backup.yml
docker stack deploy -c backup.yml backup

set +o allexport