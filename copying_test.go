package vector

import (
	"testing"
)

func TestClone(t *testing.T) {
	vectors := [][]float64{{1, 8, -9.43}, {}, {1e43, 5, -9, 2, 23, 2}}
	for _, vec := range vectors {
		clone := Clone(vec)
		if !AreClones(vec, clone) {
			t.Errorf("%v, %v aren't clones.", vec, clone)
		}
	}
}
