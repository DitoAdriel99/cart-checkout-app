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
	CartID      *uuid.UUID `json:"cart_id,omitempty" db:"cart_id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       int        `json:"price"`
	TotPrice    int        `json:"total_price,omitempty"`
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
	CartID []uuid.UUID `json:"cart_id"`
}

func (cdp CartDeletePayload) Validate() error {
	return validation.ValidateStruct(&cdp,
		validation.Field(&cdp.CartID, validation.Required),
	)
}

type CheckoutNeed struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	CartID      uuid.UUID `json:"cart_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
	QuantityReq int       `json:"quantity_request"`
	Price       int       `json:"price"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CheckoutPayload struct {
	CartsID []uuid.UUID `json:"cart_id"`
}

func (cdp CheckoutPayload) Validate() error {
	return validation.ValidateStruct(&cdp,
		validation.Field(&cdp.CartsID, validation.Required),
	)
}

func (cp CartPayload) Validate() error {
	return validation.ValidateStruct(&cp,
		validation.Field(&cp.ProductsID, validation.Required),
		validation.Field(&cp.Qty, validation.Required, validation.Min(1)),
	)
}

func (cp CartsPayload) Validate() error {
	for _, payload := range cp {
		if err := payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}
