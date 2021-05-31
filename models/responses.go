package models

type ErrorMessage struct {
	Code    string `json:"code"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Response struct {
	Success bool          `json:"success"`
	Id      string        `id:"success"`
	Created int64         `json:"created"`
	Error   *ErrorMessage `json:"error"`
}
