package product_service

import (
	"go-learn/entities"
	"go-learn/library/jwt_parse"
)

func (s *_Service) GetCart(bearer string) ([]entities.Product, error) {
	claims, err := jwt_parse.GetClaimsFromToken(bearer)
	if err != nil {
		return nil, err
	}

	resp, err := s.repo.ProductRepo.GetCart(claims.Email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
