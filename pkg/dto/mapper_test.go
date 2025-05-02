package dto_test

import (
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapper(t *testing.T) {
	var (
		input    = 3
		expected = "3"
	)
	mapper := dto.Mapper[int, string](strconv.Itoa)
	if got := mapper(input); assert.NotEmpty(t, got) {
		assert.Equal(t, expected, got)
	}

	mapper1 := mapper.PtrO()
	if got := mapper1(input); assert.NotNil(t, got) {
		assert.Equal(t, expected, *got)
	}

	mapper2 := mapper.PtrI()
	if got := mapper2(&input); assert.NotEmpty(t, got) {
		assert.Equal(t, expected, got)
	}

	mapper3 := mapper.PtrIO()
	if got := mapper3(&input); assert.NotNil(t, got) {
		assert.Equal(t, expected, *got)
	}
}
