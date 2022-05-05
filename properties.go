package vector

import "math"

func Sq(vec []float64) (sq float64) {
	for _, v := range vec {
		sq += v * v
	}
	return
}

func CalcNorm(vec []float64) (norm float64) {
	return math.Sqrt(Sq(vec))
}

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
