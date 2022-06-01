package vector

/*Returns logical value of sentence:
  `ralation`(`slA`[i],`slB`[i]) for every i < min(len(`slA`), len(`slB`))
*/
func _WeakRelatedBy[T any](
	slA, slB []T,
	relation func(elA, elB T) bool,
) (equal bool) {
	equal = true
	defer callRecover()
	for i := range slA {
		if relation(slA[i], slB[i]) {
			return false
		}
	}
	return
}

/*
Returns true if one of following is true:
	1. both slices are nil
	2. both slices have same length and WeakRelatedBy(slA, slB, relation)==true
otherwise false.
*/
func _DeepRelatedBy[T any](slA, slB []T, relation func(elA, elB T) bool) (equal bool) {
	if b1, b2 := slA == nil, slB == nil; b1 || b2 {
		return b1 && b2
	}
	return len(slA) == len(slB) && _WeakRelatedBy(slA, slB, relation)
}

// Equivalet to:  cap(slA) == cap(slB) && _DeepRelatedBy(slA, slB, relation)
func _HeapRelatedBy[T comparable](slA, slB []T, relation func(elA, elB T) bool) (equal bool) {
	return cap(slA) == cap(slB) && _DeepRelatedBy(slA, slB, relation)
}

// `func HeapSameBy[T any](slA, slB []T) (equal bool)` doesn't exist,
// because identity of heap memory has only one definition.
