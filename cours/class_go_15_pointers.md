## Les pointeurs

Les pointeurs de Golang ressemblent à ceux du C++ en un peu plus simple.  
Ils permettent de manipuler des adresses memoire.  
https://tour.golang.org/moretypes/1 

## nil

En golang la référence `null` s'écrit `nil`

## initialisation

Cette écriture initialise un pointeur nil de type `*int`
```golang
	var i *int = nil
```

Cette écriture initialise un pointeur vers un int de valeur 0
```golang
	var i *int = new(int)
	j := new(int)
```

## initialiser un pointer et sa valeur référencé en même temps

#### Dans le cas d'une struct

```golang
package main

import (
"fmt"
)

type User struct{
Name string
}

func main() {
	bob := &User{
		Name:"bob",
	}
	fmt.Println(bob)
}
```

#### Dans le cas d'un type primitif

On ne peut pas le faire en une ligne. (sauf si vous voulez la valeur par défaut)

## Tester les pointeurs

Pour tester si le pointer est nil:
```golang
if myPointer == nil {
// do something
}
```

## Note

En golang, les strings sont des types primitifs immuable, pas des pointeurs.

## Exercice 1
Implémenter la fonction `IsNumber` pour avoir le comportement attendu

```golang
package main

import (
"fmt"
)

func main() {
	var isNumber bool
	input := new(string)

	*input = "5"
	isNumber = IsNumber(input)
	fmt.Println(isNumber) // expect true

	*input = "789645"
	isNumber = IsNumber(input)
	fmt.Println(isNumber) // expect true

	*input = "coucou"
	isNumber = IsNumber(input)
	fmt.Println(isNumber) // expect false

	input = nil
	isNumber = IsNumber(input)
	fmt.Println(isNumber) // expect false
}

func IsNumber(*string) bool {
// TODO: implement me
}
```