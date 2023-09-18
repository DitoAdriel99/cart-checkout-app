package product_service

import (
	"go-learn/entities"
	"go-learn/library/jwt_parse"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Checkout(bearer string, payload *entities.CheckoutPayload) error {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return err
	}

	objects := make([]entities.CheckoutNeed, 0)
	for _, v := range payload.CartsID {
		cartDetail, err := s.repo.ProductRepo.GetCartDetail(v)
		if err != nil {
			return err
		}

		newIdCheckout, _ := uuid.NewUUID()

		totPrice := cartDetail.QtyReq * cartDetail.Price

		object := entities.CheckoutNeed{
			ID:          newIdCheckout,
			Email:       claims.Email,
			CartID:      v,
			ProductID:   cartDetail.ID,
			Quantity:    cartDetail.Qty,
			QuantityReq: cartDetail.QtyReq,
			Price:       cartDetail.Price,
			TotalPrice:  totPrice,
			CreatedAt:   time.Now().Local(),
			UpdatedAt:   time.Now().Local(),
		}
		objects = append(objects, object)
	}

	if err := s.repo.ProductRepo.Checkout(claims.Email, objects); err != nil {
		return err
	}

	return nil
}
