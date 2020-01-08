## Les packages

En golang, le code est organisé par package.
Ils correspondent à un dossier dans votre projet.  
On peut référencer du code dans un autre package via l'instruction `import`  

En golang, on travaille en général avec le code source, même pour les dépendances.  
On import du **code**, pas de librairie compilée. 

## Instruction import

On importe de la même manière les packages du projet courant, les packages standards ou d'un autre projet.

e.g.: 
```golang
import ()
	"fmt" // standard package
	"my_awsome_project/some_subpackage" // current project's sub-package
	"github.com/davecgh/go-spew" // another project
)
```

## La mauvaise mais simple façon de faire : dossier `src`

On peut utiliser la commande `go get` pour récupérer des sources et les placer dans le dossier `$GOPATH/src` où le compilateur ira chercher.  
C'est facile et rapide, mais ça pose de gros problème : tous les projets qu'on écrit partagent les mêmes dépendances et on n'a aucun contrôle sur la version utilisée (la dernière version disponible est utilisée.). 

## Exercice 1

Prenons cette application en exemple :
```golang
package main

import (
	"time"
	"github.com/cheggaaa/pb"
)

func main() {
	count := 5000
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}
```

Récupérez les dépendances dans votre dossier `$GOPATH/src` et exécutez l'application. 

## La bonne façon de faire : vendoring et package manager

#### Le package manager

Son rôle est de :
- lire les imports de votre projet
- d'y associer un fichier de config qui définissent les versions des dépendances
- générer un fichier `lock` qui garantis que les versions utilisées ne changent pas d'un build à l'autre
- générer le dossier `vendor` à partir du fichier `lock`

Il y a eu beaucoup de package manager dans le passé. 
On va tous les ignorer et se concentrer sur l'outil `go mod`.

#### RIP le dossier vendor

Il a été utilisé pendant quelques années pour stocker les dépendances résolues par les packages managers.   
Mais `go mod` ne génère pas de fichier dans le dossier `vendor`. Il suit un concept similaire mais choisi d'écrire dans le dosser `$GOPATH/pkg/mod`.  
Notez qu'il est toujours possible de générer le dossier `vendor` avec `go mod vendor`, mais même ainsi `go build` ignorera ce folder si on ne specifie par l'argument `go build -mod vendor`. 

Donc à partir de maintenant va ignorer le dossier `vendor` qui ne sert plus que comme failback pour le support de quelques vieux outils. 

#### Go mod

Depuis golang 1.11 `go mod` est disponible pour gérer les modules (pour l'instant comprenez module~=package). 
Depuis golang 1.13 il est utilisé par défaut quand on run `go get` et que le projet est **en mode module**. 

Il est officiel.  
Il est une solution long terme.  
Il est plus qu'un package manager, mais on va ignorer les autres fonctions pour l'instant.  
Il est le package manager le plus utilisé par la communauté.   

Pour initialiser le fichier de config et passer le projet **en mode module**, exécutez `go mod init`.  
Le fichier `go.mod` généré contiendra plus tard toutes les informations sur les dépendances du projet. Pour l'instant, il est plutôt vide.  

On peut remplir automatiquement les dépendances utilisées en exécutant `go get`.  
Maintenant que notre projet contient un fichier `go.mod`, le comportement de `go get` est altéré.  
Au lieu de placer les dépendances dans le dossier `src`, il va 
- remplir le fichier go.mod pour y ajouter les dépendances manquantes
- générer le fichier `go.sum` qui est le fichier `lock`

#### C'est quoi un module

Un module, c est juste un package golang qui fait partie d'un projet qui est en mode module. 
Un projet est en mode module si on trouve un fichier `go.mod` à la racine. 

#### Sementic versionning et semantic import versioning

Le fichier `go.mod` contient les les règles qui définissent les versions des dépendances que vous voulez avoir.  

Example :
```
module another_awsome_project

go 1.13

require github.com/cheggaaa/pb/v3 v3.0.3
```

**Le sementic versionning** ("semver" à partir de maintenant) c'est le "v3.0.3". Les règles pour définir une version sont standards et définies. [ici](https://semver.org/)  
Dans le cadre du fichier `go.mod`, l'outil cherchera la version la plus récente qui matches cette règle.  
Exemple :  
Avec `require github.com/cheggaaa/pb/v3 v3.0.2`, go mod prendra la version `v3.0.2`  
Mais en écrivant `require github.com/cheggaaa/pb/v3 v3`, go mod prendra la version `v3.0.3`  
  
   
**Le semantic import versioning** en revanche, c'est une règle que google impose avec `go mod`.  
C'est le "v3" a la fin de `github.com/cheggaaa/pb/v3`.  
Un module golang n'est valide que si le nom du package contient un `vX` qui correspond à la version majeure du semver.   
Ici "v3" match le "3.0.3"  


#### Exercice 2

Reprendre l'exercice précédant. 
Suprimez le projet `github.com/cheggaaa/pb` dans `$GOPATH/src/github.com/cheggaaa/pb`.  
Utilisez `go mod` récupérer les depandances du projet puis exécutez le.