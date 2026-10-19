package models

type RideRequest struct {
	City string `json:"city"`
	From string `json:"from"`
	To string `json:"to"`
	Fare string `json:"fare"`
}

type RideResponse struct {
	Price int `json:"price"`
	Currency string `json:"currency"`
}