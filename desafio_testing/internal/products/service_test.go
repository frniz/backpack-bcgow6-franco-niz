package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllByID(t *testing.T) {
	ss := NewRepositoryMock()
	service := NewService(&ss)
	ss.prods = append(ss.prods, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	ss.flagGetAllByID = false

	result, err := service.GetAllBySeller("")

	assert.Equal(t, ss.prods, result)
	assert.Nil(t, err)
	assert.True(t, ss.flagGetAllByID)
}
