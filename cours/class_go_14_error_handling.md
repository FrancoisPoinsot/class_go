## gérer les erreurs

En golang, une fonction peut retourner plusieurs paramètres.
On se sert de cette propriété pour gérer les erreurs.
Si jamais il y a une erreur à retourner, la convention veut qu'on le met dans le dernier paramètre de sortie.

```golang
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(42, 0)
	fmt.Println(result, err)


	result, err = Divide(42, 12)
	fmt.Println(result, err)
}

func Divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("b cannot == 0")
	}
	return a / b, nil
}
```

## panic 
La panic se comporte comme un `trow` en C# ou Java.  
Mais contrairement à ces deux langage, **il est fortement déconseillé** de s'appuyer sur la panic pour gérer les cas d'erreur.  
On l'appelle avec
```golang
panic("some explanation")
```

Pour récupérer d'une panic sans interrompre le programme on peut utiliser `recover`
```golang
defer func() {
	if r := recover(); r != nil {
		// If you are here, it means you had a panic
	}
}()
```


## Exercice 1
Ouvrir et lire le contenu d'un fichier.
Mettez le nom du fichier dans une constante.
Testez avec des fichiers qui existe et d'autre qui n'existe pas.
Gérez les cas d'erreur en affichant un message explicatif a la place d'une panic.

pssst: vous voulez utiliser le package `os`