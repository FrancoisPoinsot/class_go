package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_computeNewAccounts(t *testing.T) {
	accounts := []Account{
		{Amount: 10, ID: "1"},
		{Amount: 10, ID: "2"},
		{Amount: 10, ID: "3"},
	}
	operations := []Operation{
		{Amount: 2, AccountIn: "1", AccountOut: "2"},
		{Amount: 3, AccountIn: "3", AccountOut: "1"},
	}
	newAccounts := computeNewAccounts(accounts, operations)

	expected := []Account{
		{Amount: 9, ID: "1"},
		{Amount: 8, ID: "2"},
		{Amount: 13, ID: "3"},
	}
	assert.Equal(t, expected, newAccounts)
}
