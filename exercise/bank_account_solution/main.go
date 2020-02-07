package main

import (
	"encoding/json"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type Operation struct {
	AccountIn  string `json:"account_in"`
	AccountOut string `json:"account_out"`
	Amount     int64  `json:"amount"`
}

type Account struct {
	ID     string `json:"account_id",yaml:"account_id"`
	Amount int64  `json:"amount",yaml:"amount"`
}

func main() {
	accounts, operations := getData()
	newAccounts := computeNewAccounts(accounts, operations)
	//writeDownNewAccount(newAccounts)

	top10 := getTop10(newAccounts)
	writeDowntop10(top10)

	log.Println("all good")
}

func writeDowntop10(top10 []Account){
	top10Bytes ,err := yaml.Marshal(top10)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("top10.yaml", top10Bytes, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func getTop10(accounts []Account) []Account {
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Amount > accounts[j].Amount
	})

	return accounts[:10]
}

func writeDownNewAccount(newAccounts []Account) {

	newAccountsBytes, err := json.Marshal(newAccounts)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("new_accounts.json", newAccountsBytes, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func computeNewAccounts(accounts []Account, operations []Operation) []Account {
	for _, op := range operations {

		var indexIn, indexOut int
		// find account in
		for indexIn = range accounts {
			if accounts[indexIn].ID == op.AccountIn {
				break
			}
		}

		// find account out
		for indexOut = range accounts {
			if accounts[indexOut].ID == op.AccountOut {
				break
			}
		}

		accounts[indexIn].Amount += op.Amount
		accounts[indexOut].Amount -= op.Amount
	}

	return accounts
}

func getData() ([]Account, []Operation) {
	accountsBytes, err := ioutil.ReadFile(`accounts.json`)
	if err != nil {
		log.Fatal(err)
	}

	accounts := []Account{}
	err = json.Unmarshal(accountsBytes, &accounts)
	if err != nil {
		log.Fatal(err)
	}

	operationsBytes, err := ioutil.ReadFile(`operations.json`)
	if err != nil {
		log.Fatal(err)
	}

	operations := []Operation{}
	err = json.Unmarshal(operationsBytes, &operations)
	if err != nil {
		log.Fatal(err)
	}

	return accounts, operations
}
