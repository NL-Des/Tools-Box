# Tableau Commandes Docker

---

## Construction & lancement

| Action | Commande |
|--------|----------|
| Construire et lancer un container sans bloquer le terminal, en supprimant les anciens | `docker compose up -d --build --remove-orphans --force-recreate <service>` |
| Construire un container sans utiliser le cache | `docker compose build --no-cache <service>` |
| Affichage détaillé du build | `docker compose build --progress=plain <service>` |
| Lancer un container depuis une image existante | `docker run -d --name <nom> -p <port_host>:<port_container> <repository>` |
| Redémarrer un service défini dans docker-compose | `docker compose restart <service>` |

---

## Démarrage / Arrêt / Suppression des containers

| Action | Commande |
|--------|----------|
| Démarrer un container | `docker start <nom/id>` |
| Arrêter un container | `docker stop <nom/id>` |
| Supprimer un container non actif | `docker rm <nom/id>` |
| Supprimer un container actif | `docker rm -f <nom/id>` |
| Tuer un container (kill) | `docker kill <nom>` |

---

## Interaction avec un container

| Action | Commande |
|--------|----------|
| Entrer dans le terminal du container | `docker exec -ti <nom> bash` |
| Se connecter au container (équivalent) | `docker exec -it <nom/id> bash` |
| Voir les logs | `docker logs <nom>` |
| Suivre les logs en temps réel | `docker logs -f <nom/id>` |

---

## Inspection & informations

| Action | Commande |
|--------|----------|
| Lister les containers actifs | `docker ps` |
| Lister tous les containers | `docker ps -a` |
| Voir les volumes utilisés par un container | `docker inspect <nom_du_container>` |
| Lister les images | `docker images ls` |
| Lister toutes les images (intermédiaires incluses) | `docker images -a` |

---

## Nettoyage

| Action | Commande |
|--------|----------|
| Supprimer une image | `docker image rm <id/tag>` |
| Supprimer toutes les images non actives | `docker image prune -a` |

---

## Volumes

| Action | Commande |
|--------|----------|
| Lister les volumes | `docker volume ls` |
| Inspecter un volume | `docker volume inspect <volume>` |
| Supprimer un volume | `docker volume rm <volume>` |

---

## Réseaux

| Action | Commande |
|--------|----------|
| Lister les réseaux | `docker network ls` |
| Créer un réseau | `docker network create <nom>` |
| Supprimer un réseau Docker | `docker network rm <nom>` |
| Supprimer réseaux Docker non utilisés (retour à la configuration de base de Docker) | `docker network prune` |
| Forcer la suppression des réseaux utilisés ou non (retour à la configuration de base de Docker) | `docker network prune -f` |
| 

---