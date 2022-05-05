package vector

/// Multiplies vector `vec` in place by given `scalar`.
func MulByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] *= scalar
	}
}

/// Returns product of vector `vec` and `scalar` as newly created slice.
func ProdWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	MulByScalar(prod, scalar)
	return
}

/// Divides vector `vec` in place by given `scalar`.
func DivByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] /= scalar
	}
}

/// Returns quotient of vector `vec` and `scalar` as newly created slice.
func QuotWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	DivByScalar(prod, scalar)
	return
}
