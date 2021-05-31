package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ivaroliRU/KassAPI/models"
)

const (
	ProductionURL = "https://api.kass.is/v1"
	SandboxURL    = "https://api.testing.kass.is/v1"
)

type Client struct {
	BaseURL  string
	User     string
	Password string
}

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

func (c *Client) CreateCharge() *error {
	data := &models.CreateChargeRequest{
		Amount:      100,
		Description: "",
		Order:       "asd32434ifj",
		ImageURL:    "asdf",
		Recipient:   "eæpi53pm",
		Terminal:    1,
		ExpiresIn:   90,
		NotifyURL:   "https://example.com/callbacks/kass",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)

	req, _ := http.NewRequest("POST", c.BaseURL, payloadBuf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)

	if e != nil {
		return &e
	}

	defer res.Body.Close()

	return nil
}
