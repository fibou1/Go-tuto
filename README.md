# 🧠 Projet GoQuiz — Séance de Soutien Programmation Go

## 🎯 Objectif pédagogique

Créer un **petit jeu de quiz en console** avec Go pour réviser les bases de programmation (variables, conditions, boucles, fonctions, etc.)  
Le jeu pose des questions à l’utilisateur, vérifie les réponses et affiche un score final avec un petit message amusant.

---

## 🗂️ Structure du projet

GoQuiz/
│
├── main.go → point d’entrée du programme
├── quizmenu.go → menu principal
├── quizinfo.go → quiz sur l’informatique
├── quizcyber.go → quiz cybersécurité
├── quizIAdata.go → quiz sur la data et l’IA
├── go.mod → module Go (nom du projet)
├── README.md → documentation du projet

---

## 🧩 Logique du programme (vue d’ensemble)

         ┌────────────────────────────┐
         │        main.go             │
         │ (point d'entrée du jeu)    │
         └────────────┬───────────────┘
                      │
                      ▼
         ┌────────────────────────────┐
         │       quizmenu.go          │
         │  fonction : ShowMenu()     │
         │                            │
         │  1️⃣ Affiche le menu       │
         │  2️⃣ Lis le choix          │
         │  3️⃣ Appelle le bon quiz   │
         └───────┬─────────┬───────┬─┘
                 │         │       │
     ┌───────────┘         └────┐  └────────────────────────────┐
     ▼                          ▼                               ▼
┌──────────────────────┐    ┌──────────────────────┐    ┌──────────────────────┐

│    quizinfo.go       │    │     quizcyber.go     │    │    quizIAdata.go     │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
│                      │    │                      │    │                      │
└──────────────────────┘    └──────────────────────┘    └──────────────────────┘