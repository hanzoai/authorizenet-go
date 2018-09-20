package authorizenet

import (
	"testing"
)

var previousAuth string
var previousCharged string
var heldTransactionId string

func TestChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "15.90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
	}
	res, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}
	if res.Approved() {
		previousCharged = res.TransactionID()
		t.Log("#", res.TransactionID(), "Transaction was CHARGED $", newTransaction.Amount, "\n")
		t.Log("AVS Result Code: ", res.AVS().avsResultCode+"\n")
		t.Log("AVS ACVV Result Code: ", res.AVS().cavvResultCode+"\n")
		t.Log("AVS CVV Result Code: ", res.AVS().cvvResultCode+"\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestAVSDeclinedChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".75",
		CreditCard: CreditCard{
			CardNumber:     "5424000000000015",
			ExpirationDate: "08/" + RandomNumber(20, 27),
		},
		BillTo: &BillTo{
			FirstName:   RandomString(7),
			LastName:    RandomString(9),
			Address:     "1111 white ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "46205",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	res, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}

	if res.AVS().avsResultCode == "N" {
		t.Log("#", res.TransactionID(), "AVS Transaction was DECLINED due to AVS Code. $", newTransaction.Amount, "\n")
		t.Log("AVS Result Text: ", res.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", res.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", res.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", res.AVS().cvvResultCode, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
		t.Log(res.Message(), "\n")
		t.Fail()
	}
}

func TestAVSChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".75",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "08/" + RandomNumber(20, 27),
		},
		BillTo: &BillTo{
			FirstName:   RandomString(7),
			LastName:    RandomString(9),
			Address:     "1111 green ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "46203",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	res, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		heldTransactionId = res.TransactionID()
	}

	if res.Held() {
		t.Log("Transaction is being Held for Review", "\n")
	}

	if res.AVS().avsResultCode == "E" {
		t.Log("#", res.TransactionID(), "AVS Transaction was CHARGED is now on HOLD$", newTransaction.Amount, "\n")
		t.Log("AVS Result Text: ", res.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", res.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", res.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", res.AVS().cvvResultCode, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
		t.Log(res.Message(), "\n")
		t.Fail()
	}
}

func TestDeclinedChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
		BillTo: &BillTo{
			FirstName:   "Declined",
			LastName:    "User",
			Address:     "1337 Yolo Ln.",
			City:        "Beverly Hills",
			State:       "CA",
			Country:     "USA",
			Zip:         "46282",
			PhoneNumber: "8885555555",
		},
	}
	res, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		t.Fail()
	} else {
		t.Log("#", res.TransactionID(), "Transaction was DECLINED!!!", "\n")
		t.Log(res.Message(), "\n")
		t.Log("AVS Result Text: ", res.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", res.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", res.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", res.AVS().cvvResultCode, "\n")
	}
}

func TestAuthOnlyCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "100.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/27",
		},
	}
	res, err := newTransaction.AuthOnly()
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		previousAuth = res.TransactionID()
		t.Log("#", res.TransactionID(), "Transaction was AUTHORIZED $", newTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestCaptureAuth(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: "49.99",
		RefId:  previousAuth,
	}
	res, err := oldTransaction.Capture()
	if err != nil {
		t.Fail()
	}
	if res.Approved() {
		t.Log("#", res.TransactionID(), "Transaction was CAPTURED $", oldTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestChargeCardChannel(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "38.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		AuthCode: "RANDOMAUTHCODE",
	}
	res, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		previousAuth = res.TransactionID()
		t.Log("#", res.TransactionID(), "Transaction was Charged Through Channel (AuthCode) $", newTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestRefundCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "15.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		RefTransId: "0392482938402",
	}
	res, err := newTransaction.Refund()
	if err != nil {
		t.Fail()
	}
	if res.Approved() {
		t.Log("#", res.TransactionID(), "Transaction was REFUNDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestVoidCard(t *testing.T) {
	newTransaction := PreviousTransaction{
		RefId: previousCharged,
	}
	res, err := newTransaction.Void()
	if err != nil {
		t.Fail()
	}
	if res.Approved() {
		t.Log("#", res.TransactionID(), "Transaction was VOIDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}

func TestChargeCustomerProfile(t *testing.T) {

	oldProfileId := "1810921101"
	oldPaymentId := "1805617738"

	customer := Customer{
		ID:        oldProfileId,
		PaymentID: oldPaymentId,
	}

	newTransaction := NewTransaction{
		Amount: "35.00",
	}

	res, err := newTransaction.ChargeProfile(customer)
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		t.Log("#", res.TransactionID(), "Customer was Charged $", newTransaction.Amount, "\n")
	} else {
		t.Log(res.ErrorMessage(), "\n")
	}
}
