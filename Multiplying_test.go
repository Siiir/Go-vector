package vector

import (
	"reflect"
	"testing"
)

func TestForceDotProduct(t *testing.T) {
	// Definitions
	const tQuantity = 7
	ArgsTab := [tQuantity][2][]float64{
		{{}, {}},                       // Both empty.
		{{1, 90, 30}, {7, 0.5, 2}},     // Same length.
		{{200000, 200_000}, {0, 3, 0}}, // Apparently different length.
		{{1, 2, 3}, {4, 5}},            // Different length.
		{{4, 5}, {1, 2, 3}},            // Diff. len. reversed.
		{{}, {3, 5}},                   // One empty
		{{6e34, 9e+34}, {}},            // One empty reversed.
	}
	ExpectTab := [tQuantity]float64{
		0,
		7 + 45 + 60,
		600_000,
		4 + 10 + 0,
		4 + 10 + 0,
		0,
		0,
	}

	// Algorithm
	for i, Args := range ArgsTab {
		sl1, sl2 := Args[0], Args[1]
		if ex, got := ExpectTab[i], ForceDotProd(sl1, sl2); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestDotProduct(t *testing.T) {
	// Definitions
	const tQuantity = 7
	ArgsTab := [tQuantity][2][]float64{
		{{}, {}},                       // Both empty.
		{{1, 90, 30}, {7, 0.5, 2}},     // Same length.
		{{200000, 200_000}, {0, 3, 0}}, // Apparently different length.
		{{1, 2, 3}, {4, 5}},            // Different length.
		{{4, 5}, {1, 2, 3}},            // Diff. len. reversed.
		{{}, {3, 5}},                   // One empty
		{{6e34, 9e+34}, {}},            // One empty reversed.
	}
	ExpectTab := [tQuantity]struct {
		float64
		bool
	}{
		{0, false},
		{7 + 45 + 60, false},
		{0, true},
		{0, true},
		{0, true},
		{0, true},
		{0, true},
	}

	// Algorithm
	for i, Args := range ArgsTab {
		sl1, sl2 := Args[0], Args[1]
		got1, got2 := DotProd(sl1, sl2)
		if ex := ExpectTab[i]; !(ex.float64 == got1 && ex.bool == got2) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != {%v %v}", i, ex, got1, got2)
		}
	}
}

// Universal definitions.
const CrossProdTQuantity = 9

var CrossProdArgsTab = [CrossProdTQuantity][2][]float64{
	// Correct args.
	{{1, 2, 3}, {1, 5, 7}},                                   // Same length.
	{{-1, -2, 3}, {4, 0, -8}},                                // Same len. One 0.
	{{-1, -2, 3}, {4, 0, -8, 0, 0}},                          // Apparently different length.
	{{1, 2, 3, 345, 45, 94, 0}, {4.7, 5, 7e-2, 9e34, 834.4}}, // Different length.
	// Incorrect args.
	{},                   // Both empty.
	{{1, 2, 3}, {4, 5}},  // Different length.
	{{4, 5}, {1, 2, 3}},  // Diff. len. reversed.
	{{}, {3, 5}},         // One empty
	{{6e34, 9e+34}, nil}, // One nil
}

func TestUncheckedCrossProd(t *testing.T) {
	// Definitions
	ExpectTab := [CrossProdTQuantity][]float64{
		{-1, -4, 3},
		{16, 4, 8},
		{16, 4, 8},
		{-14.86, 14.030000000000001, -4.4},
	}

	// Algorithm
	for i, Args := range CrossProdArgsTab {
		sl1, sl2 := Args[0], Args[1]
		ex := ExpectTab[i]
		if ex != nil {
			if got := UncheckedCrossProd(sl1, sl2); !reflect.DeepEqual(ex, got) {
				t.Errorf("\nTest case %d failed. Expected != Got"+
					"\n%v != %v", i, ex, got)
			}
		} else {
			func() {
				defer callRecover()
				UncheckedCrossProd(sl1, sl2)
				t.Errorf("\nTest case %d failed. Should have paniced!", i)
			}()
		}

	}
}

func TestNilableCrossProduct(t *testing.T) {
	// Definitions
	ExpectTab := [CrossProdTQuantity][]float64{
		{-1, -4, 3},
		{16, 4, 8},
	}

	// Algorithm
	for i, Args := range CrossProdArgsTab {
		sl1, sl2 := Args[0], Args[1]
		ex := ExpectTab[i]
		if got := NilableCrossProd(sl1, sl2); !reflect.DeepEqual(ex, got) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestCrossProd(t *testing.T) {
	// Definitions
	ExpectTab := [CrossProdTQuantity][]float64{
		{-1, -4, 3},
		{16, 4, 8},
	}

	// Algorithm
	for i, Args := range CrossProdArgsTab {
		sl1, sl2 := Args[0], Args[1]
		ex1 := ExpectTab[i]
		ex2 := ex1 == nil
		if got1, got2 := CrossProd(sl1, sl2); !(reflect.DeepEqual(ex1, got1) && ex2 == got2) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n(%v, %v) != (%v, %v)", i, ex1, ex2, got1, got2)
		}
	}
}
