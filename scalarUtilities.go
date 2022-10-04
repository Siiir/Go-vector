package vector

import (
	"C"
)

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
