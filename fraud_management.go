package authorizenet

import (
	"encoding/json"
	"time"
)

func UnsettledBatchList(c Client) (*TransactionsList, error) {
	response, err := c.SendGetUnsettled()
	return response, err
}

func (input TransactionsList) List(c Client) ([]BatchTransaction, error) {
	response, err := c.SendGetUnsettled()
	return response.Transactions, err
}

func updateHeldTransaction() {

}

func (input TransactionsList) Count() int {
	return input.TotalNumInResultSet
}

type UnsettledTransactionsRequest struct {
	GetUnsettledTransactionListRequest GetUnsettledTransactionListRequest `json:"getUnsettledTransactionListRequest"`
}

type GetUnsettledTransactionListRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication,omitempty"`
	Status                 string                 `json:"status,omitempty"`
}

type TransactionsList struct {
	Transactions        []BatchTransaction `json:"transactions"`
	TotalNumInResultSet int                `json:"totalNumInResultSet"`
	MessagesResponse
}

type BatchTransaction struct {
	TransID           string    `json:"transId"`
	SubmitTimeUTC     time.Time `json:"submitTimeUTC"`
	SubmitTimeLocal   string    `json:"submitTimeLocal"`
	TransactionStatus string    `json:"transactionStatus"`
	InvoiceNumber     string    `json:"invoiceNumber"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	AccountType       string    `json:"accountType"`
	AccountNumber     string    `json:"accountNumber"`
	SettleAmount      float64   `json:"settleAmount"`
	MarketType        string    `json:"marketType"`
	Product           string    `json:"product"`
	FraudInformation  struct {
		FraudFilterList []string `json:"fraudFilterList"`
		FraudAction     string   `json:"fraudAction"`
	} `json:"fraudInformation"`
}

type UpdateHeldTransactionRequest struct {
	UpdateHeldTransaction UpdateHeldTransaction `json:"updateHeldTransactionRequest"`
}

type UpdateHeldTransaction struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID                  string                 `json:"refId"`
	HeldTransactionRequest HeldTransactionRequest `json:"heldTransactionRequest"`
}

type HeldTransactionRequest struct {
	Action     string `json:"action"`
	RefTransID string `json:"refTransId"`
}

func (c Client) SendTransactionUpdate(tranx PreviousTransaction, method string) (*TransactionResponse, error) {
	action := UpdateHeldTransactionRequest{
		UpdateHeldTransaction: UpdateHeldTransaction{
			MerchantAuthentication: c.GetAuthentication(),
			RefID: tranx.RefId,
			HeldTransactionRequest: HeldTransactionRequest{
				Action:     method,
				RefTransID: tranx.RefId,
			},
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := c.SendRequest(req)
	var dat TransactionResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (t PreviousTransaction) Approve(c Client) (*TransactionResponse, error) {
	response, err := c.SendTransactionUpdate(t, "approve")
	return response, err
}

func (t PreviousTransaction) Decline(c Client) (*TransactionResponse, error) {
	response, err := c.SendTransactionUpdate(t, "decline")
	return response, err
}

func (c Client) SendGetUnsettled() (*TransactionsList, error) {
	action := UnsettledTransactionsRequest{
		GetUnsettledTransactionListRequest: GetUnsettledTransactionListRequest{
			MerchantAuthentication: c.GetAuthentication(),
			Status:                 "pendingApproval",
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := c.SendRequest(req)
	var dat TransactionsList
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}
