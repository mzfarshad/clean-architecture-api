package dto

// List makes []O using []I.
func List[O any, L ~[]I, I any](inputs L, convert func(I) O) []O {
	if convert == nil || inputs == nil || len(inputs) == 0 {
		return nil
	}
	outputs := make([]O, len(inputs))
	for i := range inputs {
		outputs[i] = convert(inputs[i])
	}
	return outputs
}
