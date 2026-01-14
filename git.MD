# 📘 Tableau des commandes Git

> Taper **Q** pour quitter certaines commandes (ex : `git log`).

---

## 🏁 Initialisation

| Action | Commande |
|--------|----------|
| Initialiser un dépôt local | `git init` |

---

## 🚀 Procédure d’envoi (push)

| Action | Commande |
|--------|----------|
| Ajouter un fichier ou tout le dossier | `git add <.> <fichier>` |
| Créer un commit | `git commit -m "<message>"` |
| Envoyer vers le dépôt distant | `git push` |

---

## 📥 Récupération & clonage

| Action | Commande |
|--------|----------|
| Cloner un dépôt distant | `git clone <url>` |
| Récupérer les modifications distantes sans fusion | `git fetch` |
| Récupérer et fusionner les modifications distantes | `git pull` |

---

## 🌿 Gestion des branches

| Action | Commande |
|--------|----------|
| Créer une branche locale depuis une branche distante | `git checkout -b dev origin/dev` |
| Afficher les dépôts distants configurés | `git remote -v` |
| Lister les branches locales | `git branch` |
| Changer de branche | `git checkout <branche>` |
| Créer une nouvelle branche et s’y déplacer | `git checkout -b <nouvelle-branche>` |
| Fusionner la branche actuelle avec une autre | `git merge <branche>` |

---

## 📂 Suivi des fichiers

| Action | Commande |
|--------|----------|
| Voir les fichiers modifiés / ajoutés / supprimés | `git status` |

---

## 🕓 Historique & inspection

| Action | Commande |
|--------|----------|
| Afficher l’historique des commits | `git log` |
| Voir les différences entre index et fichiers modifiés | `git diff` |
| Voir les détails d’un commit | `git show <commit>` |

---

## 🧹 Suppression & nettoyage

| Action | Commande |
|--------|----------|
| Supprimer un fichier local + Git | `git rm <fichier>` |
| Mettre de côté des modifications non commités | `git stash` |
| Restaurer l’état d’origine d’un fichier ou branche | `git restore` |

---

## 🌳 Arborescence du projet (commande `tree`)

| Action | Commande |
|--------|----------|
| Afficher toute l’arborescence | `tree` |