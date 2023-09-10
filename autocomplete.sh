#!/bin/bash

_my_script_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    case ${COMP_CWORD} in
        1)
            opts="deploy undeploy redeploy"
            ;;
        2)
            opts="all samba network monitoring home minecraft backup"
            ;;
        3)
            opts="all athena apollon hermes artemis"
            ;;
        *)
            opts=""
            ;;
    esac
    
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}
complete -F _my_script_completion ./stack.sh