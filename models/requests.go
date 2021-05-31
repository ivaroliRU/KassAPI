package models

type CreateChargeRequest struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Order       string `json:"order"`
	Recipient   string `json:"recipient"`
	Terminal    int    `json:"terminal"`
	ExpiresIn   int    `json:"expires_in"`
	NotifyURL   string `json:"notify_url"`
}
