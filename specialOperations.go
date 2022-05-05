package vector

// Normalizes vector in place.
func Normalize(vec []float64) {
	DivByScalar(vec, CalcNorm(vec))
}

// Returns a new normalized vector.
func Normalized(vec []float64) (normalized []float64) {
	clone := Clone(vec)
	Normalize(clone)
	return clone
}

// Every `element` in given slice is transformed into f(`element`)
// in place.
func Remap(sl []float64, f func(float64) float64) {
	for i := range sl {
		sl[i] = f(sl[i])
	}
}

/*
Return new slice with every element
being an image of its counterpart from given slice.
Where `f` is projecting function.
*/
func Remaped(sl []float64,
	f func(float64) float64,
) (new []float64) {
	new = make([]float64, len(sl))
	for i := range sl {
		new[i] = f(sl[i])
	}
	return
}
