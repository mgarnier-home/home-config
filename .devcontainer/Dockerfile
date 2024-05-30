FROM mcr.microsoft.com/devcontainers/base:bookworm

RUN apt-get update && apt-get install -y bash-completion iputils-ping && \
  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /root/.local/bin && \
  echo "if [ -f /etc/bash_completion ]; then \n . /etc/bash_completion \n fi" >> /root/.bashrc



CMD [ "sleep", "infinity" ]