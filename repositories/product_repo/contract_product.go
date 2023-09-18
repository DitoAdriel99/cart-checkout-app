package product_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"
	"go-learn/library/meta"

	"github.com/google/uuid"
)

type _ProductRepoImp struct {
	conn *sql.DB
}

type ProductContract interface {
	Create(pr *entities.Product) error
	Detail(id uuid.UUID) (*entities.Product, error)
	GetAll(m *meta.Metadata) ([]entities.Product, error)
	Update(pr *entities.Product) error
	Delete(id uuid.UUID) error
	AddToCart(payload entities.CartsPayload, email string) error
	GetCart(email string) ([]entities.Product, error)
	DeleteCart(email string, products_id []uuid.UUID) error
	Checkout(email string, payload []entities.CheckoutNeed) error
	GetCartDetail(cartID uuid.UUID) (*entities.Product, error)
}

func NewProductRepositories() ProductContract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_ProductRepoImp{
		conn: conn,
	}
}
