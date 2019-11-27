## Documentation officiel

La liste complète des packages standard ce trouve ici: https://golang.org/pkg/  
Pour commencer, on peut en ignorer la pluspart.  

## Les plus important pour commencer

#### fmt

Format des string et affiche dans la console

#### math

Vous y trouverez des operations mathematiques et des RNG

#### database

Contient des interface pour utiliser des database.  
Vous devez quand même importer le driver correspondant à la database que vous ciblez.  
ex: `import _ "github.com/denisenkom/go-mssqldb"`  

#### encoding/json

Contient les serialiser/deserialiser vers le JSON.

#### net/http 

Contient des outils pour créer des serveurs ou client http.

#### net/url

Pour créer ou manipuler des url.

#### os 

Interface vers l'OS. 
Utile pour lire des fichier, écouter des signaux, éxecuter des commandes...

#### sort

Pour trier des collections

#### strings 

Quelques opérations sur les strings.
ex: split, trim, contains

#### strconv

Permet de convertir les strings vers/depuis un autre type.

#### time

Contient le type `Time` et les operations associées




## Exercice

Afficher l'heure actuelle dans la console.