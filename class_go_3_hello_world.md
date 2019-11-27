## Hello world

- créer un dossier sous `src`:  
run `mkdir $Env:GOPATH/src/myAwsomeProject; cd $Env:GOPATH/src/myAwsomeProject`   

- créer le fichier main.go avec ce contenu:

```golang
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

```

- pour exécuter l'application : `go run main.go`
