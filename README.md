# Authorize.net CIM, AIM, and ARB for Go Language

[![Build Status](https://travis-ci.org/hanzoai/authorizenet.svg?branch=master)](https://travis-ci.org/hanzoai/authorizenet)  [![Code Climate](https://lima.codeclimate.com/github/hanzoai/authorizenet/badges/gpa.svg)](https://lima.codeclimate.com/github/hanzoai/authorizenet) [![Coverage Status](https://coveralls.io/repos/github/hanzoai/authorizenet/badge.svg?branch=master)](https://coveralls.io/github/hanzoai/authorizenet?branch=master) [![GoDoc](https://godoc.org/github.com/hanzoai/authorizenet?status.svg)](https://godoc.org/github.com/hanzoai/authorizenet) [![Go Report Card](https://goreportcard.com/badge/github.com/hanzoai/authorizenet)](https://goreportcard.com/report/github.com/hanzoai/authorizenet)

Give your Go Language applications the ability to store and retrieve credit cards from Authorize.net CIM, AIM, and ARB API.
This golang package lets you create recurring subscriptions, AUTH only transactions, voids, refunds, and other functionality connected to the Authorize.net API.

***

# Features
* [AIM Payment Transactions](https://github.com/hanzoai/authorizenet#payment-transactions)
* [CIM Customer Information Manager](https://github.com/hanzoai/authorizenet#customer-profile)
* [ARB Automatic Recurring Billing](https://github.com/hanzoai/authorizenet#recurring-billing) (Subscriptions)
* [Transaction Reporting](https://github.com/hanzoai/authorizenet#transaction-reporting)
* [Fraud Management](https://github.com/hanzoai/authorizenet#fraud-management)
* Creating Users Accounts based on user's unique ID and/or email address
* Store Payment Profiles (credit card) on Authorize.net using Customer Information Manager (CIM)
* Create Subscriptions (monthly, weekly, days) with Automated Recurring Billing (ARB)
* Process transactions using customers stored credit card
* Delete and Updating payment profiles
* Add Shipping Profiles into user accounts
* Delete a customers entire account
* Tests included and examples below

```go
customer := authorizenet.Customer{
        ID: "13838",
    }

customerInfo := customer.Info()

paymentProfiles := customerInfo.PaymentProfiles()
shippingProfiles := customerInfo.ShippingProfiles()
subscriptions := customerInfo.Subscriptions()
```
***

# Usage
* Import package
```
go get github.com/hanzoai/authorizenet-go
```
```go
import "github.com/hanzoai/authorizenet-go"
```

## Set HTTP Client
This library allows you to set your own HTTP client for talking to the
Authorize.net API. This is useful for working in Appengine contexts where a
default client may not exist.

Usage:

```
import "context"

ctx := context.TODO()
httpClient := urlfetch.Client(ctx)

httpClient.Transport = &urlfetch.Transport{
    Context: ctx,
    AllowInvalidServerCertificate: appengine.IsDevAppServer(),
}

client := authorizenet.New(apiName, apiKey, tesMode)
client.SetHTTPClient(httpClient)
```

## Set Authorize.net API Keys
You can get Sandbox Access at:  https://developer.authorize.net/hello_world/sandbox/
```go
apiName := "auth_name_here"
apiKey := "auth_transaction_key_here"
authorizenet.SetAPIInfo(apiName,apiKey,"test")
// use "live" to do transactions on production server
```
***

## Included API References

:white_check_mark: Set API Creds
```go
func main() {

    apiName := "PQO38FSL"
    apiKey := "OQ8NFBAPA9DS"
    apiMode := "test"

    authorizenet.SetAPIInfo(apiName,apiKey,apiMode)

}
```
***

# Payment Transactions

:white_check_mark: chargeCard
```go
newTransaction := authorizenet.NewTransaction{
		Amount: "15.90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
	}
response, err := newTransaction.Charge()
if response.Approved() {

}
```
***

:white_check_mark: authorizeCard
```go
newTransaction := authorizenet.NewTransaction{
		Amount: "100.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/27",
		},
	}
response, err := newTransaction.AuthOnly()
if response.Approved() {

}
```
***

:white_check_mark: capturePreviousCard
```go
oldTransaction := authorizenet.PreviousTransaction{
		Amount: "49.99",
		RefId:  "AUTHCODEHERE001",
	}
response, err := oldTransaction.Capture()
if response.Approved() {

}
```
***

:white_check_mark: captureAuthorizedCardChannel
```go
newTransaction := authorizenet.NewTransaction{
		Amount: "38.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		AuthCode: "YOURAUTHCODE",
	}
response, err := newTransaction.Charge()
if response.Approved() {

}
```
***

:white_check_mark: refundTransaction
```go
newTransaction := authorizenet.NewTransaction{
		Amount: "15.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		RefTransId: "0392482938402",
	}
response, err := newTransaction.Refund()
if response.Approved() {

}
```
***

:white_check_mark: voidTransaction
```go
oldTransaction := authorizenet.PreviousTransaction{
		RefId: "3987324293834",
	}
response, err := oldTransaction.Void()
if response.Approved() {

}
```
***

:white_medium_square: updateSplitTenderGround

:white_medium_square: debitBankAccount

:white_medium_square: creditBankAccount

:white_check_mark: chargeCustomerProfile
```go
customer := authorizenet.Customer{
		ID: "49587345",
		PaymentID: "84392124324",
	}

newTransaction := authorizenet.NewTransaction{
		Amount: "35.00",
	}

response, err := newTransaction.ChargeProfile(customer)

if response.Ok() {

}
```
***

:white_medium_square: chargeTokenCard

:white_medium_square: creditAcceptPaymentTransaction

:white_medium_square: getAccessPaymentPage

:white_medium_square: getHostedPaymentPageRequest

## Transaction Responses
```go
response.Ok()                   // bool
response.Approved()             // bool
response.Message()              // string
response.ErrorMessage()         // string
response.TransactionID()        // string
response.AVS()                  // [avsResultCode,cavvResultCode,cvvResultCode]
```
***

# Fraud Management

:white_check_mark: getUnsettledTransactionListRequest
```go
transactions := authorizenet.UnsettledBatchList()
fmt.Println("Unsettled Count: ", transactions.Count)
```
***

:white_check_mark: updateHeldTransactionRequest
```go
oldTransaction := authorizenet.PreviousTransaction{
		Amount: "49.99",
		RefId:  "39824723983",
	}

	response, err := oldTransaction.Approve()
	//response := oldTransaction.Decline()

	if response.Ok() {

	}
```
***

# Recurring Billing

:white_check_mark: ARBCreateSubscriptionRequest
```go
subscription := authorizenet.Subscription{
		Name:        "New Subscription",
		Amount:      "9.00",
		TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: authorizenet.IntervalMonthly(),
		},
		Payment: &Payment{
			CreditCard: CreditCard{
				CardNumber:     "4007000000027",
				ExpirationDate: "10/23",
			},
		},
		BillTo: &BillTo{
			FirstName: "Test",
			LastName:  "User",
		},
	}

response, err := subscription.Charge()

if response.Approved() {
    fmt.Println("New Subscription ID: ",response.SubscriptionID)
}
```
###### For Intervals, you can use simple methods
```go
authorizenet.IntervalWeekly()      // runs every week (7 days)
authorizenet.IntervalMonthly()     // runs every Month
authorizenet.IntervalQuarterly()   // runs every 3 months
authorizenet.IntervalYearly()      // runs every 1 year
authorizenet.IntervalDays("15")    // runs every 15 days
authorizenet.IntervalMonths("6")   // runs every 6 months
```
***

:white_check_mark: ARBCreateSubscriptionRequest from Customer Profile
```go
subscription := authorizenet.Subscription{
		Name:        "New Customer Subscription",
		Amount:      "12.00",
		TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: authorizenet.IntervalDays("15"),
		},
		Profile: &CustomerProfiler{
			CustomerProfileID: "823928379",
			CustomerPaymentProfileID: "183949200",
			//CustomerShippingProfileID: "310282443",
		},
	}

	response, err := subscription.Charge()

	if response.Approved() {
		newSubscriptionId = response.SubscriptionID
		fmt.Println("Customer #",response.CustomerProfileId(), " Created a New Subscription: ", response.SubscriptionID)
	}
```
***

:white_check_mark: ARBGetSubscriptionRequest
```go
sub := authorizenet.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo := sub.Info()
```
***

:white_check_mark: ARBGetSubscriptionStatusRequest
```go
sub := authorizenet.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo, err := sub.Status()

fmt.Println("Subscription ID has status: ",subscriptionInfo.Status)
```
***

:white_check_mark: ARBUpdateSubscriptionRequest
```go
subscription := authorizenet.Subscription{
		Payment: Payment{
			CreditCard: CreditCard{
				CardNumber:     "5424000000000015",
				ExpirationDate: "06/25",
			},
		},
		SubscriptionId: newSubscriptionId,
	}

response, err := subscription.Update()

if response.Ok() {

}
```
***

:white_check_mark: ARBCancelSubscriptionRequest
```go
sub := authorizenet.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo, err := sub.Cancel()

fmt.Println("Subscription ID has been canceled: ", sub.Id, "\n")
```
***

:white_check_mark: ARBGetSubscriptionListRequest
```go
inactive := authorizenet.SubscriptionList("subscriptionInactive")
fmt.Println("Amount of Inactive Subscriptions: ", inactive.Count())

active := authorizenet.SubscriptionList("subscriptionActive")
fmt.Println("Amount of Active Subscriptions: ", active.Count())
```
***

# Customer Profile (CIM)

:white_check_mark: createCustomerProfileRequest
```go
customer := authorizenet.Profile{
		MerchantCustomerID: "86437",
		Email:              "info@emailhereooooo.com",
		PaymentProfiles: &PaymentProfiles{
			CustomerType: "individual",
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "10/23",
				},
			},
		},
	}

	response, err := customer.Create()

if response.Ok() {
    fmt.Println("New Customer Profile Created #",response.CustomerProfileID)
    fmt.Println("New Customer Payment Profile Created #",response.CustomerPaymentProfileID)
} else {
       fmt.Println(response.ErrorMessage())
   }
```
***

:white_check_mark: getCustomerProfileRequest
```go
customer := authorizenet.Customer{
		ID: "13838",
	}

customerInfo, err := customer.Info()

paymentProfiles := customerInfo.PaymentProfiles()
shippingProfiles := customerInfo.ShippingProfiles()
subscriptions := customerInfo.Subscriptions()
```
***

:white_check_mark: getCustomerProfileIdsRequest
```go
profiles, _ := authorizenet.GetProfileIds()
fmt.Println(profiles)
```
***

:white_check_mark: updateCustomerProfileRequest
```go
customer := authorizenet.Profile{
		MerchantCustomerID: "13838",
		CustomerProfileId: "13838",
		Description: "Updated Account",
		Email:       "info@updatedemail.com",
	}

	response := customer.Update()

if response.Ok() {

}
```
***

:white_check_mark: deleteCustomerProfileRequest
```go
customer := authorizenet.Customer{
		ID: "13838",
	}

	response, err := customer.Delete()

if response.Ok() {

}
```
***

# Customer Payment Profile

:white_check_mark: createCustomerPaymentProfileRequest
```go
paymentProfile := authorizenet.CustomerPaymentProfile{
		CustomerProfileID: "32948234232",
		PaymentProfile: PaymentProfile{
			BillTo: BillTo{
				FirstName: "okokk",
				LastName: "okok",
				Address: "1111 white ct",
				City: "los angeles",
				Country: "USA",
				PhoneNumber: "8885555555",
			},
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber: "5424000000000015",
					ExpirationDate: "04/22",
				},
			},
			DefaultPaymentProfile: "true",
		},
	}

response, err := paymentProfile.Add()

if response.Ok() {

} else {
    fmt.Println(response.ErrorMessage())
}
```
***

:white_check_mark: getCustomerPaymentProfileRequest
```go
customer := authorizenet.Customer{
		ID: "3923482487",
	}

response, err := customer.Info()

paymentProfiles := response.PaymentProfiles()
```

:white_check_mark: getCustomerPaymentProfileListRequest
```go
profileIds := authorizenet.GetPaymentProfileIds("2017-03","cardsExpiringInMonth")
```
***

:white_check_mark: validateCustomerPaymentProfileRequest
```go
customerProfile := authorizenet.Customer{
		ID: "127723778",
		PaymentID: "984583934",
	}

response, err := customerProfile.Validate()

if response.Ok() {

}
```
***

:white_check_mark: updateCustomerPaymentProfileRequest
```go
customer := authorizenet.Profile{
		CustomerProfileId:  "3838238293",
		PaymentProfileId: "83929382739",
		Email:              "info@updatedemail.com",
		PaymentProfiles: &PaymentProfiles{
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber: "4007000000027",
					ExpirationDate: "01/26",
				},
			},
			BillTo: &BillTo{
				FirstName:   "newname",
				LastName:    "golang",
				Address:     "2841 purple ct",
				City:        "los angeles",
				State:		  "CA",
				Zip:            "93939",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
		},
	}

response, err := customer.UpdatePaymentProfile()

if response.Ok() {
    fmt.Println("Customer Payment Profile was Updated")
} else {
    fmt.Println(response.ErrorMessage())
}
```
***

:white_check_mark: deleteCustomerPaymentProfileRequest
```go
customer := authorizenet.Customer{
		ID: "3724823472",
		PaymentID: "98238472349",
	}

response, err := customer.DeletePaymentProfile()

if response.Ok() {
    fmt.Println("Payment Profile was Deleted")
} else {
    fmt.Println(response.ErrorMessage())
}
```
***

# Customer Shipping Profile

:white_check_mark: createCustomerShippingAddressRequest
```go
customer := authorizenet.Profile{
		MerchantCustomerID: "86437",
		CustomerProfileId:  "7832642387",
		Email:              "info@emailhereooooo.com",
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

response, err := customer.CreateShipping()

if response.Ok() {
    fmt.Println("New Shipping Added: #",response.CustomerAddressID)
} else {
    fmt.Println(response.ErrorMessage())
}
```
***

:white_check_mark: getCustomerShippingAddressRequest
```go
customer := authorizenet.Customer{
		ID: "3842934233",
	}

response, err := customer.Info()

shippingProfiles := response.ShippingProfiles()

fmt.Println("Customer Shipping Profiles", shippingProfiles)
```
***

:white_check_mark: updateCustomerShippingAddressRequest
```go
customer := authorizenet.Profile{
		CustomerProfileId:  "398432389",
		CustomerAddressId: "848388438",
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

response, err := customer.UpdateShippingProfile()

if response.Ok() {
    fmt.Println("Shipping Profile was updated")
}
```
***

:white_check_mark: deleteCustomerShippingAddressRequest
```go
customer := authorizenet.Customer{
		ID: "128749382",
		ShippingID: "34892734829",
	}

	response, err := customer.DeleteShippingProfile()

	if response.Ok() {
		fmt.Println("Shipping Profile was Deleted")
	} else {
		fmt.Println(response.ErrorMessage())
	}
```
***

:white_medium_square: getHostedProfilePageRequest

:white_medium_square: createCustomerProfileFromTransactionRequest

# Transaction Reporting

:white_check_mark: getSettledBatchListRequest
```go
list := authorizenet.Range{
		Start: LastWeek(),
		End:   Now(),
	}

batches := list.SettledBatch().List()

for _, v := range batches {
    t.Log("Batch ID: ", v.BatchID, "\n")
    t.Log("Payment Method: ", v.PaymentMethod, "\n")
    t.Log("State: ", v.SettlementState, "\n")
}
```
***

:white_check_mark: getUnSettledBatchListRequest
```go
batches := authorizenet.UnSettledBatch().List()

for _, v := range batches {
    t.Log("Status: ",v.TransactionStatus, "\n")
    t.Log("Amount: ",v.Amount, "\n")
    t.Log("Transaction ID: #",v.TransID, "\n")
}

```
***

:white_check_mark: getTransactionListRequest
```go
list := authorizenet.Range{
		BatchId: "6933560",
	}

batches := list.Transactions().List()

for _, v := range batches {
    t.Log("Transaction ID: ", v.TransID, "\n")
    t.Log("Amount: ", v.Amount, "\n")
    t.Log("Account: ", v.AccountNumber, "\n")
}
```
***

:white_check_mark: getTransactionDetails
```go
oldTransaction := authorizenet.PreviousTransaction{
		RefId: "60019493304",
	}
response := oldTransaction.Info()

fmt.PrintLn("Transaction Status: ",response.TransactionStatus,"\n")
```
***

:white_check_mark: getBatchStatistics
```go
list := authorizenet.Range{
		BatchId: "6933560",
	}

batch := list.Statistics()

fmt.PrintLn("Refund Count: ", batch.RefundCount, "\n")
fmt.PrintLn("Charge Count: ", batch.ChargeCount, "\n")
fmt.PrintLn("Void Count: ", batch.VoidCount, "\n")
fmt.PrintLn("Charge Amount: ", batch.ChargeAmount, "\n")
fmt.PrintLn("Refund Amount: ", batch.RefundAmount, "\n")
```
***

:white_check_mark: getMerchantDetails
```go
info := authorizenet.GetMerchantDetails()

fmt.PrintLn("Test Mode: ", info.IsTestMode, "\n")
fmt.PrintLn("Merchant Name: ", info.MerchantName, "\n")
fmt.PrintLn("Gateway ID: ", info.GatewayID, "\n")
```
***

# ToDo
* Organize and refactor some areas
* Add Bank Account Support
* Make tests fail if transactions fail (skipping 'duplicate transaction')

### Authorize.net CIM Documentation
http://developer.authorize.net/api/reference/#customer-profiles

### Authorize.net Sandbox Access
https://developer.authorize.net/hello_world/sandbox/

# License
[MIT](LICENSE). Originally forked from https://github.com/hunterlong/AuthorizeCIM.
