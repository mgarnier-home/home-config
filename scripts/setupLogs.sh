
setupLogs() {
  # ssh into the server in params
  # update the /etc/docker/daemon.json file with the log driver
  # restart the docker daemon

  local sshUser=$1
  local serverIp=$2

  ssh -t $sshUser@$serverIp "
    if [ -f /etc/docker/daemon.json ]; then
      sudo mv -f /etc/docker/daemon.json /etc/docker/daemon.json.save
    fi
    echo '{
      \"log-driver\": \"syslog\",
      \"log-opts\": {
        \"syslog-address\": \"udp://$SYSLOG_IP:$SYSLOG_PORT\",
        \"tag\": \"DCMSG:{{.Name}}:{{.ID}}\"
      }
    }' | sudo tee /etc/docker/daemon.json
    sudo service docker restart
  "
  

}

set -o allexport
source ./.env

setupLogs $SSH_USER $BOREE_IP
setupLogs $SSH_USER $EUROS_IP
setupLogs $SSH_USER $NOTOS_IP
setupLogs $SSH_USER $ZEPHYR_IP
setupLogs $SSH_USER $ATHENA_IP


set +o allexport