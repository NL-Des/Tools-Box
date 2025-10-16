## Pour construire etlancer le container, sans immobiliser le terminal, et supprimer les dockers qui tournaient avant.
docker compose up -d --build --remove-orphans --force-recreate <nom du container>

## Pour construire le container et ignorer le cache du dossier où se lance le docker-compose.yml
docker compose build --no-cache forum

## Affichage détaillé du processus de build du container.
docker compose build --progress=plain <nom du container>

## Pour démarrer un container.
docker start <nom ou id du container>

## Pour entrer dans le terminal du container.
docker exec -ti <nom du container> bash

## Pour arrêter un container actif.
docker stop <nom ou id du container>

## Pour supprimer un container spécifique.
docker kill <nom du container>

## Lister les container en cours d'exécutions.
docker ps

## Lister les container en cours d'exécutions ou non.
docker ps -a

## Pour vérifier les réseaux existants.
docker network ls

## Création du réseau web.
docker network create web

## Supression des réseaux web non utilisés.
docker network prune

## Supression des réseaux web utilisés ou non.
docker network prune -f

## Suppression conteneur non-actif.
docker rm <nom ou id du container>

## Suppresion conteneur actif.
docker rm -f <nom ou id du container>

## Pour voir les logs d'un container
docker logs <nom du container>

## Pour suivre les logs d'un container en temps réel.
docker logs -f <nom ou id du container>

## Pour se connecter au container.
docker exec -it <nom ou id du container> bash

## Destruction de toutes les images non-actives (intermédiaires incluses).
docker image prune -a

## Liste des images.
docker images ls

## Liste des images (intermédiaires comprises).
docker images -a

## Supprimer une image.
docker image rm <id ou TAG>