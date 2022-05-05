package vector

import "math"

/// Return standard dot product of vector `vec` and its copy.
func Sq(vec []float64) (sumOfSq float64) {
	for _, v := range vec {
		sumOfSq += v * v
	}
	return
}

/// Returns standard norm aka length of vector.
func CalcNorm(vec []float64) (norm float64) {
	return math.Sqrt(Sq(vec))
}

/// Returns the number of trailing zeros in given vector.
func NumOfTrailingZeros(vec []float64) (quantity uint) {
	defer recover()
	for i := len(vec) - 1; ; i-- {
		if vec[i] == 0 {
			quantity++
		} else {
			return
		}
	}
}
