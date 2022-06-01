package vector

/*Returns logical value of sentence:
`slA`[i]=`slB`[i] for every i < min(len(`slA`), len(`slB`))
*/
func WeakEqual[T comparable](slA, slB []T) (equal bool) {
	equal = true
	defer callRecover()
	for i := range slA {
		if slA[i] != slB[i] {
			return false
		}
	}
	return
}

// Equivalent to: _WeakRelatedBy(vecA, vecB, f64ValEquivalence)
func _WeakEquivalent(vecA, vecB []float64) bool {
	return _WeakRelatedBy(vecA, vecB, f64ValEquivalence)
}

/*
Works like `reflect.DeepEqual`, but faster.
Returns true if one of following is true:
	1. both slices are nil
	2. both slices have same length and WeakEqual(slA, slB)==true
otherwise false.
*/
func DeepEqual[T comparable](slA, slB []T) (equal bool) {
	if b1, b2 := slA == nil, slB == nil; b1 || b2 {
		return b1 && b2
	}
	return len(slA) == len(slB) && WeakEqual(slA, slB)
}

// Equivalent to: _DeepRelatedBy(vecA, vecB, f64ValEquivalence)
func _DeepEquivalent(vecA, vecB []float64) bool {
	return _DeepRelatedBy(vecA, vecB, f64ValEquivalence)
}

/*
Returns true if
	both arrays described by slices own the same values (samely ordered) on the heap
	and len(`slA`)==len(`slB`).
*/
func HeapEqual[T comparable](slA, slB []T) (equal bool) {
	return cap(slA) == cap(slB) && DeepEqual(slA, slB)
}

// Equivalent to: _HeapRelatedBy(vecA, vecB, f64ValEquivalence)
func _HeapEquivalent(vecA, vecB []float64) bool {
	return _HeapRelatedBy(vecA, vecB, f64ValEquivalence)
}

/*
Returns true if both slices satisfy any of these conditions:
	1. are nil
	2. slice an empty array
	3. are slicing the same part of the same array
*/
func HeapSame[T any](slA, slB []T) (equal bool) {
	// Checking 2 integer components.
	if cap(slA) != cap(slB) || len(slA) != len(slB) {
		return false
	}
	// Checking pointer
	if cap(slA) == 0 {
		if b1, b2 := slA == nil, slB == nil; b1 || b2 {
			return b1 && b2
		}
		return true
	}
	return &slA[:1][0] == &slB[:1][0]
}
