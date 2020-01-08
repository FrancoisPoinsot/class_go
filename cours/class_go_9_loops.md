## Loop for

Il n'a qu'une loop en golang: `for`

#### for ever
```golang
for {
	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}
```

#### for 0...N
```golang
for i := 0; i < 10 ; i++{
	fmt.Println(i)
}
```

#### for a while
```golang
sum := 100
for sum < 1000 {
	sum += sum
}
fmt.Println(sum)
```


#### for each
```golang
mySlice := []string{"abc", "def"}
for index, valeur := range mySlice {
	fmt.Println(index, valeur)
}
```

#### continue
```golang
for i := 0; i < 4 ; i++{
	if i == 2 {
		continue
	}
	fmt.Println(i)
}
// Output: 0 1 3
```

#### break
```golang
for i := 0; i < 4 ; i++{
	if i == 2 {
		break
	}
	fmt.Println(i)
}
// Output: 0 1
```

## Exercice 1
Trouver pourquoi cette example ne marche pas:
```golang
package main

import (
	"fmt"
)

func main() {
	userAges := map[string]int{
		"user1": 23,
		"user2": 54,
		"user3": 90,
		"user4": 23,
		"user5": 10,
	}
	fmt.Println(joinAllKeys(userAges))
	// Expected : user1 user2 user3 user4 user5
}

func joinAllKeys(myMap map[string]int) string {
	result := ""
	for key := range myMap {
		result += key + " "
	}

	return result
}
```


## Exercice 2
Faire la somme des ages de ces users

## Exercice 3
Faire une fonction qui accepte la map d'user précédente et la transforme en slice de User
```
type User struct{
	name string
	age int
}
```