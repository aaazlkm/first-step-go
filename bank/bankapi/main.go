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
	http.HandleFunc("/withdraw", bankController.withdraw)
	http.HandleFunc("/send", bankController.send)

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
		return
	}

	if err := account.Deposit(amount); err != nil {
		fmt.Fprintf(w, "cannot deposit")
		return
	}

	fmt.Fprint(w, account.Statement())
}

func (bankController BankController) withdraw(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	if err := account.Withdraw(amount); err != nil {
		fmt.Fprintf(w, "cannot deposit")
		return
	}

	fmt.Fprintf(w, account.Statement())
}

func (bankController BankController) send(w http.ResponseWriter, r *http.Request) {
	accountIdData := r.URL.Query().Get("accountId")
	amountData := r.URL.Query().Get("amount")
	toAccountId := r.URL.Query().Get("toAccountId")

	accountId, err := parseAccountId(accountIdData)
	if err != nil {
		fmt.Fprintf(w, "accountId cannnot be parsed")
		return
	}

	amount, err := parseAmount(amountData)
	if err != nil {
		fmt.Fprintf(w, "amount cannnot be parsed")
		return
	}

	toAccountId, err := parseAccountId(toAccountId)
}

func parseAccountId(accountIdData string) (int64, error) {
	if accountIdData == "" {
		return -1, fmt.Errorf("accountId is empty")
	}

	accountId, err := strconv.ParseInt(accountIdData, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("accountId is empty")
	}

	return accountId, nil
}

func parseAmount(amountData string) (float64, error) {
	if amountData == "" {
		return -1, fmt.Errorf("amount is empty")
	}

	amount, err := strconv.ParseFloat(amountData, 64)
	if err != nil {
		return -1, fmt.Errorf("amount cannnot be parsed")
	}

	return amount, nil
}
