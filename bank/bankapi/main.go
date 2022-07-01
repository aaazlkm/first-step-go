package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/myuser/bankcore"
)

func main() {
	bankController := BankController{
		accounts: map[int64]bankcore.Account{},
	}

	http.HandleFunc("/statement", bankController.statement)
	http.HandleFunc("/deposit", bankController.deposit)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

type BankController struct {
	accounts map[int64]bankcore.Account
}

func (bankController BankController) statement(w http.ResponseWriter, r *http.Request) {
	accountIdData := r.URL.Query().Get("accountId")

	if accountIdData == "" {
		fmt.Fprintf(w, "accountId is empty")
		return
	}

	accountId, err := strconv.ParseInt(accountIdData, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "accountId is not a number")
		return
	}

	account, exit := bankController.accounts[accountId]

	if !exit {
		fmt.Fprintf(w, "accountId is not found")
		return
	}

	fmt.Fprint(w, account.Statement())
}

func (bankController BankController) deposit(w http.ResponseWriter, r *http.Request) {
	accountIdData := r.URL.Query().Get("accountId")
	amountData := r.URL.Query().Get("amount")

	if accountIdData == "" {
		fmt.Fprintf(w, "accountId is empty")
		return
	}

	if amountData == "" {
		fmt.Fprintf(w, "amount is empty")
		return
	}

	accountId, err := strconv.ParseInt(accountIdData, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "accountId is not int")
		return
	}

	amount, err := strconv.ParseFloat(amountData, 64)
	if err != nil {
		fmt.Fprintf(w, "amount is not float")
		return
	}

	account, exit := bankController.accounts[accountId]
	if !exit {
		fmt.Fprintf(w, "account is not found")
	}

	if err := account.Deposit(amount); err != nil {
		fmt.Fprintf(w, "cannot deposit")
		return
	}

	fmt.Fprint(w, account.Statement())
}
