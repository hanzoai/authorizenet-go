package main

import (
	"fmt"
	"github.com/hanzoai/authorizenet-go"
	"os"
)

var newTransactionId string

func main() {

	client := authorizenet.New(apiName, apiKey, true)

	if client.IsConnected() {
		fmt.Println("Connected to Authorize.net!")
	}

	client.ChargeCustomer()
	client.VoidTransaction()
}

func ChargeCustomer() {

	newTransaction := client.NewTransaction{
		Amount: "13.75",
		CreditCard: authorizenet.CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "08/25",
			CardCode:       "393",
		},
		BillTo: &authorizenet.BillTo{
			FirstName:   "Timmy",
			LastName:    "Jimmy",
			Address:     "1111 green ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "43534",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	res := newTransaction.Charge()

	if res.Approved() {
		newTransactionId = res.TransactionID()
		fmt.Println("Transaction was Approved! #", res.TransactionID())
	}
}

func VoidTransaction() {

	newTransaction := client.PreviousTransaction{
		RefId: newTransactionId,
	}
	res := newTransaction.Void()
	if res.Approved() {
		fmt.Println("Transaction was Voided!")
	}

}
