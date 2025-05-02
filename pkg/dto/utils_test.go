package dto_test

import (
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUtils(t *testing.T) {
	x := "str"
	// Pointer to x
	pX := dto.Ptr[string]()(x)
	if assert.NotNil(t, pX) {
		assert.Equal(t, x, *pX)
	}
	// value of the pointer to x
	vpX := dto.IndirectFunc[string](false)(pX)
	assert.Equal(t, x, vpX)
	// pointer of pointer to x
	ppX := dto.Ptr[*string]()(pX)
	if assert.NotNil(t, ppX) {
		assert.Equal(t, x, **ppX)
		assert.Equal(t, pX, dto.IndirectFunc[*string](false)(ppX))
	}
	// safe mode && nil input >>> zero value
	assert.Equal(t, nil, dto.IndirectFunc[any](true)(nil))
	assert.Equal(t, false, dto.IndirectFunc[bool](true)(nil))
	assert.Equal(t, "", dto.IndirectFunc[string](true)(nil))
	assert.Equal(t, struct{}{}, dto.IndirectFunc[struct{}](true)(nil))

	// unsafe mode && nil input >>> panics
	assert.Panics(t, func() { dto.IndirectFunc[any](false)(nil) })
}

func TestAddress(t *testing.T) {
	x := "str"
	pX := dto.Address(x)
	assert.Equal(t, x, *pX)

	assert.NotPanics(t, func() { dto.Address[*bool](nil) })
	pBool := dto.Address[*bool](nil)
	log.Printf(">>>> pBool: %v, is nil: %v, *pBool: %v *pBool is nil: %v", pBool, pBool == nil, *pBool, *pBool == nil)
}
