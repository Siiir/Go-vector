package vector

func Normalize(vec []float64) {
	DivByScalar(vec, CalcNorm(vec))
}
