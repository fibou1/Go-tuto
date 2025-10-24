#  TP Progressifs en Go (Débutants)
```
# Essayez de réaliser ces exercices sans l’aide d’une IA STP .
Faites appel à votre mémoire, à votre logique et à tout ce que vous avez déjà appris 
— c’est le meilleur moyen de progresser vraiment
```


---

## 🔹 TP1 — Dire bonjour 
###  Objectif :
Découvrir comment afficher un texte à l’écran.

###  Syntaxe utile :
```go
fmt.Println("texte ici")
```

###  Indice :
- Utilise `fmt.Println()` pour afficher un message.

---

## 🔹 TP2 — Lire une donnée au clavier 
###  Objectif :
Demander un prénom à l’utilisateur et l’afficher.

###  Syntaxe utile :
```go
fmt.Scan(&variable)
```

###  Indice :
- Crée une variable pour stocker le nom.
- Utilise `fmt.Print()` pour poser la question.

---

## 🔹 TP3 — Les conditions `if / else`
###  Objectif :
Demander l’âge d’une personne et dire si elle est majeure ou mineure.

###  Syntaxe utile :
```go
if condition {
    // action
} else {
    // autre action
}
```

###  Indice :
- Utilise `>=` pour comparer.
- Si l’âge est ≥ 18 → majeur, sinon → mineur.

---

## 🔹 TP4 — Boucle `for` : répéter une action
###  Objectif :
Afficher “Tour numéro X” 5 fois.

###  Syntaxe utile :
```go
for i := 1; i <= 5; i++ {
    // ton code ici
}
```

###  Indice :
- Utilise une boucle avec un compteur.

---

## 🔹 TP5 — Addition de plusieurs nombres 
###  Objectif :
Demander 3 nombres et calculer leur somme.

###  Syntaxe utile :
```go
total += n
```

###  Indice :
- Crée une variable `total` pour additionner.
- Utilise une boucle pour répéter la saisie 3 fois.

---

## 🔹 TP6 — Moyenne de notes 🎓
###  Objectif :
Calculer la moyenne de 3 notes et afficher si c’est réussi (≥ 10).

###  Syntaxe utile :
```go
moyenne := (n1 + n2 + n3) / 3
```

###  Indice :
- Utilise `float64` pour les nombres à virgule.
- Fais une condition pour afficher “Réussi” ou “Échoué”.

---

## 🔹 TP7 — Liste et boucle `range`
###  Objectif :
Afficher une liste de fruits avec une phrase du type “J’aime les ...”.

###  Syntaxe utile :
```go
for _, element := range liste {
    // ...
}
```

###  Indice :
- Crée une liste `[]string`.
- Parcours-la avec `range`.

---

## 🔹 TP8 — Struct simple : une fiche d’étudiant
###  Objectif :
Créer un modèle `Etudiant` avec nom, âge et moyenne, puis afficher ses infos.

###  Syntaxe utile :
```go
type NomStruct struct {
    Champ1 Type
    Champ2 Type
}
```

###  Indice :
- Crée une `struct` et une variable du type `Etudiant`.
- Affiche chaque champ avec `fmt.Println()`.

---

## 🔹 TP9 — Fonction simple
###  Objectif :
Créer une fonction `DireBonjour()` qui affiche un message personnalisé.

###  Syntaxe utile :
```go
func NomFonction(parametre Type) {
    // action
}
```

###  Indice :
- Crée une fonction avant `main()`.
- Appelle-la deux fois avec des prénoms différents.

---

## 🔹 TP10 — Mini calculatrice 
###  Objectif :
Lire deux nombres, les additionner et afficher le résultat avec une fonction.

###  Syntaxe utile :
```go
func addition(a, b float64) float64 {
    return a + b
}
```

###  Indice :
- Crée une fonction `addition`.
- Demande deux nombres puis affiche le résultat.

---

🎓 **Conseil :**
Fais les TP dans l’ordre, un par jour si possible.  
Lis bien les indices avant de coder, et teste chaque ligne avant de passer à la suivante.
