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
func CrossProd(vec1, vec2 []float64) float64
