package vector

func Clone(vec []float64) (clone []float64) {
	clone = make([]float64, len(vec))
	copy(clone, vec)
	return
}
