package vector

// Returns logical value of `leftInclusiveBound` > `rightInclusiveBound`
// where pointers are treated as their binary values.
func IsEmptyArray(leftInclusiveBound, rightInclusiveBound *float64) bool {
	return newUintptr(rightInclusiveBound) < newUintptr(leftInclusiveBound)
}

// Returns logical value of ⟨l1,r1⟩∩⟨l2,r2⟩=∅,
// where pointers are treated as their binary values.
//
// Warning: arrays may be moved in memory => function will give meaningless output.
// But this is hardly possible.
func AreArraysOverlapping(l1, r1, l2, r2 *float64) bool {
	// Check if there is an empty one.
	for _, l_r := range [2][2]*float64{{l1, r1}, {l2, r2}} {
		if IsEmptyArray(l_r[0], l_r[1]) {
			return false
		}
	}

	// Check the first range if in second
	for _, lOrR := range [2]*float64{l1, r1} {
		if ptrInRange(lOrR, l2, r2) {
			return true
		}
	}

	// Check the second range if in first
	for _, lOrR := range [2]*float64{l2, r2} {
		if ptrInRange(lOrR, l1, r1) {
			return true
		}
	}

	// Alternatively
	return false
}

// Returns true if given slices share some heap memory.
func AreSlicingSameArray(sl1, sl2 []float64) bool {
	// Activity check
	for _, sl := range [2][]float64{sl1, sl2} {
		if sl == nil || cap(sl) == 0 {
			return false
		}
	}

	// We extend slices for full memory check.
	sl1 = sl1[:cap(sl1)]
	sl2 = sl2[:cap(sl2)]

	//Knowing that both slices now have atleast one element.
	l1, r1 := &sl1[0], &sl1[len(sl1)-1]
	l2, r2 := &sl2[0], &sl2[len(sl2)-1]

	return AreArraysOverlapping(l1, r1, l2, r2)
}

// Fills []float64 with random (normal & abnormal) float values.
func FillWithRandF64s(slice []float64) {
	for i := range slice {
		slice[i] = randF64()
	}
}

// Returns new []float64 of given `length`,
// with random (normal & abnormal) float values.
func WithRandF64s(length uint) []float64 {
	slice := make([]float64, length)
	FillWithRandF64s(slice)
	return slice
}

// Fills []float64
// with different values obtained from rand.ExpFloat64()
// gives all values random sign (- or +).
func FillWithRandVals(slice []float64) {
	for i := range slice {
		slice[i] = randNormalF64()
	}
}

// Returns new []float64 of given `length`,
// with different values obtained from rand.ExpFloat64()
// gives all values random sign (- or +).
func WithRandVals(length uint) []float64 {
	slice := make([]float64, length)
	FillWithRandVals(slice)
	return slice
}

// Returns true if given slices are deeply equal,
// but do not share any heap memory.
func AreClones(v, u []float64) bool {
	return !AreSlicingSameArray(v, u) && DeepEqual(v, u)
}
