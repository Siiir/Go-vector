package vector

func NewFilled(l uint, scalar float64) (vec []float64) {
	vec = make([]float64, l)
	for i := range vec {
		vec[i] = scalar
	}
	return
}
