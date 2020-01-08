## Les tests

Le framework pour exécuter et écrire des test fait partie intégrante du standard go. 
Les tests sont contenus dans les fichier `*_test.go` au sein même du projet. 
Les fonctions de test ont une signature similaire a cette fonction : `func TestMyFunc(t *testing.T)`. Le nom DOIT commencer par "test". 

On exécute les tests avec la commande `go test ./...`.

## exemple

Prenons cette fonction définie dans `main.go`
```golang
	func Add(a,b int) int {
		return a+b
	}
```

On la test avec cette function définie dans `main_test.go`
```golang
	package main

	import "testing"

	func TestAdd(t *testing.T) {
		expectedValue := 6
		result := Add(2, 4)
		if result != expectedValue {
			t.Fail()
			t.Log("expected value is 6")
		}
	}
```

## tips

Il existe plusieurs librairies pour simplifier l'écriture de test. 
Voici comment écrire le même test avec `testify` :
```golang
	package main

	import (
		"github.com/stretchr/testify/assert"
		"testing"
	)

	func TestAdd(t *testing.T) {
		result := Add(2, 4)
		assert.Equal(t, result, 6, "expected value is 6")
	}
```

## Exercice 1
Reprendre l'exercice du [fizz buzz](class_go_10_flow_control.md#exercice-1-fizz-buzz).
Ecrire des tests pour la fonction `func fizzbuzz(i int) string` qui retourne la chaine de charactère à afficher dans la console pour chaque iteration `i`.

```golang
	package main

	import (
		"fmt"
		"strconv"
	)

	func main() {
		for i := 0; i < 100; i++ {
			fmt.Println(fizzbuzz(i))
		}
	}

	func fizzbuzz(i int) string {
		if i%3 == 0 && i%5 == 0 {
			return "fizzbuzz"
		} else if i%3 == 0 {
			return "fizz"
		} else if i%5 == 0 {
			return "buzz"
		}
		return strconv.Itoa(i)
	}
```