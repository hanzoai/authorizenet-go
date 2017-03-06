package AuthorizeCIM

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var api_endpoint string = "https://apitest.authorize.net/xml/v1/request.api"
var apiName *string
var apiKey *string
var testMode string

func SetAPIInfo(name string, key string, mode string) {
	apiKey = &key
	apiName = &name
	if mode == "test" {
		testMode = "testMode"
		api_endpoint = "https://apitest.authorize.net/xml/v1/request.api"
	} else {
		testMode = "liveMode"
		api_endpoint = "https://api.authorize.net/xml/v1/request.api"
	}
}


func GetAuthentication() (MerchantAuthentication) {
	auth := MerchantAuthentication{
		Name:           apiName,
		TransactionKey: apiKey,
	}
	return auth
}


func SendRequest(input []byte) []byte {
	req, err := http.NewRequest("POST", api_endpoint, bytes.NewBuffer(input))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body
}



