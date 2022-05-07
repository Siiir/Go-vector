package vector

/*
Adds in place `vec2` to `vec1`.

Dimensions of `vec1` that are not present in `vec2` are not incremented.
Dimensions of `vec2` that are not present in `vec1` are omitted.
*/
func ForceAdd(vec1, vec2 []float64) {
	// Do not panic :) This is only loop termination.
	// The easy & cheap way.
	defer recover()
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

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func LongestSumSlice(vec1, vec2 []float64) (long []float64) {
	// The longest slice is copied.
	// We force addition of shorter slice to longer,
	//  so that we archieve longer resulting vector,
	// If we forced addition of larger to smaller,
	//  the larger slice would have to be truncated (tail'd be ommited)
	var shorter, longer []float64
	if len(vec1) < len(vec2) {
		shorter = vec1
		longer = vec2
	} else {
		shorter = vec2
		longer = vec1
	}

	long = Clone(longer)
	ForceAdd(long, shorter)

	return
}

/*
Subtracts in place `vec2` from `vec1`.

Dimensions of `vec1` that are not present in `vec2` are not deincremented.
Dimensions of `vec2` that are not present in `vec1` are omitted.
*/
func ForceSub(vec1, vec2 []float64) {
	defer recover()
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

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
///func LongestDiffSlice(vec1, vec2 []float64) (long []float64)
