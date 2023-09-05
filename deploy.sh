#!/bin/bash

set -o allexport
source ./.env

deploy_stack() {
    local stack=$1
    local file=$2
    echo "Deploying $stack stack..."
    docker stack deploy -c $file $stack
}

STACK=${1:-all}

case $STACK in
    home)
        deploy_stack home home.yml
        ;;
    samba)
        deploy_stack samba samba.yml
        ;;
    backup)
        deploy_stack backup backup.yml
        ;;
    network)
        deploy_stack network network.yml
        ;;
    monitoring)
        deploy_stack monitoring monitoring.yml
        ;;
    minecraft)
        deploy_stack minecraft minecraft.yml
        ;;
    all)
        deploy_stack home home.yml
        deploy_stack samba samba.yml
        deploy_stack backup backup.yml
        deploy_stack network network.yml
        deploy_stack monitoring monitoring.yml
        deploy_stack minecraft minecraft.yml
        ;;
    *)
        echo "Unknown stack: $STACK"
        echo "Usage: $0 [home|samba|backup|network|monitoring|minecraft|all]"
        exit 1
        ;;
esac

set +o allexport