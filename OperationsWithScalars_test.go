package vector

import (
	"math"
	"reflect"
	"testing"
)

func TestMulByScalar(t *testing.T) {
	vec := []float64{4, 0, 7, 7.9}

	const tQuantity = 3
	ScalarTab := [tQuantity]float64{1, 0.02, 0}
	ExpectTab := [tQuantity][]float64{
		{4, 0, 7, 7.9}, {0.08, 0, 0.14, 0.158}, {0, 0, 0, 0},
	}

	// Algorithm
	for i, scalar := range ScalarTab {
		MulByScalar(vec, scalar)
		if ex := ExpectTab[i]; !reflect.DeepEqual(ex, vec) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, vec)
		}
	}
	vec = append(vec, 5.0)
	MulByScalar(vec, 1e7)
	if ex := []float64{0, 0, 0, 0, 5e7}; !reflect.DeepEqual(ex, vec) {
		t.Errorf("\nLast test case failed. Expected != Got"+
			"\n%v != %v", ex, vec)
	}
}

func TestDivByScalar(t *testing.T) {
	vec := []float64{4, 0, 7, 7.9}

	const tQuantity = 3
	ScalarTab := [tQuantity]float64{1, 50, math.Inf(1)}
	ExpectTab := [tQuantity][]float64{
		{4, 0, 7, 7.9}, {0.08, 0, 0.14, 0.158}, {0, 0, 0, 0},
	}

	// Algorithm
	for i, scalar := range ScalarTab {
		DivByScalar(vec, scalar)
		if ex := ExpectTab[i]; !reflect.DeepEqual(ex, vec) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, vec)
		}
	}
	vec = append(vec, 4.0)
	DivByScalar(vec, 1e-7)
	if ex := []float64{0, 0, 0, 0, 4e7}; !reflect.DeepEqual(ex, vec) {
		t.Errorf("\nLast test case failed. Expected != Got"+
			"\n%v != %v", ex, vec)
	}
}

// func TestProdWithScalar(t *testing.T) ommited, because all subfunctions are tested.

// func TestQuotWithScalar(t *testing.T) ommited, because all subfunctions are tested.
