package vector

/// Makes a deep copy (named `clone`) of a given []float64.
func Clone(vec []float64) (clone []float64) {
	clone = make([]float64, len(vec))
	copy(clone, vec)
	return
}
