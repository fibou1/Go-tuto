# Les liens entre Go et le HTML

Le HTML peut contenir des **balises spéciales Go** que l’on peut utiliser directement dans le code HTML.

Ces balises sont entourées de doubles accolades : `{{ }}`.

Elles permettent :

* d’afficher des données provenant d’un serveur back-end Go directement dans le HTML ;
* de créer des conditions ou des boucles selon l’état des variables Go. Ces instructions se terminent toujours par `{{end}}`.

---

### Exemple :

```html
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
</head>
<body>
    {{ .Message }}
    {{if eq .Nb_visiteurs 5}}
        <p>Il y a 5 visiteurs sur le site !</p>
    {{end}}
</body>
</html>
```

> Le HTML affiche la valeur de la variable Go *Title* dans le titre de l’onglet et *Message* dans le corps de la page.
> Il n’affiche la balise `<p>` que lorsque la variable Go *Nb_visiteurs* est égale à **5**.

---

## Objectifs

Dans ce cours, nous allons détailler comment **envoyer des données Go (variables)** vers le HTML afin de les afficher directement sur le site.

---

## Comprendre l’initialisation d’une page HTML via Go

Au lancement de la page HTML `"/contact"`, le serveur va :

* Appeler la fonction Go **Contact** du package *controller* (situé dans un fichier séparé) ;
* Initialiser une requête HTTP (souvent de méthode **GET**) envoyée à cette fonction, afin de récupérer et afficher les données sur la page.

Une page HTML est initialisée via la fonction Go suivante :

```go
package router

import (
	"net/http"
	"controller"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes du site
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/contact", controller.Contact)

	return mux
}
```

Fonction *Contact* :

```go
package controller

import (
	"html/template"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Traitement du formulaire POST
	}

	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Envoie-nous un message 📩",
	}
	tmpl := template.Must(template.ParseFiles("template/contact.html"))
	tmpl.Execute(w, data)
}
```

---

## Comprendre l'envoi et la réception d'un formulaire HTML vers Go

### HTML vers Go

Pour communiquer du HTML vers Go, le serveur utilise des **requêtes HTTP**, sous deux formes :

* **Méthode GET**

  > Lors du lancement du serveur, la page initialise une requête pour récupérer le contenu et l’afficher à l’utilisateur.

* **Méthode POST**

  > Avec une balise `<form>`, l’utilisateur peut envoyer des données (texte, choix, etc.) au serveur via un bouton d’envoi.

Exemple de formulaire HTML :

```html
<form method="post" action="/contact">
    <label>Nom :</label><br>
    <input type="text" name="name"><br><br>
    <input type="hidden" name="nombre" value="5"><br><br>
    <button type="submit">Envoyer</button>
</form>
```

> L'attribut `method` définit la méthode de la requête (ici `POST`). L’attribut `action` définit la route Go vers laquelle les données seront envoyées.

---


#### Composition d’une balise `<form>`

* **Balise `<input>`**

  * `type` : définit le type de saisie.
  * `name` : définit le nom du champ envoyé.
  * `value` : définit la valeur par défaut ou envoyée.

  > Le type `text` est un champ de saisie qui contiendra la valeur saisie par l'utilisateur. La valeur sera toujours une **string**.

* **Balise `<button>`**

  * `type="submit"` : permet d'envoyer le formulaire vers le serveur Go.

#### Points importants

Un formulaire envoyé en POST nécessite :

* un `<input>` pour définir les données ;
* un `<button>` pour envoyer les données.

---


### Réception du formulaire par le serveur Go

Après l'envoi d'un formulaire depuis le HTML, la fonction Go (`/contact`) reçoit deux paramètres :

* `w` de type `http.ResponseWriter`
* `r` de type `*http.Request`

> `r` contient le formulaire envoyé par le client HTML. `w` est une interface qui permet d’écrire la réponse HTTP envoyée au HTML.

#### Exemple de traitement

```go
func Contact(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        nb := r.FormValue("nombre")
        // nb sera une string, à convertir en entier si nécessaire avec strconv.Atoi()
    }
}
```

#### Points importants

* Vérifier d'abord la méthode de la requête (GET/POST).
* Récupérer les données du formulaire avec `r.FormValue()`.
* La valeur récupérée est toujours une string.

---


### Envoi des données du serveur Go vers le HTML

Après avoir traité les données du formulaire, on peut renvoyer de nouvelles données vers le HTML.

#### Préparer les données

On peut utiliser :

* **Une structure** :

```go
type ContactData struct {
    Title   string
    Message string
}

data := ContactData{
    Title:   "Contact",
    Message: "Merci " + name + " pour ton message : " + msg,
}
```

* **Une map** :

```go
data := map[string]string{
    "Title":   "Contact",
    "Message": "Merci " + name + " pour ton message : " + msg,
}
```

---


#### Envoyer les données vers le HTML

```go
tmpl := template.Must(template.ParseFiles("template/contact.html"))
tmpl.Execute(w, data)
return // On termine ici pour ne pas exécuter la partie GET
```

---
    

### Réception des données dans le HTML

Pour afficher les données envoyées depuis Go :

```html
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    <p>{{ .Message }}</p>
</body>
</html>
```

> `{{ .Title }}` et `{{ .Message }}` affichent le contenu des variables Go.

---


### Utilisation des balises Go dans le HTML

```html
{{if eq .Nb_visiteurs 5}}
    <p>Il y a 5 visiteurs sur le site !</p>
{{end}}
```

---


#### Balises Go utiles

* `{{if CONDITION}} ... {{end}}` : condition simple
* `{{if CONDITION}} ... {{else}} ... {{end}}` : condition avec sinon
* `{{range VARIABLE}} ... {{end}}` : boucle sur une liste/array

#### Opérateurs logiques

* `eq` : égal à
* `ne` : différent de
* `lt` : inférieur à
* `gt` : supérieur à
* `le` : inférieur ou égal à
* `ge` : supérieur ou égal à
* `{{.VARIABLE}}` : affiche la valeur d'une variable Go

---


## Exemple complet de fonction Go gérant un formulaire de contact

```go
package controller

import (
    "html/template"
    "net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        msg := r.FormValue("msg")

        data := map[string]string{
            "Title":   "Contact",
            "Message": "Merci " + name + " pour ton message : " + msg,
        }
        tmpl := template.Must(template.ParseFiles("template/contact.html"))
        tmpl.Execute(w, data)
        return
    }

    data := map[string]string{
        "Title":   "Contact",
        "Message": "Envoie-nous un message 📩",
    }
    tmpl := template.Must(template.ParseFiles("template/contact.html"))
    tmpl.Execute(w, data)
}
```

---


## Conclusion

Nous avons vu comment envoyer des données depuis un serveur Go vers une page HTML, comment utiliser ces données dans le HTML avec les balises Go, et comment gérer les conditions et boucles pour afficher dynamiquement le contenu.

---

