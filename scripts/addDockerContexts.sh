cd ~/docker-configs

set -o allexport
source ./.env

docker context create boree --docker "host=ssh://${SSH_USER}@${BOREE_IP}"
docker context create athena --docker "host=ssh://${SSH_USER}@${ATHENA_IP}"
docker context create euros --docker "host=ssh://${SSH_USER}@${EUROS_IP}"

set +o allexport