## Exercice : Bank account

Vous avez 2 fichiers de donnée :
- `accounts.json` contient une liste de comptes avec un id et un montant. 
Le montant est noté en centimes. Il correspond au montant disponible au début de la journée.

- `operations.json` contient la liste des operations financiaires entre les comptes depuis le debut de la journée.   
`account_out` est le compte d'origine de la transaction.
`account_in` est le compte destinatire de la transaction.
`amount` contient le montant de la transaction noté en centime. 


1) Créez un nouveau fichier `new_accounts.json` qui contient les nouveaux états pour l'ensemble des comptes du fichier `account.json`

2) L'équipe de Data Inteligence veut le top 10 des comptes ayant le plus d'argent à la fin de la journée. Mais ils utilisent un format different.  
Trouvez les 10 comptes ayant le plus gros montant et exporter dans un nouveau fichier au format YAML `top10.yaml`

example:
```yaml
- account_id : "ACT_5Ryt6TGgzkzORuYGqBd9"
  amount: 10984
- account_id : "ACT_M3vDVinSfFtCuK2wjO8h"
  amount: 9884
- account_id : "ACT_M3vDVinSfFtCuK2wjO8h"
  amount: 9698
  ...
```