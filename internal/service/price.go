package service

import (
	"Taxi/internal/models"
	"math"
)

type PriceService struct {}

func NewPriceService() *PriceService{
	return &PriceService{}
}

func (c *PriceService) CalculatePrice(req models.RideRequest) models.RideResponse{
	var distance float64
	var duration float64
	var weather string
	var fare = req.Fare

	var boardingPrice int
	var priceForKm int
	var priceForMinute int
	var weatherMult float64
	var initialPrice float64

	switch fare{
	case "economy":
		boardingPrice = 9000
		priceForKm = 2200
		priceForMinute = 600
	case "comfort":
		boardingPrice = 13000
		priceForKm = 2800
		priceForMinute = 800
	case "comfort_plus":
		boardingPrice = 18000
		priceForKm = 3500
		priceForMinute = 1100
	case "business":
		boardingPrice = 29000
		priceForKm = 5500
		priceForMinute = 1800
	default:
		boardingPrice = 9000
		priceForKm = 2200
		priceForMinute = 600
	}

	switch weather {
	case "light-rain", "rain":
    	weatherMult = 1.2
	case "heavy-rain", "showers", "thunderstorm-with-rain":
    	weatherMult = 1.4
	case "light-snow", "snow":
    	weatherMult = 1.3
	case "snow-showers":
    	weatherMult = 1.6
	default:
    	weatherMult = 1.0
	}
	initialPrice := (float64(boardingPrice) + (distance * float64(priceForKm)) + (duration * float64(priceForMinute))) * weatherMult
	return models.RideResponse{
		Price: int(math.Ceil(initialPrice)),
		Currency: "RUB",
	}
}