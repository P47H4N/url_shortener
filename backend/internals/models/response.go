package models

type Response struct {
	Status  string `json:"status" example:"success/failed"`
	Error   string `json:"error,omitempty" example:"Error Message. (If have any)"`
	Message string `json:"message,omitempty" example:"Success Message. (If have any)"`
	Data    any    `json:"data,omitempty"`
}
