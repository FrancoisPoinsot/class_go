## Composition
Il n'y a pas d'héritage en golang.  
Tout se fait par composition.  

```golang
type Engine struct {
	Model string
	Power int
}

type Wheel struct {
	size int
}

type Car struct {
	Engine Engine
	Wheel  Wheel
}
```

On peut aussi faire "hériter" le type parent de toute les propriétés des enfants lorsqu'on omet le nom de l'enfant.
```golang
type Animal struct {
	Size string
}

type Pet struct {
	Animal
	Name string
}

func main() {
	myPet := Pet{
		Name: "bibi",
		Animal: Animal{
			Size: "small",
		},
	}

	fmt.Println(myPet.Size)
}
```
