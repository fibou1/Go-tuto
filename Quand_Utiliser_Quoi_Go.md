# 🧭 Quand utiliser quoi en Go ?  
*(Guide simple pour comprendre la logique de base)*

---

## 💡 Avant tout

Un bon programmeur ne pense pas d’abord au code.  
Il se demande :  
> “Qu’est-ce que je veux faire exactement ?”  

Ensuite seulement, il choisit **l’outil adapté** :
- une **variable** pour garder une donnée,  
- un **if** pour faire un choix,  
- une **boucle** pour répéter,  
- une **fonction** pour ranger le code,  
- une **struct** pour regrouper des infos.  

---

## 1️⃣ Les variables — quand tu veux stocker une information

> “Je veux garder une valeur en mémoire pour la réutiliser.”

```go
package main
import "fmt"

func main() {
    nom := "Firas"
    age := 25
    fmt.Println("Nom :", nom)
    fmt.Println("Âge :", age)
}
```

🧠 **Quand l’utiliser :**
- Tu veux **stocker** une donnée (nom, âge, score, prix…)
- Tu veux **réutiliser** cette donnée plus tard dans ton code

---

## 2️⃣ Les conditions (`if / else`) — quand tu veux faire un choix

> “Je veux que le programme fasse quelque chose selon une situation.”

```go
package main
import "fmt"

func main() {
    temperature := 28

    if temperature > 30 {
        fmt.Println("Il fait très chaud ☀️")
    } else if temperature > 20 {
        fmt.Println("Il fait bon 😎")
    } else {
        fmt.Println("Il fait froid ❄️")
    }
}
```

🧠 **Quand l’utiliser :**
- Tu veux tester **une condition**
- Tu veux exécuter **différentes actions selon les cas**

---

## 3️⃣ Les boucles (`for`) — quand tu veux répéter une action

> “Je veux faire la même chose plusieurs fois.”

```go
package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Tour numéro", i)
    }
}
```

🧠 **Quand l’utiliser :**
- Tu veux **répéter** une instruction un certain nombre de fois
- Tu veux **parcourir une liste de valeurs**

---

## 4️⃣ Les fonctions (`func`) — quand tu veux réutiliser une action

> “Je veux éviter de copier-coller du code.”

```go
package main
import "fmt"

func DireBonjour() {
    fmt.Println("Bonjour 👋")
}

func main() {
    DireBonjour()
    DireBonjour()
}
```

🧠 **Quand l’utiliser :**
- Tu veux **ranger un bloc de code**
- Tu veux **réutiliser la même action** plusieurs fois
- Tu veux rendre ton programme **plus clair et organisé**

---

## 5️⃣ Les structures (`struct`) — quand plusieurs informations vont ensemble

> “Je veux regrouper plusieurs infos qui décrivent la même chose.”

```go
package main
import "fmt"

type Personne struct {
    Nom  string
    Age  int
}

func main() {
    p1 := Personne{"Firas", 25}
    fmt.Println("Nom :", p1.Nom)
    fmt.Println("Âge :", p1.Age)
}
```

🧠 **Quand l’utiliser :**
- Tu veux **décrire un objet réel** (personne, voiture, produit, etc.)
- Tu veux **regrouper** plusieurs données liées entre elles

---

## 6️⃣ Les listes (`slice`) — quand tu veux plusieurs éléments du même type

> “J’ai plusieurs valeurs similaires à ranger.”

```go
package main
import "fmt"

func main() {
    nombres := []int{10, 20, 30, 40}

    for _, n := range nombres {
        fmt.Println("Valeur :", n)
    }
}
```

🧠 **Quand l’utiliser :**
- Tu veux **stocker plusieurs valeurs**
- Tu veux **parcourir** une série d’éléments

---

## 🧩 Résumé pratique

| Situation | Outil à utiliser | Exemple rapide |
|------------|------------------|----------------|
| Je veux garder une donnée | Variable | `nom := "Firas"` |
| Je veux choisir entre 2 actions | `if / else` | `if age > 18 {...}` |
| Je veux répéter quelque chose | `for` | `for i := 0; i < 5; i++` |
| Je veux regrouper des infos liées | `struct` | `type Personne struct {...}` |
| Je veux ranger un bloc de code | `func` | `func DireBonjour()` |
| Je veux plusieurs valeurs du même type | `slice` | `[]int{1,2,3}` |

---

## 💡 Comment raisonner avant de coder

Pose-toi toujours ces 5 questions simples :

1. **Est-ce que je dois stocker une valeur ?** → ➜ variable  
2. **Est-ce que je dois choisir entre deux cas ?** → ➜ if  
3. **Est-ce que je dois répéter quelque chose ?** → ➜ for  
4. **Est-ce que j’ai plusieurs infos liées entre elles ?** → ➜ struct  
5. **Est-ce que mon code se répète ?** → ➜ func  

---

## 🎓 En résumé

> - **Variable** → garder une info  
> - **If / else** → faire un choix  
> - **For** → répéter une action  
> - **Struct** → regrouper plusieurs infos  
> - **Func** → ranger et réutiliser du code  
> - **Slice** → gérer plusieurs valeurs  

---

💬 **Conseil final :**
> Commence toujours simple.  
> Si tu répètes, mets une boucle.  
> Si tu choisis, mets un if.  
> Si tu regroupes, mets une struct.  
> Si tu veux réutiliser, mets une fonction.  
>  
> 👉 C’est tout. C’est ça, la logique d’un vrai programmeur.
