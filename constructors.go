package vector

/*Constructs a new []T (named `vec`) with length = `length`
filled with given `scalar`.
*/
func NewFilled[T interface{}](length uint, scalar T) (vec []T) {
	vec = make([]T, length)
	for i := range vec {
		vec[i] = scalar
	}
	return
}
