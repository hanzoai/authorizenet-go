package authorizenet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultHTTPTimeout = 80 * time.Second

type Client struct {
	APIName   string
	APIKey    string
	Endpoint  string
	Client    *http.Client
	Live      bool
	Connected bool
	Verbose   bool
}

func New(apiName string, apiKey string, testMode bool) *Client {
	// Default to production endpoint / livemode
	endpoint := "https://api.authorize.net/xml/v1/request.api"
	mode := "liveMode"

	// Use test endpoints if testMode is true
	if testMode {
		endpoint = "https://apitest.authorize.net/xml/v1/request.api"
		mode = "testMode"
	}

	return &Client{
		APIKey:   apiKey,
		APIName:  apiName,
		Endpoint: endpoint,
		Client:   &http.Client{Timeout: defaultHTTPTimeout},
	}
}

func (c *Client) IsConnected() (bool, error) {
	info, err := GetMerchantDetails()
	if err != nil {
		return false, err
	}
	if info.Ok() {
		return true, err
	}
	return false, err
}

func (c *Client) GetAuthentication() MerchantAuthentication {
	auth := MerchantAuthentication{
		Name:           c.APIName,
		TransactionKey: c.APIKey,
	}
	return auth
}

func (c *Client) SendRequest(input []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(input))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	if c.Verbose {
		fmt.Println(string(body))
	}
	return body, err
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

func (r AVS) Text() string {
	var response string
	switch r.avsResultCode {
	case "E":
		response = "AVS data provided is invalid or AVS is not allowed for the card type that was used."
	case "R":
		response = "The AVS system was unavailable at the time of processing."
	case "G":
		response = "The card issuing bank is of non-U.S. origin and does not support AVS"
	case "U":
		response = "The address information for the cardholder is unavailable."
	case "S":
		response = "The U.S. card issuing bank does not support AVS."
	case "N":
		response = "Address: No Match ZIP Code: No Match"
	case "A":
		response = "Address: Match ZIP Code: No Match"
	case "Z":
		response = "Address: No Match ZIP Code: Match"
	case "W":
		response = "Address: No Match ZIP Code: Matched 9 digits"
	case "X":
		response = "Address: Match ZIP Code: Matched 9 digits"
	case "Y":
		response = "Address: Match ZIP: Matched first 5 digits"
	}
	return response
}
