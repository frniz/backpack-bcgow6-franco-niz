package products

import "fmt"

type repositoryMock struct {
	flagGetAllByID bool
	prods          []Product
}

func NewRepositoryMock() repositoryMock {
	return repositoryMock{}
}

func (r *repositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	r.flagGetAllByID = true
	return []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}, nil
}

type repositoryError struct{}

func NewRepositoryError() repositoryError {
	return repositoryError{}
}

func (r *repositoryError) GetAllBySeller(sellerID string) ([]Product, error) {
	return nil, fmt.Errorf("error test")
}
