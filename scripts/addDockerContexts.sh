SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR/..

set -o allexport
source .env

docker context create boree --docker "host=ssh://${SSH_USER}@${BOREE_IP}"
docker context create athena --docker "host=ssh://${SSH_USER}@${ATHENA_IP}"
docker context create euros --docker "host=ssh://${SSH_USER}@${EUROS_IP}"
docker context create notos --docker "host=ssh://${SSH_USER}@${NOTOS_IP}"
docker context create zephyr --docker "host=ssh://${SSH_USER}@${ZEPHYR_IP}"

set +o allexport