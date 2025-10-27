# 📦 Workflow Git complet — Branche de fonctionnalité
> Proposé par Emrick Rivet - B1 Info - Octobre 2025

---

## Objectif
Ce guide détaille un workflow Git complet pour travailler avec des branches de fonctionnalité. Il couvre la création, le travail, et la fusion de branches dans le cadre d'un projet collaboratif.

---

## 🧩 1. Création de la branche ou cloner une branche d'un distant

### Créer une branche: 
* 1️⃣ Se placer sur la branche principale
    ```git
	git checkout main  
    ```    
    > Pour être sûr d'être sur la branche main

    ```git
     git pull origin main  
    ```
    > récupérer les dernières mises à jour du main sur le dépôt distant

* 2️⃣ Créer une branche pour une nouvelle fonctionnalité
    ```git
	git checkout -b nom_de_branche  
    ```
    > -b permet de créer la branche en même temps du checkout

    * Exemple : 
    ```git
    git checkout -b feature/nom-de-fonctionnalité
    ```

---

### Cloner une branche après avoir cloner un repositoire GitHub:
* 1️⃣ Regarder les branche disponibles sur le distant:
    ```git
	git branch -a 
    ```	
    > En rouge les branches disponibles sur le distant
	
    * Exemple de résultat : 
    remotes/origin/feature/les_bases 	
    > Ici on a une branche "feature/les_bases" sur le distant
			
* 2️⃣ Cloner la branche dans son local:
    ```git
	git checkout nom_de_branche 
    ```	
    > On se déplace automatiquement dans la branche

    * Exemple : 
    ```git
    git checkout feature/les_bases
    ```

---

## 🧰 2. Travailler sur la branche
* 1️⃣ Ajouter ou modifier des fichiers
	```git
    git add . 
    ```
    > ajoute tous les fichiers modifiés
    ```git
	git commit -m "description_du_commit"
    ```
    > crée un commit avec les fichiers ajoutés

* 2️⃣ Pousser la branche sur GitHub (⚠️ Premier push)
    ```git
	git push -u origin feature/nom-de-fonctionnalite
    ```
    > L’option -u (ou --set-upstream) lie la branche locale à la branche distante pour simplifier les futurs push / pull.

* 3️⃣ Récupérer des mises à jour si d’autres ont travaillé sur la branche
	* Sur ta branche locale :
    ```git
	 git pull origin feature/nom-de-fonctionnalite
     ```

	* Si des changements ont été faits sur le main :
    ```git
	git checkout main
	git pull origin main
    ```

	* Puis intégrer ces changements dans ta branche :
    ```git
	git checkout feature/nom-de-fonctionnalite
	git merge main    
    ```
    > intègre les derniers changements du main dans ta branche

---

## 🔀 3. Merge de la branche vers main
* 1️⃣ Récupérer la dernière version du main
    ```git
	git checkout main
	git pull origin main
    ```

* 2️⃣ Fusionner la branche de fonctionnalité dans main
    ```git
	git merge feature/nom-de-fonctionnalite
    ```
    > On merge toujours dans la branche sur laquelle on se trouve, ici main.

* 3️⃣ Résoudre les conflits si nécessaire
	> Si des conflits apparaissent, Git te le signalera. 
    > Il faudra alors les résoudre manuellement dans les fichiers concernés.
	
* 4️⃣ Commit du merge
    ```git
	git commit -m "Merge de feature/nom-de-fonctionnalite"
    ```
    > Si des conflits ont été résolus manuellement, il faut faire un commit pour finaliser la fusion.

	* ⚠️ Git crée souvent un commit automatique s’il n’y a pas de conflit.
        > Il peut afficher un message dans l’éditeur (type Vim) :
        ```vim
        Merge branch 'MenuPrincipale'								
        # Please enter a commit message to explain why this merge is necessary, 		
        # especially if it merges an updated upstream into a topic branch. 			
        # Lines starting with '#' will be ignored, and an empty message aborts the commit..	
        .git/MERGE_MSG[+] [unix] (14:59 21/10/2025)						
        ```

		* 📝 Modifier le nom du commit dans Vim :

			- Appuyer sur la touche "i" pour passer en mode insertion

			- Modifier le texte du commit dans la première ligne (souvent en orange)

			- Appuyer sur "Échap" pour sortir du mode insertion

			- Taper `:wq` (write and quit) puis appuyer sur "Entrée" pour enregistrer et quitter

* 5️⃣ Pousser la branche principale mise à jour
    ```git
	git push origin main
    ```

* 6️⃣ Supprimer la branche une fois merge
	* Localement:
    ```git
	git branch -d feature/nom-de-fonctionnalite
    ```
	* Sur GitHub (dépôt distant) :
    ```git
	git push origin --delete feature/nom-de-fonctionnalite
    ```

---

## Conclusion

Ce workflow Git permet de gérer efficacement le développement de nouvelles fonctionnalités en utilisant des branches. En suivant ces étapes, tu peux travailler de manière collaborative tout en maintenant une base de code propre et organisée.

✅ Fin du Workflow

---