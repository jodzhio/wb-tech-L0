package ordgen

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jodzhio/wb-tech-L0/models"
)

func GenerateOrder() models.Order {
	delivery := models.Delivery{
		Name:    gofakeit.Name(),
		Phone:   gofakeit.Phone(),
		Zip:     gofakeit.Zip(),
		City:    gofakeit.City(),
		Address: gofakeit.Address().Address,
		Region:  gofakeit.City(),
		Email:   gofakeit.Email(),
	}

	item := models.Item{
		ChrtID:      int(gofakeit.Uint32()),
		TrackNumber: gofakeit.Word(),
		Price:       int(gofakeit.Price()),
		Rid:         gofakeit.Word(),
		Name:        gofakeit.FirstName(),
		Sale:        gofakeit.Number(0, 100000),
		Size:        strconv.Itoa(gofakeit.Number(1, 100)),
		TotalPrice:  int(gofakeit.Price(0, 100000)),
		NmID:        gofakeit.Number(0, 100000),
		Brand:       gofakeit.Name(),
		Status:      gofakeit.HTTPStatusCode(),
	}

	payment := models.Payment{
		Transaction:  gofakeit.UUID(),
		RequestID:    gofakeit.UUID(),
		Currency:     gofakeit.CurrencyShort(),
		Provider:     gofakeit.Name(),
		PaymentDT:    gofakeit.Date(),
		Bank:         gofakeit.Name(),
		DeliveryCost: gofakeit.Number(0, 100000),
		GoodsTotal:   gofakeit.Number(0, 100000),
		CustomFee:    gofakeit.Number(0, 100000),
	}

	order := models.Order{
		OrderUID:          gofakeit.UUID(),
		TrackNumber:       gofakeit.Name(),
		Entry:             gofakeit.Name(),
		Delivery:          delivery,
		Payment:           payment,
		Items:             []models.Item{item},
		Locale:            gofakeit.CountryAbr(),
		InternalSignature: gofakeit.UUID(),
		CustomerID:        gofakeit.UUID(),
		DeliveryService:   gofakeit.Name(),
		Shardkey:          gofakeit.UUID(),
		SmID:              gofakeit.Number(0, 100000),
		DateCreated:       gofakeit.Date(),
		OofShard:          strconv.Itoa(gofakeit.Number(1, 100)),
	}

	return order
}
