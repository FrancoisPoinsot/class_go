## L'interface vide
Une interface qui ne déclare aucune méthode est valide. C'est l'interface vide.  

```golang
	var somthing interface{}
	somthing = 5
	fmt.Println(somthing)
```

En golang il n'y a pas de générique. On utilise des interfaces à la place.  
L'interface vide est implémenté par m'importe quel type.  

La fonction `fmt.Println` en est un exemple. Elle accepte n'importe quoi.

## type assertion
On peut vérifier le type d'un objet pendant l'exécution du programme.  
```golang
func main() {
	var something interface{}
	something = 5

	stringValue, ok := something.(string)
	if ok {
		fmt.Println(stringValue)
	} else {
		fmt.Println("this is not a string")
	}
}
```

## cast != type assertion
Le cast change le type d'un objet, **si possible**, au moment de la **compilation** et s'écrit ainsi
```golang
	var something int64
	something = 5
	somethingElse := int(something)
	fmt.Println(somethingElse)
```

Le type assertion se résout pendant l'exécution, sa réussite n'est pas garantie et il ne change pas le type.
```golang
func main() {
	var something interface{}
	something = 5

	fmt.Printf("the type is: %T\n", something)	// the type is: int

	somethingElse, ok := something.(int)
	if ok {
		fmt.Printf("now the type is: %T\n", somethingElse)	// now the type is: int
	}
}
```

```golang
func main() {
	var something int
	something = 5

	fmt.Printf("the type is: %T\n", something)	// the type is: int

	somethingElse := int64(something)
	fmt.Printf("now the type is: %T\n", somethingElse)	// now the type is: int64
}
```

#### Types switch 
On peut aussi utiliser le switch pour gérer plusieurs cas de type assertion facilement.  

```golang
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

## Exercice 1
Implémenter `numberOfDigit` pour retourner le nombre de chiffres contenu dans le paramètre `number`:  
```golang
package main

import (
	"fmt"
)

func main() {
	var digitsNumber int
	digitsNumber = numberOfDigit(42)
	fmt.Println(digitsNumber) // expect 2

	digitsNumber = numberOfDigit("8763")
	fmt.Println(digitsNumber) // expect 4
}

// Returns the number of digits
// You can assume any string given contains only digits
func numberOfDigit(number interface{}) int {
	// TODO
}
```