package vector

func Negate(vec []float64) {
	for i := range vec {
		vec[i] = -vec[i]
	}
}

func Negated(vec []float64) (negated []float64) {
	negated = make([]float64, len(vec))
	for i := range vec {
		negated[i] = -vec[i]
	}
	return
}
