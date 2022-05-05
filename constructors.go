package vector

/*Constructs a new []float64 (named `vec`) with length = `length`
filled with given `scalar`.
*/
func NewFilled(length uint, scalar float64) (vec []float64) {
	vec = make([]float64, length)
	for i := range vec {
		vec[i] = scalar
	}
	return
}
