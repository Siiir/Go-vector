package vector

import "testing"

func TestSq(t *testing.T) {
	// Definitions
	const tQuantity = 5
	randVec := WithRandVals(87)
	ArgTab := [tQuantity][]float64{
		{},
		{0, 0, 0},
		{8, 9},
		{5.7},
		randVec,
	}

	ExpectTab := [tQuantity]float64{
		0, 0, 64 + 81, 5.7 * 5.7,
		ForceDotProd(randVec, randVec)}

	// Algorithm
	for i, vec := range ArgTab {
		ex := ExpectTab[i]
		got := Sq(vec)
		if ex != got {
			t.Errorf("Test case %d failed! expect != got\n%v != %v",
				i, ex, got)
		}
	}
}

func TestCalcNorm(t *testing.T) {
	// Definitions
	const tQuantity = 4
	ArgTab := [tQuantity][]float64{
		{},
		{0, 0, 0},
		{4, 3},
		{-5.7},
	}

	ExpectTab := [tQuantity]float64{0, 0, 5, 5.7}

	// Algorithm
	for i, vec := range ArgTab {
		ex := ExpectTab[i]
		got := CalcNorm(vec)
		if ex != got {
			t.Errorf("Test case %d failed! expect != got\n%v != %v",
				i, ex, got)
		}
	}
}

func TestNumOfTrailingZeros(t *testing.T) {
	// Definitions
	const tQuantity = 4
	ArgTab := [tQuantity][]float64{
		{},
		{0, 0, 0},
		{4, 3, 0e-100, 0.0},
		{-5.7},
	}

	ExpectTab := [tQuantity]uint{0, 3, 2, 0}

	// Algorithm
	for i, vec := range ArgTab {
		ex := ExpectTab[i]
		got := NumOfTrailingZeros(vec)
		if ex != got {
			t.Errorf("Test case %d failed! expect != got\n%v != %v",
				i, ex, got)
		}
	}
}
