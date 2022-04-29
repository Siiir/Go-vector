package vector

func MultByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] *= scalar
	}
}
func DivByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] /= scalar
	}
}
