#!/bin/bash

set -o allexport
source ./.env

manage_stack() {
    local stack=$1
    local action=$2

    cd compose/$stack

    local additional_compose_files="-f ../volumes.yml"

    for file in $(ls *.$stack.yml); do
        local host_name=$(echo $file | cut -d'.' -f1)

        local compose_cmd="docker compose -f $file $additional_compose_files"
        
        export DOCKER_CONTEXT=$host_name



        case $action in 
            deploy)
                echo "Deploying $stack stack on $host_name..."
                $compose_cmd pull
                $compose_cmd up -d
                ;;
            undeploy)
                echo "Undeploying $stack stack on $host_name..."
                $compose_cmd down
                ;;
        esac
    done
    
    cd ../..

}



ACTION=${1}

case $ACTION in
    deploy|undeploy)
        ;;
    *)
        echo "Unknown action: $ACTION"
        echo "Usage: $0 [deploy|undeploy|ps]"
        exit 1
        ;;
esac


STACK=${2:-all}

case $STACK in
    samba)
        manage_stack samba $ACTION
        ;;
    network)
        manage_stack network $ACTION
        ;;
    monitoring)
        manage_stack monitoring $ACTION
        ;;
    home)
        manage_stack home $ACTION
        ;;
    minecraft)
        manage_stack minecraft $ACTION
        ;;
    backup)
        manage_stack backup $ACTION
        ;;
    all)
        manage_stack samba $ACTION
        manage_stack network $ACTION
        manage_stack monitoring $ACTION
        manage_stack home $ACTION
        manage_stack minecraft $ACTION
        manage_stack backup $ACTION
        ;;
    *)
        echo "Unknown stack: $STACK"
        echo "Usage: $0 [samba|network|monitoring|minecraft|all]"
        exit 1
        ;;
esac

set +o allexport