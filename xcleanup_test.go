package authorizenet

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestCancelSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo, err := sub.Cancel(client)
	if err != nil {
		t.Fail()
	}

	if subscriptionInfo.Ok() {
		t.Log("Subscription ID has been canceled: ", sub.Id, "\n")
		t.Log(subscriptionInfo.ErrorMessage(), "\n")
	} else {
		t.Log(subscriptionInfo.ErrorMessage())
		t.Fail()
	}

}

func TestCancelSecondSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSecondSubscriptionId,
	}

	subscriptionInfo, err := sub.Cancel(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if subscriptionInfo.Ok() {
		t.Log("Second Subscription ID has been canceled: ", sub.Id, "\n")
		t.Log(subscriptionInfo.ErrorMessage(), "\n")
	} else {
		t.Log(subscriptionInfo.ErrorMessage())
		t.Fail()
	}

}

func TestDeleteCustomerShippingProfile(t *testing.T) {
	customer := Customer{
		ID:         newCustomerProfileId,
		ShippingID: newCustomerShippingId,
	}

	res, err := customer.DeleteShippingProfile(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res.Ok() {
		t.Log("Shipping Profile was Deleted")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}
}

func TestDeleteCustomerPaymentProfile(t *testing.T) {
	customer := Customer{
		ID:        newCustomerProfileId,
		PaymentID: newCustomerPaymentId,
	}

	res, err := customer.DeletePaymentProfile(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res.Ok() {
		t.Log("Payment Profile was Deleted")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}
}

func TestDeleteCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	res, err := customer.DeleteProfile(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res.Ok() {
		t.Log("Customer was Deleted")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestDeleteSecondCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newSecondCustomerProfileId,
	}

	res, err := customer.DeleteProfile(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res.Ok() {
		t.Log("Second Customer was Deleted")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestDeclineTransaction(t *testing.T) {
	oldTransaction := PreviousTransaction{
		//Amount: "49.99",
		RefId: heldTransactionId,
	}

	res, err := oldTransaction.Decline(client)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res.Approved() {
		t.Log("DECLINED the previous transasction that was on Hold. ID #", oldTransaction.RefId)
		t.Log(res.TransactionID())
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}
}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomNumber(min, max int) string {
	num := rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
