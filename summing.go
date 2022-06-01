package vector

/*
Adds in place `vec2` to `vec1`.

Dimensions of `vec1` that are not present in `vec2` are not incremented.
Dimensions of `vec2` that are not present in `vec1` are ignored.
*/
func ForceAdd(vec1, vec2 []float64) {
	// Do not panic :) This is only loop termination.
	// The easy & cheap way.
	defer callRecover()
	for i := 0; ; i++ {
		vec1[i] += vec2[i]
	}
}

// If vectors are of the same length, adds in place `vec2` to `vec1`
// and returns false.
// Otherwise returns true.
func Add(vec1, vec2 []float64) (impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	ForceAdd(vec1, vec2)
	return
}

/*
If vectors are of the same length,
	returns their sum as newly created []float64
		and false as the second value.
Otherwise returns (nil, true)
*/
func Sum(vec1, vec2 []float64) (sum []float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	sum = Clone(vec1)
	ForceAdd(sum, vec2)
	return
}

/*
Return sum of vectors `vec1` & `vec2`.
Returned vector will have max(`vec1`.len(),`vec2`.len()) length.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func LongestSumSlice(vec1, vec2 []float64) (long []float64) {
	// Loading slice
	long = make([]float64, maxInt(len(vec1), len(vec2)))
	copy(long, vec1)
	// Subtracting
	ForceAdd(long, vec2)
	// Rest
	return
}

/*
Subtracts in place `vec2` from `vec1`.

Dimensions of `vec1` that are not present in `vec2` are not deincremented.
Dimensions of `vec2` that are not present in `vec1` are ignored.
*/
func ForceSub(vec1, vec2 []float64) {
	defer callRecover()
	for i := 0; ; i++ {
		vec1[i] -= vec2[i]
	}
}

// If vectors are of the same length, subtracts in place `vec2` from `vec1`
// and returns false.
// Otherwise returns true.
func Sub(vec1, vec2 []float64) (impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	ForceSub(vec1, vec2)
	return
}

/*
If vectors are of the same length,
	returns their difference as newly created []float64
		and false as the second value.
Otherwise returns (nil, true)
*/
func Diff(vec1, vec2 []float64) (diff []float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	diff = Clone(vec1)
	ForceSub(diff, vec2)
	return
}

/*
Return difference of vectors `vec1` & `vec2`.
Returned vector will have max(`vec1`.len(),`vec2`.len()) length.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func LongestDiffSlice(vec1, vec2 []float64) (long []float64) {
	// Loading slice
	long = make([]float64, maxInt(len(vec1), len(vec2)))
	copy(long, vec1)
	// Subtracting
	ForceSub(long, vec2)
	// Rest
	return
}
