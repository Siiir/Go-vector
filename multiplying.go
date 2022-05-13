package vector

/*
Returns standard dot product of two given vectors.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
And so it can operate on vectors of apparently different number of dimensions.
*/
func ForceDotProd(vec1, vec2 []float64) (sumOfProd float64) {
	defer callRecover()
	for i := 0; ; i++ {
		sumOfProd += vec1[i] * vec2[i]
	}
}

// Returns standard dot product of two given vectors.
// [last argument `impossible` = true] means result was not computed
// due to slices having different lengths.
func DotProd(vec1, vec2 []float64) (sumOfProd float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	sumOfProd = ForceDotProd(vec1, vec2)
	return
}

// `UncheckedCrossProd` trust caller that given vectors v, u have at least 3-dimesions
// and calculates their cross product as if they had exactly 3 dimensions.
func UncheckedCrossProd(v, u []float64) []float64 {
	return []float64{
		v[1]*u[2] - u[1]*v[2],
		-v[0]*u[2] + u[0]*v[2],
		v[0]*u[1] - u[0]*v[1],
	}
}

// Returns cross product of two vectors if they are 3-dimensional.
// nil otherwise.
func NilableCrossProd(v, u []float64) (prodOrNil []float64) {
	impossible := len(v) != 3 || len(u) != 3
	if impossible {
		return
	}

	return UncheckedCrossProd(v, u)
}

// Returns cross product of two vectors and false,
// if they are 3-dimensional. (nil, true) otherwise.
func CrossProd(v, u []float64) ( //For 3D-3D
	prodOrNil []float64, impossible bool,
) {
	prodOrNil = NilableCrossProd(v, u)
	impossible = prodOrNil == nil
	return
}
