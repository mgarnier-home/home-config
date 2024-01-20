#!/bin/bash

_my_script_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    case ${COMP_CWORD} in
        1)
            opts="deploy undeploy redeploy pull"
            ;;
        2)
            opts="all backup file_server minecraft monitoring network nextcloud paperless palworld runner samba db home_assistant media syslog"
            ;;
        3)
            opts="all athena euros boree notos zephyr"
            ;;
        *)
            opts=""
            ;;
    esac
    
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}
complete -F _my_script_completion my-stack.sh