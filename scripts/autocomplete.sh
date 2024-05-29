#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# # Read JSON file
# STACK_JSON_FILE="$SCRIPT_DIR/stack.json"
# if [ ! -f "$STACK_JSON_FILE" ]; then
#     echo "Error: $STACK_JSON_FILE not found."
#     exit 1
# fi

# ACTIONS=($(jq -r '.actions[]' $STACK_JSON_FILE | tr -d '\r'))
# HOSTS=($(jq -r '.hosts[]' $STACK_JSON_FILE | tr -d '\r'))
# STACKS=($(jq -r '.stacks[]' $STACK_JSON_FILE | tr -d '\r'))

_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    case ${COMP_CWORD} in
        1)
            opts="deploy undeploy redeploy pull"
            ;;
        2)
            opts="all backup file_server minecraft monitoring network nextcloud paperless plex samba db jellyfin home_assistant navidrome"
            ;;
        3)
            opts="all athena euros boree"
            ;;
        *)
            opts=""
            ;;
    esac
    
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}
complete -F _completion home-cli