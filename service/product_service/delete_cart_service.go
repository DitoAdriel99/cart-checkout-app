package product_service

import (
	"go-learn/library/jwt_parse"

	"github.com/google/uuid"
)

func (s *_Service) DeleteCart(bearer string, cart_id []uuid.UUID) error {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return err
	}
	for _, v := range cart_id {
		_, err := s.repo.ProductRepo.GetCartDetail(v)
		if err != nil {
			return err
		}
	}

	if err := s.repo.ProductRepo.DeleteCart(claims.Email, cart_id); err != nil {
		return err
	}

	return nil
}
