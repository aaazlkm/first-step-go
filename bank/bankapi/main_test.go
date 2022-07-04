package main

import (
	"net/http/httptest"
	"testing"

	"github.com/myuser/bankcore"
)

func TestStatement(t *testing.T) {
	bankController := testBankController()

	request := httptest.NewRequest("Get", "/statemnet", nil)
	query := request.URL.Query()
	query.Add("accountId", "1001")
	request.URL.RawQuery = query.Encode()
	response := httptest.NewRecorder()

	bankController.statement(response, request)

	if response.Body.String() != "1001 - John - 0" {
		t.Error("statement is not working")
	}
}

func TestDeposit(t *testing.T) {
	bankController := testBankController()

	response := httptest.NewRecorder()
	request := httptest.NewRequest("Get", "/deposit", nil)
	query := request.URL.Query()
	query.Add("accountId", "1001")
	query.Add("amount", "100")
	request.URL.RawQuery = query.Encode()

	bankController.deposit(response, request)

	if response.Body.String() != "1001 - John - 100" {
		t.Error("deposit is not working")
	}
}

func TestWithDraw(t *testing.T) {
	baseController := testBankController()

	response := httptest.NewRecorder()
	request := httptest.NewRequest("Get", "/withdraw", nil)
	query := request.URL.Query()
	query.Add("accountId", "1003")
	query.Add("amount", "100")
	request.URL.RawQuery = query.Encode()

	baseController.withdraw(response, request)

	if response.Body.String() != "1003 - test3 - 0" {
		t.Error("withdraw is not working")
	}
}

func TestSend(t *testing.T) {
	baseController := testBankController()

	response := httptest.NewRecorder()
	request := httptest.NewRequest("Get", "/send", nil)
	query := request.URL.Query()
	query.Add("accountId", "1003")
	query.Add("amount", "100")
	query.Add("toAccountId", "1002")
	request.URL.RawQuery = query.Encode()

	baseController.send(response, request)

	if response.Body.String() != "1003 - test3 - 0" {
		t.Error("send is not working")
	}
}

func testBankController() BankController {
	return BankController{
		map[int64]bankcore.Account{
			1001: bankcore.Account{
				Customer: bankcore.Customer{
					Name:    "John",
					Address: "Los Angeles, California",
					Phone:   "(213) 555 0147",
				},
				Number:  1001,
				Balance: 0,
			},
			1002: bankcore.Account{
				Customer: bankcore.Customer{
					Name:    "test2",
					Address: "Los Angeles, California",
					Phone:   "(213) 555 0147",
				},
				Number:  1002,
				Balance: 0,
			},
			1003: bankcore.Account{
				Customer: bankcore.Customer{
					Name:    "test3",
					Address: "Los Angeles, California",
					Phone:   "(213) 555 0147",
				},
				Number:  1003,
				Balance: 100,
			},
		},
	}
}
