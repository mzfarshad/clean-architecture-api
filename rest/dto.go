package rest

import (
	file "github.com/h2non/filetype/types"
	"github.com/mzfarshad/music_store_api/pkg/dto"
)

// idto represents any single or list input/output data transfer object.
// It prevents directly sending domain entity/aggregate to the client side.
// As a rest input idto, we can use some functionalities like parsing, and validating input data.
type idto interface{ dto() }

type DTO struct{}

func (DTO) dto() {}

// Map implements idto interface and can be use when we need a map response data.
type Map map[string]any

func (Map) dto() {}

// List implements idto interface and represents a list of the given idto type.
type List[D idto] []D

func (List[D]) dto() {}

func NewList[D idto, SI ~[]I, I any](inputs SI, mapper func(I) D) List[D] {
	return dto.List(inputs, mapper)
}

type File struct {
	DTO
	Name  string
	Type  file.Type
	Bytes []byte
}
