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

	sum = Clone(vec1)
	ForceAdd(sum, vec2)
	return
}

func LongestSumSlice(vec1, vec2 []float64) (long []float64) {
	//! The longest slice is copied.
	//! We force addition of shorter slice to longer,
	//!  so that we archieve longer resulting vector,
	//! If we forced addition of larger to smaller,
	//!  the larger slice would have to be truncated (tail'd be ommited)
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

func Diff(vec1, vec2 []float64) (diff []float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	diff = Clone(vec1)
	ForceSub(diff, vec2)
	return
}

func LongestDiffSlice(vec1, vec2 []float64) (long []float64)
