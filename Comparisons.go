package vector

/*Returns logical value of sentence:
`sl1`[i]=`sl2`[i] for every i < min(len(`sl1`), len(`sl2`))
*/
func WeakEqual[T comparable](sl1, sl2 []T) (equal bool) {
	equal = true
	defer callRecover()
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return
}

/*!
Works like `reflect.DeepEqual`, but faster. No benchmark yet!
Returns true if one of following is true:
	1. both slices are nil
	2. both slices have same length and WeakEqual(sl1, sl2)==true
otherwise false.
*/
func DeepEqual[T comparable](sl1, sl2 []T) (equal bool) {
	if b1, b2 := sl1 == nil, sl2 == nil; b1 || b2 {
		return b1 && b2
	}
	return len(sl1) == len(sl2) && WeakEqual(sl1, sl2)
}

/*
Returns true if
	both arrays own the same values (samely ordered) on the heap
	and len(`sl1`)==len(`sl2`).
*/
func HeapEqual[T comparable](sl1, sl2 []T) (equal bool) {
	return cap(sl1) == cap(sl2) && DeepEqual(sl1, sl2)
}

/*
Returns true if both slices satisfy any of these conditions:
	1. are nil
	2. don't slice any array at all and are not nil
	3. are slicing the same part of the same array
*/
func HeapSame[T any](sl1, sl2 []T) (equal bool) {
	// Checking 2 integer components.
	if cap(sl1) != cap(sl2) || len(sl1) != len(sl2) {
		return false
	}
	// Checking pointer
	if cap(sl1) == 0 {
		if b1, b2 := sl1 == nil, sl2 == nil; b1 || b2 {
			return b1 && b2
		}
		return true
	}
	return &sl1[:1][0] == &sl2[:1][0]
}
