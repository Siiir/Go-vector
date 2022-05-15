package vector

import (
	"math/rand"
	"unsafe"
)

// Casts *T into uintptr
func newUintptr[T interface{}](val *T) uintptr {
	return uintptr(unsafe.Pointer(val))
}

// Casts uintptr into *float64
func newFloat64Ptr(uPtr uintptr) *float64 {
	return (*float64)(unsafe.Pointer(uPtr))
}

// Returns logical value of `val`∈⟨`leftInclusiveBound`,`rightInclusiveBound`⟩
func ptrInRange[T interface{}](val, leftInclusiveBound, rightInclusiveBound *T) bool {
	v := newUintptr(val)
	return (newUintptr(leftInclusiveBound) <= v && v <= newUintptr(rightInclusiveBound))
}

//Returns either `fl` or `-fl`
func giveRandSign(fl float64) float64 {
	if (rand.Int31() & 1) == 1 {
		fl *= -1
	}
	return fl
}

/* Returns completly random float64.
Might be abnormal or special.*/
func randF64() float64 {
	i64 := rand.Uint64()
	return *(*float64)(unsafe.Pointer(&i64))
}

/*Returns random, not-abnormal, not-special (not NaN, not Inf) `float64`*/
func randNormalF64() float64 {
	var u64 uint64
	for {
		u64 = rand.Uint64()
		if u64 == 0 {
			break
		}
		u16 := *(*uint16)(unsafe.Pointer(&u64))
		u16 = (u16 << 1) >> 5
		if u16 != 0 && u16 != 0b111_1111_1111 {
			break
		}
	}
	return *(*float64)(unsafe.Pointer(&u64))
}

func minInt(i1, i2 int) (theSmaller int) {
	if i1 < i2 {
		return i1
	}
	return i2
}

func maxInt(i1, i2 int) (theLarger int) {
	if i1 < i2 {
		return i2
	}
	return i1
}
