## Maps

Les maps sont un type de collection indexée.  
Elles sont similaires aux `Dictionnary` du C# ou `Map` du Java.

L'ordre est d'accès par une boucle `for` volontairement aléatoire!

Les maps n'ont pas de `cap`, juste une `len`.

#### Déclaration / Instanciation
```golang
myMap := map[string]string{"myKey":"myValue", "notherKey":"anotherValue"}
```

#### Modifier / ajouter une valeur
```golang
myMap["myKey"] := "myValue" 
```

#### Acceder à une valeur
```golang
myVar := myMap["myKey"] 
```

#### Tester l'existence d'un element
```golang
myVar, ok := myMap["myKey"] 
if ok {
	// ...
}
```

#### Length
```golang
myMap := map[string]string{"myKey":"myValue", "notherKey":"anotherValue"}
myVar := len(myMap)
```

#### Retirer un élément
```golang
mySlice := map[string]string{"myKey": "def"}
delete(mySlice, "myKey")
```

#### Les types
La clé et la valeur peuvent être de n'importe quel type

```golang
myMap := map[int]float32{123: 3.14}
```

```golang
type User struct {
	name string
	age  int
}

myUser := User{
	name: "bob",
	age:  20,
}

myMap := map[string]User{myUser.name: myUser}
fmt.Println(myMap)
```

## Exercice 1
Créer un set d'`User` dont l'unicité est garantie, sans utiliser d'id



