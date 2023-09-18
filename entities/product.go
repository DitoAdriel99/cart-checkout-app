package entities

import (
	"go-learn/library/errbank"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

const (
	ErrAlreadyInCart errbank.Error = "You Have This Item In your Cart!"
)

type Product struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       int        `json:"price"`
	Qty         int        `json:"quantity"`
	QtyReq      int        `json:"quantity_request,omitempty"`
	Rating      float64    `json:"rating"`
	Image       string     `json:"image"`
	Type        string     `json:"type"`
	Banner      string     `json:"banner"`
	Info        string     `json:"info,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (l Product) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.Description, validation.Required),
		validation.Field(&l.Image, validation.Required),
		validation.Field(&l.Price, validation.Required),
		validation.Field(&l.Qty, validation.Required),
		validation.Field(&l.Type, validation.Required),
		validation.Field(&l.Banner, validation.Required),
		validation.Field(&l.Info, validation.Required),
	)
}

type CartsPayload []CartPayload

type CartPayload struct {
	ProductsID uuid.UUID `json:"product_id"`
	Qty        int       `json:"quantity"`
}

type CartDeletePayload struct {
	ProductsID []uuid.UUID `json:"products_id"`
}
