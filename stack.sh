#!/bin/bash

set -o allexport
source ./.env
source ./ports.env

additional_compose_files="-f ../volumes.yml"

manage_stack() {
    cd compose/$stack

    local stack=$1
    local action=$2
    local host=$3
    local file="$host.$stack.yml"

    if [ ! -f "$file" ]; then
        echo "File $file does not exist"
    else
        local compose_cmd="docker compose -f $file $additional_compose_files"
        
        export DOCKER_CONTEXT=$host

        case $action in 
            deploy)
                echo "Deploying $stack stack on $host..."
                $compose_cmd up -d
                ;;
            undeploy)
                echo "Undeploying $stack stack on $host..."
                $compose_cmd down
                ;;
            redeploy)
                echo "Redeploying $stack stack on $host..."
                $compose_cmd down
                $compose_cmd up -d
                ;;
            pull)
                echo "Pulling $stack stack on $host..."
                $compose_cmd down
                $compose_cmd pull
                $compose_cmd up -d
                ;;
        esac
    fi

    cd ../..
}

manage_stacks() {
    local stack=$1
    local action=$2
    local host=$3


    if [ "$host" != "all" ]; then
        manage_stack $stack $action $host
        return
    fi

    for file in $(ls *.$stack.yml); do
        local host_name=$(echo $file | cut -d'.' -f1)

        manage_stack $stack $action $host_name
    done
}



ACTION=${1}
HOSTNAME=${3:-all}

case $ACTION in
    deploy|undeploy|redeploy|pull)
        ;;
    *)
        echo "Unknown action: $ACTION"
        echo "Usage: $0 [deploy|undeploy|redeploy|pull]"
        exit 1
        ;;
esac

case $HOSTNAME in
    all|athena|apollon|artemis|hermes)
        ;;
    *)
        echo "Unknown host: $HOSTNAME"
        echo "Usage: $0 [athena|apollon|hermes|artemis|all]"
        exit 1
        ;;
esac

STACK=${2:-all}

echo "Args are: $ACTION $STACK $HOSTNAME"

case $STACK in
    samba)
        manage_stacks samba $ACTION $HOSTNAME
        ;;
    network)
        manage_stacks network $ACTION $HOSTNAME
        ;;
    monitoring)
        manage_stacks monitoring $ACTION $HOSTNAME
        ;;
    file_server)
        manage_stacks file_server $ACTION $HOSTNAME
        ;;
    plex)
        manage_stacks plex $ACTION $HOSTNAME
        ;;
    nextcloud)
        manage_stacks nextcloud $ACTION $HOSTNAME
        ;;
    paperless)
        manage_stacks paperless $ACTION $HOSTNAME
        ;;
    minecraft)
        manage_stacks minecraft $ACTION $HOSTNAME
        ;;
    backup)
        manage_stacks backup $ACTION $HOSTNAME
        ;;
    db)
        manage_stacks db $ACTION $HOSTNAME
        ;;
    all)
        manage_stacks samba $ACTION $HOSTNAME
        manage_stacks network $ACTION $HOSTNAME
        manage_stacks monitoring $ACTION $HOSTNAME
        manage_stacks file_server $ACTION $HOSTNAME
        manage_stacks plex $ACTION $HOSTNAME
        manage_stacks nextcloud $ACTION $HOSTNAME
        manage_stacks paperless $ACTION $HOSTNAME
        manage_stacks minecraft $ACTION $HOSTNAME
        manage_stacks backup $ACTION $HOSTNAME
        manage_stacks db $ACTION $HOSTNAME
        ;;
    *)
        echo "Unknown stack: $STACK"
        echo "Usage: $0 [samba|network|monitoring|file_server|plex|nextcloud|paperless|minecraft|backup|db|all]"
        exit 1
        ;;
esac

set +o allexport