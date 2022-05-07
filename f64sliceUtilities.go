package vector

import (
	"reflect"
	"unsafe"
)

// Casts *float64 into uintptr
func f64PtrVal(val *float64) uintptr {
	return uintptr(unsafe.Pointer(val))
}

// Returns logical value of `val`∈⟨`leftInclusiveBound`,`rightInclusiveBound`⟩
func ptrInRange(val, leftInclusiveBound, rightInclusiveBound *float64) bool {
	v := f64PtrVal(val)
	return (f64PtrVal(leftInclusiveBound) <= v && v <= f64PtrVal(rightInclusiveBound))
}

// Returns logical value of ⟨l1,r1⟩∩⟨l2,r2⟩=∅,
// where pointers are treated as their binary values.
func AreArraysOverlapping(l1, r1, l2, r2 *float64) bool {
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

// Returns true if given slices are deeply equal,
// but do not share any heap memory.
func AreClones(v, u []float64) bool {
	return !AreSlicingSameArray(v, u) && reflect.DeepEqual(v, u)
}
