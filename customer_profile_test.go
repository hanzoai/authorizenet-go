package authorizenet

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var newCustomerProfileId string
var newCustomerPaymentId string
var newCustomerShippingId string
var newSecondCustomerProfileId string
var client Client

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSetAPIInfo(t *testing.T) {
	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	//apiMode := os.Getenv("mode")
	client = *New(apiName, apiKey, true)
	t.Log("API Info Set")
}

func TestIsConnected(t *testing.T) {
	authenticated, err := client.IsConnected()
	if err != nil {
		t.Fail()
	}
	if !authenticated {
		t.Fail()
	}
}

func TestCreateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: RandomNumber(1000, 9999),
		Email:              "info@" + RandomString(8) + ".com",
		PaymentProfiles: &PaymentProfiles{
			CustomerType: "individual",
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "10/26",
					//CardCode: "384",
				},
			},
		},
	}

	res, err := customer.CreateProfile(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		newCustomerProfileId = res.CustomerProfileID
		t.Log("New Customer Profile Created #", res.CustomerProfileID)
	} else {
		t.Fail()
		t.Log(res.ErrorMessage())
	}

}

func TestGetProfileIds(t *testing.T) {
	profiles, _ := client.GetProfileIds()

	for _, p := range profiles {
		t.Log("Profile ID #", p)
	}

	if len(profiles) == 0 {
		t.Fail()
	}

	t.Log(profiles)
}

func TestUpdateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: newCustomerProfileId,
		CustomerProfileId:  newCustomerProfileId,
		Description:        "Updated Account",
		Email:              "info@updatedemail.com",
	}

	res, err := customer.UpdateProfile(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		t.Log("Customer Profile was Updated")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestCreateCustomerPaymentProfile(t *testing.T) {

	paymentProfile := CustomerPaymentProfile{
		CustomerProfileID: newCustomerProfileId,
		PaymentProfile: PaymentProfile{
			BillTo: &BillTo{
				FirstName:   "okokk",
				LastName:    "okok",
				Address:     "1111 white ct",
				City:        "los angeles",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
			Payment: &Payment{
				CreditCard: CreditCard{
					CardNumber:     "5424000000000015",
					ExpirationDate: "04/22",
				},
			},
			DefaultPaymentProfile: "true",
		},
	}

	res, err := paymentProfile.Add(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		newCustomerPaymentId = res.CustomerPaymentProfileID
		t.Log("Created new Payment Profile #", res.CustomerPaymentProfileID, "for Customer ID: ", res.CustomerProfileId)
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestGetCustomerPaymentProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	res, err := customer.Info(client)
	if err != nil {
		t.Fail()
	}

	paymentProfiles := res.PaymentProfiles()

	t.Log("Customer Payment Profiles", paymentProfiles)

	if len(paymentProfiles) == 0 {
		t.Fail()
	}

}

func TestGetCustomerPaymentProfileList(t *testing.T) {

	profileIds, err := client.GetPaymentProfileIds("2020-03", "cardsExpiringInMonth")
	if err != nil {
		t.Fail()
	}

	t.Log(profileIds)
}

func TestValidateCustomerPaymentProfile(t *testing.T) {

	customerProfile := Customer{
		ID:        newCustomerProfileId,
		PaymentID: newCustomerPaymentId,
	}

	res, err := customerProfile.Validate(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		t.Log("Customer Payment Profile is VALID")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestUpdateCustomerPaymentProfile(t *testing.T) {

	customer := Profile{
		CustomerProfileId: newCustomerProfileId,
		PaymentProfileId:  newCustomerPaymentId,
		Description:       "Updated Account",
		Email:             "info@" + RandomString(8) + ".com",
		PaymentProfiles: &PaymentProfiles{
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "01/26",
				},
			},
			BillTo: &BillTo{
				FirstName:   "newname",
				LastName:    "golang",
				Address:     "2841 purple ct",
				City:        "los angeles",
				State:       "CA",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
		},
	}

	res, err := customer.UpdatePaymentProfile(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		t.Log("Customer Payment Profile was Updated")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestCreateCustomerShippingProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: "86437",
		CustomerProfileId:  newCustomerProfileId,
		Email:              "info@" + RandomString(8) + ".com",
		Shipping: &Address{
			FirstName:   "My",
			LastName:    "Name",
			Company:     "none",
			Address:     "1111 yellow ave.",
			City:        "Los Angeles",
			State:       "CA",
			Zip:         "92039",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}

	res, err := customer.CreateShipping(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		newCustomerShippingId = res.CustomerAddressID
		t.Log("New Shipping Added: #", res.CustomerAddressID)
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}
}

func TestGetCustomerShippingProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	res, err := customer.Info(client)
	if err != nil {
		t.Fail()
	}

	shippingProfiles := res.ShippingProfiles()

	t.Log("Customer Shipping Profiles", shippingProfiles)

	if shippingProfiles[0].Zip != "92039" {
		t.Fail()
	}

}

func TestUpdateCustomerShippingProfile(t *testing.T) {

	customer := Profile{
		CustomerProfileId: newCustomerProfileId,
		CustomerAddressId: newCustomerShippingId,
		Shipping: &Address{
			FirstName:   "My",
			LastName:    "Name",
			Company:     "none",
			Address:     "1111 yellow ave.",
			City:        "Los Angeles",
			State:       "CA",
			Zip:         "92039",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}

	res, err := customer.UpdateShippingProfile(client)
	if err != nil {
		t.Fail()
	}

	if res.Ok() {
		t.Log("Shipping Address Profile was updated")
	} else {
		t.Log(res.ErrorMessage())
		t.Fail()
	}

}

func TestAcceptProfilePage(t *testing.T) {

}

func TestCreateCustomerProfileFromTransaction(t *testing.T) {

}

func TestCreateSubscriptionCustomerProfile(t *testing.T) {

	amount := RandomNumber(5, 99) + "." + RandomNumber(10, 99)

	subscription := Subscription{
		Name:   "New Customer Profile Subscription",
		Amount: amount,
		//TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			//TrialOccurrences: "0",
			Interval: IntervalMonthly(),
		},
		Profile: &CustomerProfiler{
			CustomerProfileID:         newCustomerProfileId,
			CustomerPaymentProfileID:  newCustomerPaymentId,
			CustomerShippingProfileID: newCustomerShippingId,
		},
	}

	res, err := subscription.Charge(client)
	if err != nil {
		t.Fail()
	}

	if res.Approved() {
		newSubscriptionId = res.SubscriptionID
		t.Log("Customer #", res.CustomerProfileId(), " Created a New Subscription: ", res.SubscriptionID)
	} else {
		t.Log(res.ErrorMessage(), "\n")
		t.Fail()
	}

}

func TestGetCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	res, err := customer.Info(client)
	if err != nil {
		t.Fail()
	}

	paymentProfiles := res.PaymentProfiles()
	shippingProfiles := res.ShippingProfiles()
	subscriptions := res.Subscriptions()

	t.Log("Customer Profile", res)

	t.Log("Customer Payment Profiles", paymentProfiles)
	t.Log("Customer Shipping Profiles", shippingProfiles)
	t.Log("Customer Subscription IDs", subscriptions)

}
