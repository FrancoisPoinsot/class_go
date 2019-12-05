## Privé/public

En golang on déclare les propriétés privée en commençant leur nom par une minuscule et public par une majuscule.  
Ce n'est pss juste une convention, c'est le langage.  

Contrairement à la plupart des langage orienté objet, les propriétés privée sont visibles depuis tout le package.  

```golang
type someRandomType struct {
	thisPropertyIsPrivate string
	ThisPropertyIsPublic  string
}
```


## Methods

Puisque les field privé sont visible depuis le package entier, les méthodes n'ont pas besoin d'accès particulier aux objets auquel elles sont attachées.  
En golang, les méthodes sont juste du "syntaxic sugar" par dessus des fonctions classique.  

```golang
package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func (d Dog) Bark() {
	fmt.Printf("%s says woof\n", d.name)
}

func Bark(d Dog) {
	fmt.Printf("%s says woof\n", d.name)
}

func main() {
	myPet := Dog{
		name: "bibi",
	}

	myPet.Bark()
	Bark(myPet)
}
```

## Interfaces

Les interfaces sont des types.  
Comme en C#/Java, elles permettent de décrire une liste de fonction qu'un autre type peut implémenter.  

La seule différence en golang, c'est qu'on peut implémenter une interface **implicitement**  

#### Déclaration

```golang
package main

import (
	"fmt"
)

type Barker interface {
	Bark() string
}

type Dog struct{}

func (d Dog) Bark() string {
	return "woof"
}

func main() {
	myPet := Dog{}
	Listen(myPet)
}

func Listen(pet Barker) {

	fmt.Println(pet.Bark())
}
```

#### Note

On peut aussi implémenter explicitement une interface, mais c'est rarement utilisé.

#### fonction qui implémentent une interface

Puisque les méthodes ne sont que du "syntaxic sugar" au dessus d'une fonction, rien n'empêche une fonction d'avoir une méthode.  
Ainsi une fonction peut implémenter une interface.  

En vrai c'est un trick, mieux vaut éviter.  

On trouve un exemple dans le package standard du net/http.  
```golang
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

## Exercice 1

`error` est une interface très utilisée en golang.
Voici sa signature
```golang
type error interface {
	Error() string
}
```

Implémentez cette interface avec votre propre type.
