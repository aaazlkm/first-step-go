package bankcore

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithDraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)
	if account.Balance != 0 {
		t.Error("balance is not being updated after a withdraw")
	}
}

func TestWithDrawInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	if err := account.Withdraw(20); err == nil {
		t.Error("only positive numbers should be allowed to withdraw")
	}
}

func TestStatements(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "test",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	statement := account.Statement()
	println(statement, "\n")
	if statement != "1001 - test - 10" {
		t.Error("statement is not correct")
	}
}
