set -o allexport
source ./.env

docker context create apollon --docker "host=ssh://${SSH_USER}@${APOLLON_IP}"
docker context create athena --docker "host=ssh://${SSH_USER}@${ATHENA_IP}"
docker context create artemis --docker "host=ssh://${SSH_USER}@${ARTEMIS_IP}"
docker context create hermes --docker "host=ssh://${SSH_USER}@${HERMES_IP}"

set +o allexport