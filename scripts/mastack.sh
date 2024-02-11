#!/bin/bash

# DRY=true
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR/..

# Read JSON file
STACK_JSON_FILE="$SCRIPT_DIR/stack.json"
if [ ! -f "$STACK_JSON_FILE" ]; then
    echo "Error: $STACK_JSON_FILE not found."
    exit 1
fi

ACTIONS=($(jq -r '.actions[]' $STACK_JSON_FILE | tr -d '\r'))
HOSTS=($(jq -r '.hosts[]' $STACK_JSON_FILE | tr -d '\r'))
STACKS=($(jq -r '.stacks[]' $STACK_JSON_FILE | tr -d '\r'))

declare -A AFTER_SCRIPTS
while IFS="=" read -r key value; do
    AFTER_SCRIPTS["$key"]="$value"
done < <(jq -r '.afterScript | to_entries[] | "\(.key)=\(.value)"' $STACK_JSON_FILE | tr -d '\r')


# Retrieve environment variables from .env and ports.env
set -o allexport
source ./env/.env
source ./env/ports.env

# Set the additional compose files
additional_compose_files="-f ../volumes.yml"

print_array() {
    local array=("$1")
    for item in "${array[@]}"; do
        echo "$item"
    done
}

check_param() {
    local paramName=$1
    local param=$2
    local array=("$3")

    if [ $(echo ${array[*]} | grep -o $param | wc -l) -eq 0 ]; then
        echo "Unknown $paramName: $param"
        echo "Available $paramName: "
        print_array "${array[*]}"
        exit 1
    fi
}

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
        local script_key="$action $stack $host"

        if [ "$DRY" = true ]; then
            compose_cmd="echo DRY RUN $compose_cmd"
        fi
        
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

        if [ -n "${AFTER_SCRIPTS[$script_key]}" ]; then
            local after_script=${AFTER_SCRIPTS[$script_key]}
            echo "Running after script for $script_key : $after_script"
            $after_script
        fi

        unset DOCKER_CONTEXT
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
STACK=${2:-all}
HOSTNAME=${3:-all}

check_param "action" $ACTION "${ACTIONS[*]}"
check_param "stack" $STACK "${STACKS[*]}"
check_param "host" $HOSTNAME "${HOSTS[*]}"

echo "Args are: $ACTION $STACK $HOSTNAME"

case $STACK in
    all)
        # If the selected stack is 'all', loop through all stacks
        for stack in "${STACKS[@]}"; do
            echo "Managing $stack stack..."
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