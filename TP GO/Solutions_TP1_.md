# ✅ Solutions des TP Progressifs en Go

---

## 🔹 TP1 — Dire bonjour
```go
package main
import "fmt"

func main() {
    fmt.Println("Bonjour ! Bienvenue dans le monde de Go 🚀")
}
```
🧠 **Explication :**
- `fmt.Println()` affiche du texte à l’écran.

---

## 🔹 TP2 — Lire une donnée au clavier
```go
package main
import "fmt"

func main() {
    var nom string
    fmt.Print("Quel est ton prénom ? ")
    fmt.Scan(&nom)
    fmt.Println("Enchanté,", nom, "!")
}
```
🧠 **Explication :**
- `fmt.Scan(&nom)` lit la saisie de l’utilisateur.

---

## 🔹 TP3 — Condition if / else
```go
package main
import "fmt"

func main() {
    var age int
    fmt.Print("Quel âge as-tu ? ")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("Tu es majeur ✅")
    } else {
        fmt.Println("Tu es mineur ❌")
    }
}
```
🧠 **Explication :**
- `if` vérifie la condition `age >= 18`.

---

## 🔹 TP4 — Boucle for
```go
package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Tour numéro", i)
    }
}
```
🧠 **Explication :**
- La boucle `for` répète le code 5 fois.

---

## 🔹 TP5 — Addition de plusieurs nombres
```go
package main
import "fmt"

func main() {
    var total int
    for i := 1; i <= 3; i++ {
        var n int
        fmt.Print("Entre un nombre : ")
        fmt.Scan(&n)
        total += n
    }
    fmt.Println("La somme totale est :", total)
}
```
🧠 **Explication :**
- On additionne 3 valeurs lues dans la boucle.

---

## 🔹 TP6 — Moyenne de notes
```go
package main
import "fmt"

func main() {
    var n1, n2, n3 float64
    fmt.Println("Entre trois notes :")
    fmt.Scan(&n1, &n2, &n3)

    moyenne := (n1 + n2 + n3) / 3

    if moyenne >= 10 {
        fmt.Println("Réussi ✅ (moyenne =", moyenne, ")")
    } else {
        fmt.Println("Échoué ❌ (moyenne =", moyenne, ")")
    }
}
```
🧠 **Explication :**
- Utilise `float64` pour les nombres à virgule.  
- Compare `moyenne >= 10`.

---

## 🔹 TP7 — Liste et boucle range
```go
package main
import "fmt"

func main() {
    fruits := []string{"Pomme", "Banane", "Fraise"}
    for _, fruit := range fruits {
        fmt.Println("J’aime les", fruit)
    }
}
```
🧠 **Explication :**
- `range` parcourt la liste.  
- `_` ignore l’index.

---

## 🔹 TP8 — Struct : fiche d’étudiant
```go
package main
import "fmt"

type Etudiant struct {
    Nom     string
    Age     int
    Moyenne float64
}

func main() {
    e1 := Etudiant{"Firas", 25, 14.5}
    fmt.Println("Nom :", e1.Nom)
    fmt.Println("Âge :", e1.Age)
    fmt.Println("Moyenne :", e1.Moyenne)
}
```
🧠 **Explication :**
- La `struct` regroupe plusieurs infos liées.

---

## 🔹 TP9 — Fonction simple
```go
package main
import "fmt"

func DireBonjour(nom string) {
    fmt.Println("Bonjour", nom, "!")
}

func main() {
    DireBonjour("Firas")
    DireBonjour("Ranya")
}
```
🧠 **Explication :**
- Crée une fonction réutilisable.

---

## 🔹 TP10 — Mini calculatrice
```go
package main
import "fmt"

func addition(a, b float64) float64 {
    return a + b
}

func main() {
    var x, y float64
    fmt.Println("Entrez deux nombres :")
    fmt.Scan(&x, &y)
    fmt.Println("La somme est :", addition(x, y))
}
```
🧠 **Explication :**
- Combine lecture, fonction, calcul et affichage.
