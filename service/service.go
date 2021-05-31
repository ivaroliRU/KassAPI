package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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

	req, _ := http.NewRequest("POST", c.BaseURL+ChargeEndpoint, payloadBuf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(c.User, c.Password))

	client := &http.Client{}
	res, _ := client.Do(req)

	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))

	return nil
}
