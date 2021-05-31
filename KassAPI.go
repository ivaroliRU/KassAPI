package KassAPI

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ivaroliRU/KassAPI/models"
)

const (
	ProductionURL = "https://api.kass.is/v1"
	SandboxURL    = "https://api.testing.kass.is/v1"

	ChargeEndpoint = "/payments"
)

type Client struct {
	BaseURL  string
	User     string
	Password string
}

// create client
func New(production bool, user string, password string) *Client {
	url := ProductionURL

	if !production {
		url = SandboxURL
	}

	return &Client{
		BaseURL:  url,
		User:     user,
		Password: password,
	}
}

// turn username:password to base64
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// create charge and post
func (c *Client) CreateCharge(amount int, description, imageURL, order, recipient string, terminal, expiresIn int, notifyUrl string) (*models.Response, *error) {
	data := &models.CreateChargeRequest{
		Amount:      amount,
		Description: description,
		ImageURL:    imageURL,
		Order:       order,
		Recipient:   recipient,
		Terminal:    terminal,
		ExpiresIn:   expiresIn,
		NotifyURL:   notifyUrl,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)

	req, _ := http.NewRequest("POST", c.BaseURL+ChargeEndpoint, payloadBuf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(c.User, c.Password))

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, &err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var response models.Response

	json.Unmarshal(body, &response)

	return &response, nil
}
