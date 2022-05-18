package vector

import "testing"

func genVecsWithRandVals(amount uintptr) (vecs [][]float64) {
	vecs = make([][]float64, amount)
	for i := range vecs {
		vecs[i] = WithRandVals(uint(i))
	}
	return
}

func TestNegate(t *testing.T) {
	const tcQuantity = 20
	vecs := genVecsWithRandVals(tcQuantity)
	for i, vec := range vecs {
		clone := Clone(vec)
		Negate(clone)
		MulByScalar(clone, -1)
		if !DeepEqual(vec, clone) {
			t.Errorf("Tc %d failed! expected!=got\n%v != %v",
				i, vec, clone)
		}
	}
}

func TestNegated(t *testing.T) {
	const tcQuantity = 20
	vecs := genVecsWithRandVals(tcQuantity)
	for i, vec := range vecs {
		ex := ProdWithScalar(vec, -1)
		got := Negated(vec)

		if !DeepEqual(ex, got) {
			t.Errorf("Tc %d failed! expected!=got\n%v != %v",
				i, ex, got)
		}
	}
}
