FROM golang:1.22.3-alpine3.18 as builder

RUN apk add curl

WORKDIR /build

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

COPY home-cli/ home-cli/
COPY taskfile.yml .

RUN ls -la

RUN task build

FROM jamesread/olivetin:latest as runtime

USER root

RUN microdnf install -y --nodocs --noplugins --setopt=keepcache=0 --setopt=install_weak_deps=0 \
  ansible rsync

COPY --from=builder /build/bin/home-cli /usr/bin/home-cli

COPY ansible/ansible.cfg /etc/ansible/ansible.cfg
COPY ansible/roles/compose-v2 /etc/ansible/roles/compose-v2
COPY ansible/playbooks/compose.playbook.yml /etc/ansible/playbooks/compose.playbook.yml
COPY ansible/inventories/tailscale_hosts.yml /etc/ansible/inventories/tailscale_hosts.yml

ENV ANSIBLE_DIR /etc/ansible
ENV COMPOSE_DIR /compose
ENV OLIVETIN_CONFIG_DIR /config

ENTRYPOINT ["/bin/sh", "-c", "/usr/bin/home-cli build-olivetin && /usr/bin/OliveTin"] 
