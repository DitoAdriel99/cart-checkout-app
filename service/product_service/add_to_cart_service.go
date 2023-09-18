package product_service

import (
	"go-learn/entities"
	"go-learn/library/jwt_parse"
)

func (s *_Service) AddToCart(payload entities.CartsPayload, bearer string) error {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return err
	}
	for _, v := range payload {
		_, err := s.repo.ProductRepo.Detail(v.ProductsID)
		if err != nil {
			return err
		}
	}

	err = s.repo.ProductRepo.AddToCart(payload, claims.Email)
	if err != nil {
		return err
	}

	return nil
}
