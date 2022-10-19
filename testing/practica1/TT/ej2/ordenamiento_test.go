package ej2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	sliceDePrueba := []int{7, 2, 8, 1, 5, 3, 4, 6}
	sliceDeResultado := []int{1, 2, 3, 4, 5, 6, 7, 8}

	sliceQueRecibo := OrdenarSliceDeEnteros(sliceDePrueba)

	assert.Equal(t, sliceDeResultado, sliceQueRecibo)
}
