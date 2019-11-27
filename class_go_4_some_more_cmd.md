## Quelques commandes

#### Compiler l'application 
`go build main.go`  

Ceci créer un exécutable nommé `main.exe`. Vous pouvez l'exécuter avec `./main.exe` 

#### Compiler vers un autre OS
 `$Env:GOOS = 'linux'; go build main.go`  

Le changement de la variable d'environnement `GOOS` n'est pas persisté.

#### tester le projet
 `go test ./...`  

#### récupérer d'autre projets
 `go get github.com/golang/example/hello`  

Le package sera disponible dans `$Env:GOPATH/src/github.com/golang/example/hello` 

#### Obtenir de l'aide
`go -h`  
`go build -h`  
`go test -h`
  
#### Compiler et placer le binaire dans le dossier `bin`

 `go install`  

#### Vérifier code
 `go vet`  

#### Formater code
 `go fmt`

#### Lister les package du projet courant 
`go list`

#### Note

La plupart de ces commandes accepte l'argument `./...` pour appliquer leur effet à tout les sous-dossier. 


## Exercice 1

Lister les packages du projet `github.com/denisenkom/go-mssqldb`  

## Exercice 2

Récupérer et exécuter `github.com/farhaanbukhsh/chuck`