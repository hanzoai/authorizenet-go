![alt tag](http://pichoster.net/images/2016/05/30/authcimyjIpi.jpg)

# Authorize.net CIM for golang
[![Build Status](https://travis-ci.org/Hunterlong/AuthorizeCIM.svg?branch=master)](https://travis-ci.org/Hunterlong/AuthorizeCIM)  [![Code Climate](https://codeclimate.com/github/Hunterlong/AuthorizeCIM/badges/gpa.svg)](https://codeclimate.com/github/Hunterlong/AuthorizeCIM)  [![Coverage Status](https://coveralls.io/repos/github/Hunterlong/AuthorizeCIM/badge.svg?branch=master)](https://coveralls.io/github/Hunterlong/AuthorizeCIM?branch=master)  [![GoDoc](https://godoc.org/github.com/Hunterlong/AuthorizeCIM?status.svg)](https://godoc.org/github.com/Hunterlong/AuthorizeCIM)

Give your Go Language applications the ability to store and retrieve credit cards from Authorize.net.


## Usage

* Import package
```
go get github.com/hunterlong/authorizecim
```
```go
import "github.com/hunterlong/authorizecim"
```

* Set Authorize.net API Keys
```go
apiName = os.Getenv("apiName")
apiKey = os.Getenv("apiKey")
AuthorizeCIM.SetAPIInfo(api_name,api_key)
```

## Features
* Store Billing Profile (credit card) on Authorize.net using CIM
* Fetch customers billing profiles in a simple array
* Process transactions using customers billing profile
* Delete and Updating billing profiles

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

## Examples
Below you'll find the useful functions to get you up and running in no time!


#### Test Correct API Key
```go
connected := AuthorizeCIM.TestConnection()
// true or false
```


#### Create new Customer Account
```go
customer_info := AuthorizeCIM.AuthUser{
                  "54",
                  "email@domain.com",
                  "Test Account"
                  }
new_customer, err := AuthorizeCIM.CreateCustomerProfile(customer_info)
// outputs new user profile ID
```

#### Create Payment Profile for Customer
```go
address := AuthorizeCIM.Address{
                  FirstName: "Test", 
                  LastName: "User", 
                  Address: "1234 Road St", 
                  City: "City Name", 
                  State:" California",
                  Zip: "93063", 
                  Country: "USA", 
                  PhoneNumber: "5555555555"
                  }
credit_card := AuthorizeCIM.CreditCard{
                  CardNumber: "4111111111111111", 
                  ExpirationDate: "2020-12"
                  }
profile_id := "53"
AuthorizeCIM.CreateCustomerBillingProfile(profile_id, address, credit_card)
```
#### Get Customers stored billing accounts
```go
profile_id := "30089822"
billing_accounts := AuthorizeCIM.GetCustomerProfile(profile_id)
```

#### Delete Customer Profile
```go
profile_id := "30089822"
billing_accounts = AuthorizeCIM.DeleteCustomerProfile(profile_id)
```

#### Get detailed information about a billing profile from customer
```go
profile_id := "30089822"
payment_id := "1200079812"
stored_card := AuthorizeCIM.GetCustomerPaymentProfile(profile_id,payment_id)
```

#### Delete customers billing profile
```go
profile_id := "30089822"
payment_id := "1200079812"
AuthorizeCIM.DeleteCustomerPaymentProfile(profile_id,payment_id)
```

#### Update a single billing profile with new information
```go
new_address := AuthorizeCIM.Address{
                  FirstName: "Test", 
                  LastName: "User", 
                  Address: "1234 Road St", 
                  City: "City Name", 
                  State:" California",
                  Zip: "93063", 
                  Country: "USA", 
                  PhoneNumber: "5555555555"
                  }
credit_card := AuthorizeCIM.CreditCard{
                  CardNumber: "4111111111111111", 
                  ExpirationDate: "2020-12"
                  }
profile_id := "53"
payment_id := "416"
AuthorizeCIM.UpdateCustomerPaymentProfile(profile_id,payment_id,new_address,credit_card)
```

#### Create a transaction that will be charged on customers billing profile
```go
item := LineItem{
                  ItemID: "S0897", 
                  Name: "New Product", 
                  Description: "brand new", 
                  Quantity: "1", 
                  UnitPrice: "14.43"
                  }
amount := "14.43"
profile_id := "53"
payment_id := "416"

tranx, _ := CreateTransaction(profile_id, payment_id, item, amount)

var tranx_id string
fixtransx, _ := tranx["transactionResponse"].(map[string]interface{})
if tranx["transactionResponse"]==nil {
    // this transaction failed!
} else {
	tranx_id = fixtransx["transId"].(string)
	fmt.Println("Received Transaction ID: "+tranx_id)
}
```

# Testing
```go
go test -v
```
##### This will run a test of each function, make sure you have correct API keys for Authorize.net

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

# ToDo
* Make cleaner maps for outputs
* Subscriptions (create, status, delete)
* Functions to search Subscriptions (active, expired, etc)
* Get information about Transactions ID
* Void and Refund Transactions
* Add Bank Account Support
* Authorize Only methods

### Authorize.net CIM Documentation
http://developer.authorize.net/api/reference/#customer-profiles

### Authorize.net Sandbox Access
https://developer.authorize.net/hello_world/sandbox/

# License
This golang package is release under MIT license. This software gets reponses in JSON, but Authorize.net currently says "JSON Support is in BETA, please contact us if you intend to use it in production." Make sure you test in sandbox mode!
