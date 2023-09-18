package product_service

import (
	"go-learn/entities"
	"go-learn/library/meta"
	"go-learn/repositories"

	"github.com/google/uuid"
)

type Contract interface {
	Create(payload *entities.Product) error
	Update(id uuid.UUID, payload *entities.Product) error
	Delete(id uuid.UUID) error
	Detail(id uuid.UUID) (*entities.Product, error)
	GetAll(m *meta.Metadata) ([]entities.Product, error)
	AddToCart(payload entities.CartsPayload, bearer string) error
	GetCart(bearer string) ([]entities.Product, error)
	DeleteCart(bearer string, cart_id []uuid.UUID) error
	Checkout(bearer string, payload *entities.CheckoutPayload) error
}

type _Service struct {
	repo *repositories.Repo
}

func NewProductService(repo *repositories.Repo) Contract {
	return &_Service{repo}
}
