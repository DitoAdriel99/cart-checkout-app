package product_service

import (
	"go-learn/library/jwt_parse"

	"github.com/google/uuid"
)

func (s *_Service) DeleteCart(bearer string, products_id []uuid.UUID) error {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return err
	}
	for _, v := range products_id {
		_, err := s.repo.ProductRepo.Detail(v)
		if err != nil {
			return err
		}
	}

	if err := s.repo.ProductRepo.DeleteCart(claims.Email, products_id); err != nil {
		return err
	}

	return nil
}
