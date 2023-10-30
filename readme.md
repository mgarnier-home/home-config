## Docker-configs

Ensemble de plusieurs fichiers docker-compose pour des projets personnels.
Divisés en plusieurs `stack`, chaque `stack` peut etre déployée sur un ou plusieurs `host`

### Avant de commencer

#### Host principal

- Vérifier que les scripts sont executable `chmod +x scripts/*.sh`
- Ajouter des lignes dans le fichier `~/.bashrc`
  - dossier de script `PATH` : `export PATH=$PATH:~/docker-configs/scripts/`
  - script d'autocomplete : `source ~/docker-configs/scripts/autocomplete.sh`
- Installer docker
- Installer node 18
- Vérifier que les différents hosts sont bien accessibles en ssh

#### Host secondaire

- Installer docker
- Ajouter la clé publique de l'host principal dans authorized_keys

### Ajouter un nouvel host

- Ajouter deux lignes dans le fichier `.env`
  - NEWHOST_IP=a.b.c.d
  - NEWHOST_HOSTNAME=newhost
- Créer un fichier dans la stack choisie `compose/STACK/NEWHOST.STACK.yml`
- Ajouter les différents ports des nouveaux services dans le fichier `ports.env` avec les noms `NEWHOST_SERVICE_PORT`
- Pour activer `traefik-config` sur cet host, ajouter un nouvel item `host` dans le `config.json`
  - Penser à ajouter `annotations: traefik-ports: NEWHOST_SERVICE_PORT` pour chaque service qui requiert traefik
- Ajouter un nouveau contexte dans le script `scripts/addDockerContexts.sh`
  - `docker context create NEWHOST --docker "host=ssh://${SSH_USER}@${NEWHOST_IP}"`
  - Lancer le script `scripts/addDockerContexts.sh`
- Modifier les scripts `my-stack.sh` et `autocomplete.sh` pour y ajouter le nouvel host
