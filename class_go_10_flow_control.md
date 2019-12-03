#### if, else
Ça existe

### Switch

#### Switch classique
La syntaxe classique:
```golang
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
```


#### Switch sans condition
```golang
	myVar := 10
	switch {
	case myVar > 50:
		fmt.Printf("%d > 50\n", myVar)
	case myVar > 8:
		fmt.Printf("%d > 50\n", myVar)
	case myVar > 2:
		fmt.Printf("%d > 50\n", myVar)
	default:
		fmt.Printf("%d exist!\n", myVar)
	}
```

#### Switch avec fallthrough
Attention en golang le `switch` utilise `break` par défaut.  
Si vous voulez `fallthrough` il faut le préciser.  

```golang
	myVar := 10
	switch {
	case myVar > 50:
		fmt.Printf("%d > 50\n", myVar)
		fallthrough
	case myVar > 8:
		fmt.Printf("%d > 8\n", myVar)
		fallthrough
	case myVar > 2:
		fmt.Printf("%d > 2\n", myVar)
		fallthrough
	default:
		fmt.Printf("%d exist!\n", myVar)
	}
```

### Defer
`defer` assure qu'une fonction est appelée à la fin de la fonction actuelle.  
```golang
	defer fmt.Println("world")
	fmt.Println("hello")
	// Output: world hello
```

> Note: `defer` est utilisé pour gérer les `try`, `catch`, `finally`
pas d'exemple ici, on voit ça plus tard


## Exercice 1: fizz buzz!
Faire une loop de 1 à 100.
Pour tout les nombre multiple de 3: afficher fizz
Pour tout les nombre multiple de 5: afficher buzz
Pour tout les nombre multiple de 3 et 5: afficher fizz buzz
Pour tout les autres, afficher le chiffre

