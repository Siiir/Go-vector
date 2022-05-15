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

func TestPow(t *testing.T) {
	// Definitions
	const tQuantity = 6
	ArgTab := [tQuantity]struct {
		vec []float64
		exp uint
	}{
		{[]float64{}, 9043233},
		{[]float64{1, 3, 45, 324e-5}, 1},
		{[]float64{0, 0, 0, 0}, 34598234},
		{[]float64{7, 0, 22, -8}, 2},
		{[]float64{7, 0, 22, -8}, 6},
		{[]float64{4, 5, 1}, 3},
	}
	ExpectTab := []interface{}{
		[]float64{},
		[]float64{1, 3, 45, 324e-5},
		0.0,
		49.0 + 0.0 + 484.0 + 64.0,
		math.Pow(49.0+0.0+484.0+64.0, 3),
		[]float64{42 * 4, 42 * 5, 42 * 1},
	}
	for i, args := range ArgTab {
		vec, exp := args.vec, args.exp
		ex := ExpectTab[i]
		got := Pow(vec, exp)
		if !reflect.DeepEqual(ex, got) {
			t.Errorf("Test case %d failed! expect != got\n%v != %v",
				i, ex, got)
		}
	}
}
