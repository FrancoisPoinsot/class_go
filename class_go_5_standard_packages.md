## Documentation officielle

La liste complète des packages standard se trouve ici: https://golang.org/pkg/  
Pour commencer, on peut en ignorer la plupart.  

## Les plus importants pour commencer

#### fmt

Format des string et affiche dans la console

#### math

Vous y trouverez des opérations mathématiques et des RNG

#### database

Contient des interfaces pour utiliser des bases de donnée.  
Vous devez quand même importer le driver correspondant à la bases de donnée que vous ciblez.  
ex: `import _ "github.com/denisenkom/go-mssqldb"`  

#### encoding/json

Contient les serialiser/deserialiser vers le JSON.

#### net/http 

Contient des outils pour créer des serveurs ou client http.

#### net/url

Pour créer ou manipuler des url.

#### os 

Interface vers l'OS. 
Utile pour lire des fichiers, écouter des signaux, exécuter des commandes...

#### sort

Pour trier des collections

#### strings 

Quelques opérations sur les strings.
ex: split, trim, contains

#### strconv

Permet de convertir les strings vers/depuis un autre type.

#### time

Contient le type `Time` et les opérations associées




## Exercice 1

Afficher l'heure actuelle dans la console.

## Exercice 2

Afficher le timestamp unix actuel dans la console.