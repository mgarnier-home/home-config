#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR/..

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

    for filePath in $(ls compose/$stack/*.$stack.yml); do
        local file=$(basename $filePath)
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
    all|athena|euros|boree|notos|zephyr)
        ;;
    *)
        echo "Unknown host: $HOSTNAME"
        echo "Usage: $0 [athena|euros|boree|notos|zephyr|all]"
        exit 1
        ;;
esac

STACK=${2:-all}

echo "Args are: $ACTION $STACK $HOSTNAME"


STACKS=("samba" "backup" "db" "file_server" "home_assistant" "media" "minecraft" "misc" "monitoring" "network" "nextcloud" "paperless" "palworld" "runner" "syslog")

case $STACK in
    all)
        # If the selected stack is 'all', loop through all stacks
        for stack in "${STACKS[@]}"; do
            manage_stacks "$stack" "$ACTION" "$HOSTNAME"
        done
        ;;
    *)
        # Check if the selected stack is in the STACKS array
        if [[ " ${STACKS[*]} " =~ " $STACK " ]]; then
            # Call manage_stacks for the selected stack
            manage_stacks "$STACK" "$ACTION" "$HOSTNAME"
        else
            # If the stack is not found, show an error message
            echo "Unknown stack: $STACK"
            echo "Available stacks: ${STACKS[*]}"
            exit 1
        fi
        ;;
esac

set +o allexport