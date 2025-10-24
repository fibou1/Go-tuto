# GoQuiz

GoQuiz est un quiz éducatif en ligne de commande écrit en Go. Il permet de réviser plusieurs thématiques liées à l'informatique, à la cybersécurité et à la data.

## Fonctionnalités

- Sélection d'un thème parmi trois catégories : Informatique générale, Cybersécurité, Data & IA
- Questions à choix multiples avec validation instantanée
- Attribution de titres en fonction du score final
- Easter egg lorsque l'utilisateur tente une commande secrète
- Enregistrement et affichage des meilleurs scores dans un fichier JSON

## Prérequis

- Go 1.21 ou version supérieure

## Installation

```bash
git clone <votre-url-depot>
cd Go-tuto
go run .
```

## Utilisation

1. Lancer le jeu avec `go run .`
2. Saisir un nom ou pseudo
3. Choisir un thème parmi la liste proposée
4. Répondre aux questions en tapant la lettre correspondant à votre choix
5. Consulter votre titre et votre classement

Le fichier `scores.json` est créé automatiquement pour stocker vos meilleurs scores.

## Exemple de sortie

```
============================
        GoQuiz v1.0
Apprends et révise en t'amusant !
============================

Quel est ton prénom ou pseudo ? Alex
Choisis un thème :
1. Informatique générale
2. Cybersécurité
3. Data & IA
Tape 'help' pour obtenir de l'aide.

Entre le numéro de la catégorie : 2

--- CYBERSÉCURITÉ ---

Question 1: Quel est le but d'un pare-feu ?
  A) Sauvegarder les données
  B) Empêcher l'accès non autorisé
  C) Accélérer le réseau
  D) Analyser les virus
Ta réponse (A, B, C...) : b
✅ Correct !
...
```

## Hooks Git (optionnel)

Pour ajouter un hook `pre-commit` simple qui exécute `gofmt`, créez le fichier `.git/hooks/pre-commit` avec le contenu suivant :

```bash
#!/bin/sh
go fmt ./...
```

Rendez le script exécutable :

```bash
chmod +x .git/hooks/pre-commit
```

## Licence

MIT

