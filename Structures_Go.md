# 🧱 Comprendre les Structures (`struct`) en Go 

---

## 💬 Imagine la scène

Tu veux ranger des informations sur plusieurs personnes :  
leur nom, leur âge, leur ville.  
Tu pourrais créer plusieurs variables :

```go
nom := "titi"
age := 25
ville := "Nice"
```

Mais si tu veux faire ça pour 10 personnes, ton code devient vite le chaos 😵

💡 En Go, une **structure (`struct`)** sert à **ranger plusieurs informations ensemble** dans une seule “fiche”.

---

## 🎯 1. Créer une structure

```go
type Personne struct {
    Nom   string
    Age   int
    Ville string
}
```

🧠 Ici on crée un **nouveau type de donnée** appelé `Personne`.  
C’est un “modèle” qui dit à Go :  
> “Une personne, c’est toujours un nom, un âge et une ville.”

---

## 🧾 2. Créer une fiche à partir du modèle

```go
firas := Personne{
    Nom:   "titi",
    Age:   25,
    Ville: "Nice",
}
```

💬 Ici :
- `titi` est une variable de type `Personne`
- On a rempli chaque champ avec des valeurs concrètes

---

## 👀 3. Lire ou modifier les informations

### Lire :
```go
fmt.Println(firas.Nom)
fmt.Println(firas.Age)
```

### Modifier :
```go
firas.Age = 26
fmt.Println("Nouvel âge :", firas.Age)
```

---

## 👨‍👩‍👧‍👦 4. Plusieurs structures dans une liste

Tu peux créer plusieurs personnes et les ranger ensemble :

```go
personnes := []Personne{
    {"titi", 25, "Nice"},
    {"Ranya", 22, "Cagnes"},
    {"Kantin", 23, "Antibes"},
}

for _, p := range personnes {
    fmt.Println(p.Nom, "a", p.Age, "ans et habite à", p.Ville)
}
```

💡 Très pratique pour gérer une liste d’objets similaires (personnes, voitures, produits, etc.).

---

## 🚗 5. Autre exemple : une voiture

```go
type Voiture struct {
    Marque  string
    Vitesse int
}

func main() {
    v1 := Voiture{"Peugeot", 0}

    fmt.Println(v1.Marque, "roule à", v1.Vitesse, "km/h")

    v1.Vitesse = 50
    fmt.Println("Nouvelle vitesse :", v1.Vitesse)
}
```

---

## 📦 6. Lien avec ton projet **GoQuiz**

Dans ton projet **GoQuiz**, tu peux utiliser une `struct` pour **représenter chaque question** du quiz.

```go


type Question struct {
    Texte   string
    Reponse string
}

```

💡 Ici :
- `Question` est notre modèle
- Chaque question a un **texte** et une **réponse**

---


## 🧠 Résumé

| Élément | À quoi ça sert | Exemple |
|----------|----------------|----------|
| `struct` | Créer un modèle de données | `type Question struct {...}` |
| Instance | Créer une fiche  | `q := Question{"Texte", "Réponse"}` |
| Champ | Lire ou modifier une info | `q.Texte` |
| Liste | Plusieurs structs | `[]Question{...}` |

---

**********************************************

Les **structures** servent à :
- Regrouper plusieurs informations ensemble  
- Garder un code clair et organisé  
- Créer des modèles réutilisables (ex : `Personne`, `Question`, `Voiture`)  
- Rendre ton programme **plus logique et plus proche de la vraie vie**

---
