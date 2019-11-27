## Installer golang

- installer chocolatey

https://chocolatey.org/docs/installation

- re-ouvrir powershell (as admin)

- executer `choco install golang`
- re-ouvrir powershell

executer `go version` pour verifier  


- setup le dossier golang
executer `mkdir $Env:GOPATH; mkdir $Env:GOPATH/src; mkdir $Env:GOPATH/bin`   

Le dossier golang se situe est placé à la valeur de la variable d'environment `GOPATH`.  
Dans le cadre de ce tutorial ça devrait être `~/go`.  
Il contient 3 folders:  
	- `src` : contient les source de tout ce que vous écrirez
	- `bin` : contient les binaries de tout ce que vous compilerez
	- 'pkg' : contient des fichiers intermediaire de compilation

C'est pour cette raison que ce setup a mis le dossier `$GOPATH/bin` dans la variable d'environnement `$PATH`.  
Ainsi tout les binaries compilé en go sont directement accessible depuis la ligne de commande. 
