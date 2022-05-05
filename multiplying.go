package vector

func ForceDotProd(vec1, vec2 []float64) (sumOfProd float64) {
	defer recover()
	for i := 0; ; i++ {
		sumOfProd += vec1[i] * vec2[i]
	}
	return
}

func DotProd(vec1, vec2 []float64) (sumOfProd float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	sumOfProd = ForceDotProd(vec1, vec2)
	return
}

func UncheckedCrossProd(v, u []float64) []float64 {
	return []float64{
		v[1]*u[2] - u[1]*v[2],
		-v[0]*u[2] + u[0]*v[2],
		v[0]*u[1] - u[0]*v[1],
	}
}

func NilableCrossProd(v, u []float64) (prodOrNil []float64) {
	impossible := len(v) != 3 || len(u) != 3
	if impossible {
		return
	}

	return UncheckedCrossProd(v, u)
}

func CrossProd(v, u []float64) ( //For 3D-3D
	prodOrNil []float64, impossible bool,
) {
	prodOrNil = NilableCrossProd(v, u)
	impossible = prodOrNil == nil
	return
}
