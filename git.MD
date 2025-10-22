Taper Q pour quitter la commande en cours.

# Initialiser un dépôt git local.
git init

# Procédure d'envoi des push.
git add <.> <nom_du_fichier>
git commit -m "<Commentaire_du_code>"
git push

# Récupération et clonage.
## Cloner un dépôt distant sur ma machine.
git clone <url>

## Récupérer les dernières modifications du dépôt distant sans les fusionner avec le local.
git fetch

## Récupérer et fusionner les modifications du dépôt distant dans ma branche locale.
git pull

# Gestion des branches
## Création d'une branche locale dev, à partir d'une branche distante dev.
git checkout -b dev origin/dev

## Affichage des dépôts distants configurés.
git remote -v

## Liste des branches locales.
git branch

## Changer de branche locale.
git checkout <branche>

## Créer une branche à partir d'une branche existante, puis aller sur cette nouvelle branche.
git checkout -b <nouvelle-branche>

## Fusionner la branche actuelle avec celle qui est nommée.
git merge <branche>

# Suivi des fichiers.
## Affiche les fichiers modifiés, supprimés ou ajoutés différents entre la branche et le local.
git status

# Historique et inspection
## Historique des commits
git log

## Montre les différences entre l'index et les fichiers modifiés.
git diff

## Montre les détails d'un commit.
git show <commit>

# Suppression et nettoyage.
## Supprimer un fichier en local et dans le git.
git rm <fichier>

## Mise en tampon de modifications non commités.
git stash

## Retourner à l'état d'origine de la branche.
git restore

# Arborescence de l'architecture
## Tout afficher.
tree

## Arbre avec une limite de profondeur.
tree -L2

## Arbre avec les fichiers cachés.
tree -a