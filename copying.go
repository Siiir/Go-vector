package vector

/// Makes a deep copy (named `clone`) of a given []T.
func Clone[T interface{}](vec []T) (cl []T) {
	cl = make([]T, len(vec))
	copy(cl, vec)
	return
}
