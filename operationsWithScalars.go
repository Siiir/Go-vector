package vector

func MulByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] *= scalar
	}
}
func ProdWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	MulByScalar(prod, scalar)
	return
}
func DivByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] /= scalar
	}
}
func QuotWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	DivByScalar(prod, scalar)
	return
}
