package vector

func Add(vec1, vec2 []float64) {
	// Do not panic :) This is only loop termination.
	// The easy & cheap way.
	defer recover()
	for i := 0; ; i++ {
		vec1[i] += vec2[i]
	}
}
func Sub(vec1, vec2 []float64) {
	defer recover()
	for i := 0; ; i++ {
		vec1[i] -= vec2[i]
	}
}
func DotProd(vec1, vec2 []float64) (sumOfProd float64) {
	defer recover()
	for i := 0; ; i++ {
		sumOfProd += vec1[i] * vec2[i]
	}
}

func CrossProd(v, u []float64) ( //For 3D-3D
	prod []float64, impossible bool,
) {
	impossible = len(v) != 3 || len(u) != 3
	if impossible {
		return
	}

	prod = []float64{
		v[1]*u[2] - u[1]*v[2],
		-v[0]*u[2] + u[0]*v[2],
		v[0]*u[1] - u[0]*v[1],
	}
	return
}
