# 📦 Workflow Git Complet — Branches de Fonctionnalité

> Proposé par Emrick Rivet - B1 Info - Octobre 2025

---

## Objectif

Ce guide détaille un workflow Git complet pour travailler avec des branches de fonctionnalité. Il couvre la création, le travail et la fusion de branches dans le cadre d'un projet collaboratif.

---

## 🧩 1. Création ou clonage d'une branche

### Créer une branche

* 1️⃣ Se placer sur la branche principale :

```git
git checkout main
```

> Pour s'assurer d'être sur la branche main.

```git
git pull origin main
```

> Récupérer les dernières mises à jour de la branche main depuis le dépôt distant.

* 2️⃣ Créer une branche pour une nouvelle fonctionnalité :

```git
git checkout -b nom_de_branche
```

> L’option `-b` permet de créer la branche et d'y basculer en une seule commande.

* Exemple :

```git
git checkout -b feature/nom-de-fonctionnalité
```

---

### Cloner une branche existante après avoir cloné un dépôt GitHub

* 1️⃣ Lister les branches disponibles sur le dépôt distant :

```git
git branch -a
```

> Les branches distantes apparaissent en rouge.

* Exemple de résultat :

```
remotes/origin/feature/les_bases
```

> Ici, la branche `feature/les_bases` est disponible sur le dépôt distant.

* 2️⃣ Cloner la branche en local :

```git
git checkout nom_de_branche
```

> Cette commande place automatiquement sur la branche choisie.

* Exemple :

```git
git checkout feature/les_bases
```

---

## 🧰 2. Travailler sur la branche

* 1️⃣ Ajouter ou modifier des fichiers :

```git
git add .
```

> Ajoute tous les fichiers modifiés.

```git
git commit -m "description_du_commit"
```

> Crée un commit avec les fichiers ajoutés.

* 2️⃣ Pousser la branche sur GitHub (⚠️ Premier push) :

```git
git push -u origin feature/nom-de-fonctionnalité
```

> L’option `-u` (ou `--set-upstream`) lie la branche locale à la branche distante pour simplifier les futurs `push` et `pull`.

* 3️⃣ Récupérer les mises à jour si d’autres ont travaillé sur la branche :

```git
git pull origin feature/nom-de-fonctionnalité
```

* Si des changements ont été faits sur la branche main :

```git
git checkout main
git pull origin main
```

* Puis intégrer ces changements dans votre branche :

```git
git checkout feature/nom-de-fonctionnalité
git merge main
```

> Intègre les derniers changements de la branche main dans votre branche.

---

## 🔀 3. Fusion de la branche vers main

* 1️⃣ Récupérer la dernière version de la branche main :

```git
git checkout main
git pull origin main
```

* 2️⃣ Fusionner la branche de fonctionnalité dans main :

```git
git merge feature/nom-de-fonctionnalité
```

> On merge toujours dans la branche sur laquelle on se trouve, ici `main`.

* 3️⃣ Résoudre les conflits si nécessaire :

> Si des conflits apparaissent, Git le signalera. Il faudra alors les résoudre manuellement dans les fichiers concernés.

* 4️⃣ Commit du merge :

```git
git commit -m "Merge de feature/nom-de-fonctionnalité"
```

> Si des conflits ont été résolus manuellement, il faut effectuer un commit pour finaliser la fusion.

* ⚠️ Git crée souvent un commit automatique s’il n’y a pas de conflit. Il peut afficher un message dans l’éditeur (type Vim) :

```
Merge branch 'MenuPrincipale'
# Please enter a commit message to explain why this merge is necessary,
# especially if it merges an updated upstream into a topic branch.
# Lines starting with '#' will be ignored, and an empty message aborts the commit.
.git/MERGE_MSG[+] [unix] (14:59 21/10/2025)
```

* 📝 Modifier le nom du commit dans Vim :

  * Appuyer sur la touche "i" pour passer en mode insertion.
  * Modifier le texte du commit dans la première ligne.
  * Appuyer sur "Échap" pour sortir du mode insertion.
  * Taper `:wq` puis appuyer sur "Entrée" pour enregistrer et quitter.

* 5️⃣ Pousser la branche principale mise à jour :

```git
git push origin main
```

* 6️⃣ Supprimer la branche une fois le merge effectué :

  * Localement :

```git
git branch -d feature/nom-de-fonctionnalité
```

* Sur GitHub (dépôt distant) :

```git
git push origin --delete feature/nom-de-fonctionnalité
```

---

## Conclusion

Ce workflow Git permet de gérer efficacement le développement de nouvelles fonctionnalités en utilisant des branches. En suivant ces étapes, vous pouvez travailler de manière collaborative tout en maintenant une base de code propre et organisée.

✅ Fin du Workflow

---