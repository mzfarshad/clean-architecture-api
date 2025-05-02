package dto

type (
	Mapper[I, O any] func(I) O
)

func (convert Mapper[I, O]) PtrO() func(I) *O {
	return func(input I) *O {
		output := convert(input)
		return &output
	}
}

func (convert Mapper[I, O]) PtrI() func(*I) O {
	return func(input *I) O {
		if input != nil {
			return convert(*input)
		}
		var zero O
		return zero
	}
}

func (convert Mapper[I, O]) PtrIO() func(*I) *O {
	return func(input *I) *O {
		if input != nil {
			return convert.PtrO()(*input)
		}
		return nil
	}
}
