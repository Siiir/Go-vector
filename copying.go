package vector

/// Makes a deep copy (named `clone`) of a given []float64.
func Clone(vec []float64) (cl []float64) {
	cl = make([]float64, len(vec))
	copy(cl, vec)
	return
}
