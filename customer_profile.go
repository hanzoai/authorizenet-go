package authorizenet

import (
	"encoding/json"
)

func (c *Client) GetPaymentProfileIds(month string, method string) (*GetCustomerPaymentProfileListResponse, error) {
	action := GetCustomerPaymentProfileListRequest{
		GetCustomerPaymentProfileList: GetCustomerPaymentProfileList{
			MerchantAuthentication: c.GetAuthentication(),
			SearchType:             method,
			Month:                  month,
			Sorting: Sorting{
				OrderBy:         "id",
				OrderDescending: "false",
			},
			Paging: Paging{
				Limit:  "1000",
				Offset: "1",
			},
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat GetCustomerPaymentProfileListResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (profile Profile) CreateProfile(c Client) (*CustomProfileResponse, error) {
	res, err := c.CreateProfile(profile)
	return res, err
}

func (profile Profile) CreateShipping(c Client) (*CreateCustomerShippingAddressResponse, error) {
	res, err := c.CreateShipping(profile)
	return res, err
}

func (customer Customer) Info(c Client) (*GetCustomerProfileResponse, error) {
	res, err := c.GetProfile(customer)
	return res, err
}

func (customer Customer) Validate(c Client) (*ValidateCustomerPaymentProfileResponse, error) {
	res, err := c.ValidatePaymentProfile(customer)
	return res, err
}

func (customer Customer) DeleteProfile(c Client) (*MessagesResponse, error) {
	res, err := c.DeleteProfile(customer)
	return res, err
}

func (customer Customer) DeletePaymentProfile(c Client) (*MessagesResponse, error) {
	res, err := c.DeletePaymentProfile(customer)
	return res, err
}

func (customer Customer) DeleteShippingProfile(c Client) (*MessagesResponse, error) {
	res, err := c.DeleteShippingProfile(customer)
	return res, err
}

func (payment CustomerPaymentProfile) Add(c Client) (*CustomerPaymentProfileResponse, error) {
	res, err := c.CreatePaymentProfile(payment)
	return res, err
}

func (res GetCustomerProfileResponse) PaymentProfiles() []GetPaymentProfiles {
	return res.Profile.PaymentProfiles
}

func (res GetCustomerProfileResponse) ShippingProfiles() []GetShippingProfiles {
	return res.Profile.ShippingProfiles
}

func (res GetCustomerProfileResponse) Subscriptions() []string {
	return res.SubscriptionIds
}

func (profile Profile) UpdateProfile(c Client) (*MessagesResponse, error) {
	res, err := c.UpdateProfile(profile)
	return res, err
}

func (profile Profile) UpdatePaymentProfile(c Client) (*MessagesResponse, error) {
	res, err := c.UpdatePaymentProfile(profile)
	return res, err
}

func (profile Profile) UpdateShippingProfile(c Client) (*MessagesResponse, error) {
	res, err := c.UpdateShippingProfile(profile)
	return res, err
}

func (c *Client) GetProfileIds() ([]string, error) {
	action := GetCustomerProfileIdsRequest{
		CustomerProfileIdsRequest: CustomerProfileIdsRequest{
			MerchantAuthentication: c.GetAuthentication(),
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return []string{}, err
	}
	res, err := c.SendRequest(req)
	var dat CustomerProfileIdsResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return []string{}, err
	}
	return dat.Ids, err
}

func (c *Client) ValidatePaymentProfile(customer Customer) (*ValidateCustomerPaymentProfileResponse, error) {
	action := ValidateCustomerPaymentProfileRequest{
		ValidateCustomerPaymentProfile: ValidateCustomerPaymentProfile{
			MerchantAuthentication:   c.GetAuthentication(),
			CustomerProfileID:        customer.ID,
			CustomerPaymentProfileID: customer.PaymentID,
			ValidationMode:           c.Mode,
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat ValidateCustomerPaymentProfileResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (c *Client) GetProfile(customer Customer) (*GetCustomerProfileResponse, error) {
	action := CustomerProfileRequest{
		GetCustomerProfile: GetCustomerProfile{
			MerchantAuthentication: c.GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat GetCustomerProfileResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (c *Client) CreateProfile(profile Profile) (*CustomProfileResponse, error) {
	action := CreateCustomerProfileRequest{
		CreateCustomerProfile: CreateCustomerProfile{
			MerchantAuthentication: c.GetAuthentication(),
			Profile:                profile,
			ValidationMode:         c.Mode,
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}

	res, err := c.SendRequest(req)
	var dat CustomProfileResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (c *Client) CreateShipping(profile Profile) (*CreateCustomerShippingAddressResponse, error) {
	action := CreateCustomerShippingAddressRequest{
		CreateCustomerShippingAddress: CreateCustomerShippingAddress{
			MerchantAuthentication: c.GetAuthentication(),
			Address:                profile.Shipping,
			CustomerProfileID:      profile.CustomerProfileId,
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat CreateCustomerShippingAddressResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (c *Client) UpdateProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerProfileRequest{
		UpdateCustomerProfile: UpdateCustomerProfile{
			MerchantAuthentication: c.GetAuthentication(),
			Profile:                profile,
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) UpdatePaymentProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerPaymentProfileRequest{
		UpdateCustomerPaymentProfile: UpdateCustomerPaymentProfile{
			CustomerProfileID:      profile.CustomerProfileId,
			MerchantAuthentication: c.GetAuthentication(),
			UpPaymentProfile: UpPaymentProfile{
				BillTo:                   profile.PaymentProfiles.BillTo,
				Payment:                  profile.PaymentProfiles.Payment,
				CustomerPaymentProfileID: profile.PaymentProfileId,
			},
			ValidationMode: c.Mode,
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) UpdateShippingProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerShippingAddressRequest{
		UpdateCustomerShippingAddress: UpdateCustomerShippingAddress{
			CustomerProfileID:      profile.CustomerProfileId,
			MerchantAuthentication: c.GetAuthentication(),
			Address: Address{
				FirstName:         "My",
				LastName:          "Name",
				Address:           "38485 New Road ave.",
				City:              "Los Angeles",
				State:             "CA",
				Zip:               "283848",
				Country:           "USA",
				PhoneNumber:       "8885555555",
				CustomerAddressID: profile.CustomerAddressId,
			},
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) DeleteProfile(customer Customer) (*MessagesResponse, error) {
	action := DeleteCustomerProfileRequest{
		DeleteCustomerProfile: DeleteCustomerProfile{
			MerchantAuthentication: c.GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) MessageResponder(d interface{}) (*MessagesResponse, error) {
	req, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat MessagesResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (c *Client) DeletePaymentProfile(customer Customer) (*MessagesResponse, error) {
	action := DeleteCustomerPaymentProfileRequest{
		DeleteCustomerPaymentProfile: DeleteCustomerPaymentProfile{
			MerchantAuthentication:   c.GetAuthentication(),
			CustomerProfileID:        customer.ID,
			CustomerPaymentProfileID: customer.PaymentID,
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) DeleteShippingProfile(customer Customer) (*MessagesResponse, error) {
	action := DeleteCustomerShippingProfileRequest{
		DeleteCustomerShippingProfile: DeleteCustomerShippingProfile{
			MerchantAuthentication: c.GetAuthentication(),
			CustomerProfileID:      customer.ID,
			CustomerShippingID:     customer.ShippingID,
		},
	}
	dat, err := c.MessageResponder(action)
	return dat, err
}

func (c *Client) CreatePaymentProfile(profile CustomerPaymentProfile) (*CustomerPaymentProfileResponse, error) {
	action := CreateCustomerPaymentProfile{
		CreateCustomerPaymentProfileRequest: CreateCustomerPaymentProfileRequest{
			MerchantAuthentication: c.GetAuthentication(),
			CustomerProfileID:      profile.CustomerProfileID,
			PaymentProfile: PaymentProfile{
				BillTo:                profile.PaymentProfile.BillTo,
				Payment:               profile.PaymentProfile.Payment,
				DefaultPaymentProfile: profile.PaymentProfile.DefaultPaymentProfile,
			},
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat CustomerPaymentProfileResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

type CreateCustomerProfileRequest struct {
	CreateCustomerProfile CreateCustomerProfile `json:"createCustomerProfileRequest"`
}

type CreateCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile                Profile                `json:"profile"`
	ValidationMode         string                 `json:"validationMode"`
}

type CustomerProfiler struct {
	CustomerProfileID         string `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileID  string `json:"customerPaymentProfileId,omitempty"`
	CustomerShippingProfileID string `json:"customerAddressId,omitempty"`
}

type Profile struct {
	MerchantCustomerID string           `json:"merchantCustomerId,omitempty"`
	Description        string           `json:"description,omitempty"`
	Email              string           `json:"email,omitempty"`
	CustomerProfileId  string           `json:"customerProfileId,omitempty"`
	PaymentProfiles    *PaymentProfiles `json:"paymentProfiles,omitempty"`
	PaymentProfileId   string           `json:"customerPaymentProfileId,omitempty"`
	Shipping           *Address         `json:"address,omitempty"`
	CustomerAddressId  string           `json:"customerAddressId,omitempty"`
	PaymentProfile     *PaymentProfile  `json:"paymentProfile,omitempty"`
}

type PaymentProfiles struct {
	CustomerType string  `json:"customerType,omitempty"`
	Payment      Payment `json:"payment,omitempty"`
	BillTo       *BillTo `json:"billTo,omitempty"`
	PaymentId    string  `json:"paymentProfileId,omitempty"`
}

type CustomProfileResponse struct {
	CustomerProfileID             string        `json:"customerProfileId"`
	CustomerPaymentProfileIDList  []string      `json:"customerPaymentProfileIdList"`
	CustomerShippingAddressIDList []interface{} `json:"customerShippingAddressIdList"`
	ValidationDirectResponseList  []string      `json:"validationDirectResponseList"`
	MessagesResponse
}

type CustomerProfileRequest struct {
	GetCustomerProfile GetCustomerProfile `json:"getCustomerProfileRequest"`
}

type GetCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
}

type GetCustomerProfileResponse struct {
	Profile struct {
		PaymentProfiles    []GetPaymentProfiles  `json:"paymentProfiles,omitempty"`
		ShippingProfiles   []GetShippingProfiles `json:"shipToList,omitempty"`
		CustomerProfileID  string                `json:"customerProfileId"`
		MerchantCustomerID string                `json:"merchantCustomerId,omitempty"`
		Description        string                `json:"description,omitempty"`
		Email              string                `json:"email,omitempty"`
	} `json:"profile"`
	SubscriptionIds []string `json:"subscriptionIds"`
	MessagesResponse
}

type DeleteCustomerPaymentProfileRequest struct {
	DeleteCustomerPaymentProfile DeleteCustomerPaymentProfile `json:"deleteCustomerPaymentProfileRequest"`
}

type DeleteCustomerPaymentProfile struct {
	MerchantAuthentication   MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID        string                 `json:"customerProfileId"`
	CustomerPaymentProfileID string                 `json:"customerPaymentProfileId"`
}

type DeleteCustomerShippingProfileRequest struct {
	DeleteCustomerShippingProfile DeleteCustomerShippingProfile `json:"deleteCustomerShippingAddressRequest"`
}

type DeleteCustomerShippingProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	CustomerShippingID     string                 `json:"customerAddressId"`
}

type GetShippingProfiles struct {
	CustomerAddressID string `json:"customerAddressId"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Company           string `json:"company,omitempty"`
	Address           string `json:"address,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Zip               string `json:"zip,omitempty"`
	Country           string `json:"country,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
}

type GetPaymentProfiles struct {
	CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
	Payment                  struct {
		CreditCard struct {
			CardNumber     string `json:"cardNumber"`
			ExpirationDate string `json:"expirationDate"`
		} `json:"creditCard"`
	} `json:"payment"`
	CustomerType string `json:"customerType"`
	BillTo       struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"billTo"`
}

type GetCustomerProfileIdsRequest struct {
	CustomerProfileIdsRequest CustomerProfileIdsRequest `json:"getCustomerProfileIdsRequest"`
}

type CustomerProfileIdsRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
}

type CustomerProfileIdsResponse struct {
	Ids      []string `json:"ids"`
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}

type UpdateCustomerProfileRequest struct {
	UpdateCustomerProfile UpdateCustomerProfile `json:"updateCustomerProfileRequest"`
}

type UpdateCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile                Profile                `json:"profile"`
}

type DeleteCustomerProfileRequest struct {
	DeleteCustomerProfile DeleteCustomerProfile `json:"deleteCustomerProfileRequest"`
}

type DeleteCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
}

type MessagesResponse struct {
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}

type MessageResponse struct {
	ResultCode string `json:"resultCode"`
	Message    struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"message"`
}

type CustomerPaymentProfile struct {
	CustomerProfileID string         `json:"customerProfileId"`
	PaymentProfile    PaymentProfile `json:"paymentProfile"`
}

type CreateCustomerPaymentProfile struct {
	CreateCustomerPaymentProfileRequest CreateCustomerPaymentProfileRequest `json:"createCustomerPaymentProfileRequest"`
}

type CreateCustomerPaymentProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	PaymentProfile         PaymentProfile         `json:"paymentProfile"`
}

type PaymentProfile struct {
	BillTo                *BillTo  `json:"billTo,omitempty"`
	Payment               *Payment `json:"payment,omitempty"`
	DefaultPaymentProfile string   `json:"defaultPaymentProfile,omitempty"`
	PaymentProfileId      string   `json:"paymentProfileId,omitempty"`
}

type CustomerPaymentProfileResponse struct {
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
	ValidationDirectResponse string `json:"validationDirectResponse"`
	MessagesResponse
}

type GetCustomerPaymentProfileListRequest struct {
	GetCustomerPaymentProfileList GetCustomerPaymentProfileList `json:"getCustomerPaymentProfileListRequest"`
}

type GetCustomerPaymentProfileList struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	SearchType             string                 `json:"searchType"`
	Month                  string                 `json:"month"`
	Sorting                Sorting                `json:"sorting"`
	Paging                 Paging                 `json:"paging"`
}

type GetCustomerPaymentProfileListResponse struct {
	GetCustomerPaymentProfileList struct {
		MessagesResponse
		TotalNumInResultSet string `json:"totalNumInResultSet"`
		PaymentProfiles     struct {
			PaymentProfile []PaymentProfile `json:"paymentProfile"`
		} `json:"paymentProfiles"`
	} `json:"getCustomerPaymentProfileListResponse"`
}

type ValidateCustomerPaymentProfileRequest struct {
	ValidateCustomerPaymentProfile ValidateCustomerPaymentProfile `json:"validateCustomerPaymentProfileRequest"`
}

type ValidateCustomerPaymentProfile struct {
	MerchantAuthentication   MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID        string                 `json:"customerProfileId"`
	CustomerPaymentProfileID string                 `json:"customerPaymentProfileId"`
	ValidationMode           string                 `json:"validationMode"`
}

type ValidateCustomerPaymentProfileResponse struct {
	DirectResponse string `json:"directResponse"`
	MessagesResponse
}

type CreateCustomerShippingAddressRequest struct {
	CreateCustomerShippingAddress CreateCustomerShippingAddress `json:"createCustomerShippingAddressRequest"`
}

type CreateCustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId,omitempty"`
	Address                *Address               `json:"address,omitempty"`
}

type CreateCustomerShippingAddressResponse struct {
	CustomerAddressID string `json:"customerAddressId,omitempty"`
	MessagesResponse
}

type UpdateCustomerPaymentProfileRequest struct {
	UpdateCustomerPaymentProfile UpdateCustomerPaymentProfile `json:"updateCustomerPaymentProfileRequest"`
}

type UpdateCustomerPaymentProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	UpPaymentProfile       UpPaymentProfile       `json:"paymentProfile"`
	ValidationMode         string                 `json:"validationMode"`
}

type UpPaymentProfile struct {
	BillTo                   *BillTo `json:"billTo"`
	Payment                  Payment `json:"payment"`
	CustomerPaymentProfileID string  `json:"customerPaymentProfileId"`
}

type UpdateCustomerShippingAddressRequest struct {
	UpdateCustomerShippingAddress UpdateCustomerShippingAddress `json:"updateCustomerShippingAddressRequest"`
}

type UpdateCustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	Address                Address                `json:"address"`
}
