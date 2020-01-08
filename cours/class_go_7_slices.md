## Arrays

Les arrays sont des collections d'objet de taille fixe.
Il est plus difficile à utiliser que une slice et apporte moins de features.
En toute occasion, il vaut mieux utiliser slice.

#### Instantation

Créer un array de string de taille 10

```golang
myArray := [10]string{"abc", "def"}
```


## Slices

Les slices sont des collections, proches des List ou arrayList de C#/Java.  
Elles ne sont pas limitées par leur taille.  

Elles utilisent en fond un array.  
Dans le cas où la taille de cet array n'est pas suffisant, un autre plus grand est créé et remplace l'ancien.  
Toutes les valeurs sont alors copiée de l'ancien vers le nouveau.  

La taille de la slice définie le nombre d'éléments dedans.  
La capacité de la slice definit la taille de l'array sous-jacent, et donc la limite d'élément que peut contenir la slice
avant de créer un nouveau array.

Il n'y a pas de génériques en golang. Vous devez spécifier un type au moment de la déclaration

#### Déclaration / Instanciation

Créer une slice et définir la capacité/taille en fonction des éléments fournis:
```golang
mySlice := []string{"abc", "def"}
```

Créer une slice avec une taille de 10:
(Ce n'est probablement pas ce que vous voulez faire)
```golang
mySlice := make([]string,10)
```

Créeer une slice avec une capacité de 10 et une taille de 0:
```golang
mySlice := make([]string,10, 0)
```

#### len / cap
Obtenir le nombre d'éléments dans une slice:
```golang
len(mySlice)
```

Obtenir la capacité de l'array sous-jacent:
```golang
cap(mySlice)
```


#### Ajouter un élément
```golang
mySlice := []string{"abc", "def"}
mySlice = append(mySlice2, "ghi")
```

#### modifier un élément


#### Récupérer un élément
L'index commence à 0
```golang
mySlice := []string{"abc", "def"}
fmt.PrintLn(mySlice[1])
```


#### Slicer une slice
On peut découper une slice en une slice plus petite.

```golang
mySlice := []string{"abc", "def"}
mySlice2 := mySlice[1:2]
fmt.Println(mySlice, mySlice2)
//output: 	[abc def] [def]
```

L'opération est rapide mais les valeurs ne sont pas copiées: modifier la nouvelle slice modifie l'ancienne.  
Demo WTF:
```golang
	mySlice := []string{"abc", "def"}

	mySlice2 := mySlice[1:2]
	fmt.Println(cap(mySlice2))

	mySlice2 = append(mySlice2, "ghi")
	fmt.Println(cap(mySlice2))

	mySlice2[0] = "something"
	fmt.Println(mySlice, mySlice2)
```


#### Exploser une slice
Pour les fonction qui acceptent des parametres variadics  
```golang
mySlice := []interface{}{"abc", "def"}
fmt.Println(mySlice...)	// <- exploding slice 
```


#### Copier une slice
```golang
	mySlice := []string{"abc", "def"}
	myOtherSlice := make([]string, cap(mySlice), len(mySlice))
	copy(myOtherSlice,mySlice)

	fmt.Println(myOtherSlice)
	// Output: [abc def]
```

#### les strings sont des slices deguisé
```golang
	myString := "hello world"
	fmt.Println(myString[2:8])
	// Output: llo wo
```

## Exercice 1
Unir 2 slices

## Exercice 2 
Retirer un element d'une slice
