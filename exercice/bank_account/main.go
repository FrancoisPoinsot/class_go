package main

import (
	"encoding/json"
	"github.com/thanhpk/randstr"
	"io/ioutil"
	"log"
	"math/rand"
)

type Operation struct {
	AccountIn  string `json:"account_in"`
	AccountOut string `json:"account_out"`
	Amount     int    `json:"amount"`
}

type Account struct {
	ID     string `json:"account_id"`
	Amount int    `json:"amount"`
}

func main() {
	totalOperationsCount := 10000
	totalAccounts := 100

	operations := make([]Operation, 0, totalOperationsCount)
	accounts := make([]Account, 0, totalAccounts)
	accountIndex := make(map[string]bool, cap(accounts))

	// get accounts
	for i := 0; i < totalAccounts; i++ {
		a := Account{
			ID:     "ACT_" + randstr.String(20),
			Amount: rand.Intn(10000),
		}

		// check unique account ID
		if _, exist := accountIndex[a.ID]; exist {
			i--
			continue
		}

		accountIndex[a.ID] = true
		accounts = append(accounts, a)
	}

	// get operations
	for i := 0; i < totalOperationsCount; i++ {
		in := accounts[rand.Intn(len(accounts))]
		out := accounts[rand.Intn(len(accounts))]

		// check in != out
		if in.ID == out.ID {
			i--
			continue
		}

		o := Operation{
			Amount:     rand.Intn(1000) + rand.Intn(10),
			AccountIn:  in.ID,
			AccountOut: out.ID,
		}
		operations = append(operations, o)
	}

	// dump to file
	{
		jsonBytes, err := json.MarshalIndent(operations, "", "	")
		if err != nil {
			log.Panic(err)
		}
		err = ioutil.WriteFile("operations.json", jsonBytes, 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
	{
		jsonBytes, err := json.MarshalIndent(accounts, "", "	")
		if err != nil {
			log.Panic(err)
		}
		err = ioutil.WriteFile("accounts.json", jsonBytes, 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
