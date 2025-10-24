# 🧠 Syntaxe de base du langage Go (Golang)

---

## 🎯 Objectif

> Comprendre comment est structuré un programme Go :  
> les mots-clés essentiels, les packages, les imports, les fonctions, et le rôle du fichier `go.mod`.

---

## 📦 1. Structure générale d’un programme Go

Un programme Go est **organisé** en trois parties principales :

1. Le **package** → le "groupe" ou module auquel appartient le fichier  
2. Les **imports** → les bibliothèques externes à utiliser  
3. Les **fonctions** → les blocs d’instructions exécutables

### 🧩 Exemple minimal :

```go
package main          // 1️⃣ Déclare le package principal
import "fmt"          // 2️⃣ Importe une bibliothèque standard

func main() {         // 3️⃣ Fonction principale (point d’entrée)
    fmt.Println("Hello, Go!")  // 4️⃣ Affiche un message
}
```

---

## 🧱 2. Explication ligne par ligne

### `package main`
- Tous les fichiers Go doivent **appartenir à un package**.  
- Le mot `main` indique que **ce package contient le point d’entrée** du programme.
- Si ton package s’appelait autrement (ex. `utils`), il ne pourrait pas être exécuté seul.

💡 **En résumé :**
> `package main` = "Ceci est mon fichier principal, exécutable."

---

### `import "fmt"`
- `import` permet d’utiliser un **package externe** (déjà inclus dans Go).  
- `"fmt"` signifie *format* → il contient des fonctions d’entrée/sortie :
  - `Print()` → affiche sans retour à la ligne  
  - `Println()` → affiche avec retour à la ligne  
  - `Scan()` → lit une saisie clavier

💡 **Exemple :**
```go
fmt.Print("Entre ton nom : ")
fmt.Scan(&nom)
fmt.Println("Bonjour", nom)
```

> ⚙️ `fmt` fonctionne comme une **boîte à outils** pour afficher et lire des valeurs.

---

### `func main() { ... }`
- `func` = mot-clé pour **définir une fonction**
- `main()` = **fonction principale exécutée au lancement**
- `{ ... }` = délimite le **bloc d’instructions**

💡 **Exemple visuel :**
```go
func main() {
    // Tout ce code sera exécuté au démarrage
}
```

> 🎯 C’est l’équivalent du **"quand drapeau vert cliqué"** de Scratch.

---

## 🧰 3. Les éléments de syntaxe de base

---

### 🧩 Déclaration de variable

#### Forme courte (recommandée)
```go
nom := "Firas"
age := 25
```
➡️ Go déduit automatiquement le type (`string`, `int`, etc.)

#### Forme longue (optionnelle)
```go
var nom string = "Firas"
var age int = 25
```

💡 `var` = mot-clé pour déclarer une variable manuellement.

---

### 🔢 Types de base

| Type Go | Exemple | Description |
|----------|----------|-------------|
| `string` | `"Hello"` | Chaîne de caractères |
| `int` | `42` | Nombre entier |
| `float64` | `3.14` | Nombre décimal |
| `bool` | `true` ou `false` | Valeur booléenne |
| `rune` | `'A'` | Caractère unique |
| `array` / `slice` | `[3]int{1,2,3}` | Tableau ou liste dynamique |
| `map` | `map[string]int{"Alice":10}` | Dictionnaire clé/valeur |

---

### 🧠 Constantes

Les **constantes** ne changent jamais :
```go
const Pi = 3.14159
```

---

### 🔀 Conditions (if / else)

```go
if age >= 18 {
    fmt.Println("Tu es majeur ✅")
} else {
    fmt.Println("Tu es mineur ❌")
}
```

> 🧩 Même logique que le bloc “si / sinon” en Scratch.

---

### 🔁 Boucles (for)

#### Répéter un certain nombre de fois :
```go
for i := 1; i <= 5; i++ {
    fmt.Println("Tour", i)
}
```

#### Boucle infinie :
```go
for {
    fmt.Println("Je tourne sans fin 😵")
}
```

---

### ⚙️ Fonctions

Créer un bloc réutilisable :

```go
func Bonjour(nom string) {
    fmt.Println("Bonjour", nom)
}

func main() {
    Bonjour("Ynov")
}
```

💡 `func` = mot-clé pour **créer une fonction**  
🧠 C’est l’équivalent du bloc “Créer un bloc” dans Scratch.

---

## 🧩 4. Packages et imports en détail

---

### 🔹 Qu’est-ce qu’un package ?
Un **package** = un dossier ou groupe de fichiers qui partagent un même nom logique.

Exemple :
```
monprojet/
├── main.go        (package main)
└── utils/
    └── utils.go   (package utils)
```

**utils.go :**
```go
package utils

func Bonjour(nom string) string {
    return "Bonjour " + nom
}
```

**main.go :**
```go
package main
import (
    "fmt"
    "monprojet/utils"
)

func main() {
    message := utils.Bonjour("Firas")
    fmt.Println(message)
}
```

💡 `import "monprojet/utils"` → permet d’utiliser les fonctions de `utils.go`

---

### 🔹 Pourquoi utiliser plusieurs packages ?
- Pour **organiser ton code** (comme des chapitres de livre)
- Pour **réutiliser des fonctions**
- Pour **garder ton code clair et lisible**

---

## 📘 5. Le fichier `go.mod` — le cœur du projet Go

---

### 📦 Qu’est-ce que `go.mod` ?
C’est le **fichier de configuration** du projet Go.  
Il indique :
- Le **nom du module** (ton projet)
- La **version de Go utilisée**
- Les **dépendances** importées

### 💡 Exemple :

```txt
module monprojet
go 1.22

require (
    github.com/gin-gonic/gin v1.9.0
)
```

---

### 🧠 À quoi ça sert ?
- Identifier ton projet comme un **module Go**  
- Suivre les **versions des dépendances**  
- Permettre à `go run` et `go build` de fonctionner correctement

---

### ⚙️ Comment le créer ?
1. Ouvre ton terminal dans le dossier du projet  
2. Tape :
```bash
go mod init monprojet
```
3. Go crée un fichier `go.mod` à la racine.

> 🔁 Chaque fois que tu fais un `import`, Go met à jour automatiquement les dépendances.

---

## 🧠 6. Schéma global d’un projet Go

```
monprojet/
├── go.mod
├── main.go
└── utils/
    └── utils.go
```

### main.go :
```go
package main
import (
    "fmt"
    "monprojet/utils"
)

func main() {
    fmt.Println("Programme Go démarré 🚀")
    message := utils.Bonjour("Firas")
    fmt.Println(message)
}
```

### utils/utils.go :
```go
package utils

func Bonjour(nom string) string {
    return "Bonjour " + nom
}
```

🧠 **Structure typique :**
- `package main` → exécutable principal  
- `package utils` → fonctions annexes  
- `go.mod` → configuration du projet  

---

## 🎓 En résumé

| Élément | Exemple | Rôle |
|----------|----------|------|
| `package main` | | Point d’entrée du programme |
| `import "fmt"` | | Utiliser une bibliothèque |
| `func main()` | | Démarrage du code |
| `var` / `:=` | | Déclaration de variables |
| `if / else` | | Condition logique |
| `for` | | Boucle répétitive |
| `func` | | Créer une fonction réutilisable |
| `go.mod` | | Gérer le projet et ses dépendances |

---

> 💡 **En un mot :**  
> Le langage Go est comme Scratch, mais sans les blocs :  
> tu dis à la machine **quoi faire**, **dans quel ordre**, et **avec quelles règles**.
https://www.youtube.com/watch?v=t8b9f5M9yoY