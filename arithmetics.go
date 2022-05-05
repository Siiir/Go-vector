package vector

func ForceAdd(vec1, vec2 []float64) {
	// Do not panic :) This is only loop termination.
	// The easy & cheap way.
	defer recover()
	for i := 0; ; i++ {
		vec1[i] += vec2[i]
	}
}
func Add(vec1, vec2 []float64) (impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	ForceAdd(vec1, vec2)
	return
}

func Sum(vec1, vec2 []float64) (sum []float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	sum = make([]float64, len(vec1))
	copy(sum, vec1)
	ForceAdd(sum, vec2)
	return
}

func LongestSum(vec1, vec2 []float64) (long []float64) {
	//! The longest slice is copied.
	//! We force addition of shorter slice to longer,
	//!  so that we archieve longer resulting vector,
	//! If we forced addition of larger to smaller,
	//!  the larger slice would have to be truncated (tail'd be ommited)
	if len(vec1) < len(vec2) {
		long = make([]float64, len(vec2))
		copy(long, vec2)
	} else {
		long = make([]float64, len(vec1))
		copy(long, vec1)
	}
	return
}

func ForceSub(vec1, vec2 []float64) {
	defer recover()
	for i := 0; ; i++ {
		vec1[i] -= vec2[i]
	}
}

func Sub(vec1, vec2 []float64) (impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	ForceSub(vec1, vec2)
	return
}

func LongestDiff(vec1, vec2 []float64) (long []float64)

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
