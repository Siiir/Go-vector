package vector

// Multiplies vector `vec` in place by given `scalar`.
func MulByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] *= scalar
	}
}

// Returns product of vector `vec` and `scalar` as newly created slice.
func ProdWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	MulByScalar(prod, scalar)
	return
}

// Divides vector `vec` in place by given `scalar`.
func DivByScalar(vec []float64, scalar float64) {
	for i := range vec {
		vec[i] /= scalar
	}
}

// Returns quotient of vector `vec` and `scalar` as newly created slice.
func QuotWithScalar(vec []float64, scalar float64) (prod []float64) {
	prod = Clone(vec)
	DivByScalar(prod, scalar)
	return
}

/*
Returns `vec` to the power of `exp`.
If exp mod 2 == 0, returns instance of float64 .
Otherwise returns instance of []float64 .
*/
func Pow(vec []float64, exp uint) (scalarOrVec interface{}) {
	doReturnVec := exp&1 == 1
	y := 1.0
	mult := Sq(vec)
	for {
		exp >>= 1
		if exp == 0 {
			break
		}
		if exp&1 == 1 {
			y *= mult
		}
		mult *= mult
	}
	if doReturnVec {
		return ProdWithScalar(vec, y)
	}
	return y
}
