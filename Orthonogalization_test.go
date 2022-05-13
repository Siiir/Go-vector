package vector

import (
	"reflect"
	"testing"
)

const ProjTQuantity = 7

var ProjArgTab = [ProjTQuantity][2][]float64{
	{{}, {}},                        // Both empty.
	{{1, 90, 30}, {7, 0.5, 2}},      // Same length.
	{{200_000, 200_000}, {0, 3, 0}}, // Apparently different length.
	{{1, 2, 3}, {4, 5}},             // Different length.
	{{4, 5}, {1, 2, 3}},             // Diff. len. reversed.
	{{}, {3, 5}},                    // One empty
	{{6e34, 9e+34}, {}},             // One empty reversed.
}

func TestForceProj(t *testing.T) {
	// Definitions
	ExpectTab := [ProjTQuantity][]float64{
		{},
		{14.723004694835681, 1.051643192488263, 4.206572769953052},
		{0, 200000, 0},
		{1.3658536585365855, 1.707317073170732},
		{1, 2, 3},
		{0, 0},
		{},
	}

	// Algorithm
	for i, Args := range ProjArgTab {
		sl1, sl2 := Args[0], Args[1]
		ex := ExpectTab[i]
		if got := ForceProj(sl1, sl2, nil); !reflect.DeepEqual(ex, got) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestNilableProj(t *testing.T) {
	// Definitions
	ExpectTab := [ProjTQuantity][]float64{
		{},
		{14.723004694835681, 1.051643192488263, 4.206572769953052},
	}

	// Algorithm
	for i, Args := range ProjArgTab {
		sl1, sl2 := Args[0], Args[1]
		ex := ExpectTab[i]
		if got := NilableProj(sl1, sl2, nil); !reflect.DeepEqual(ex, got) {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

// func TestProj(t *testing.T) ommited, because all subfunctions have been tested.

func TestNilableOrthogonalizedWith(t *testing.T) {
	// Definitions
	const tQuantity = 9
	ArgTab := [tQuantity][][]float64{
		/*Correct arguments.*/
		{{1, 90, 30}, {1, 0, 0}, {0, 2, 0}},         // Same length & `targets` orthogonal. 3 vec. : 3 dim.
		{{4, 3, 5, 5}, {1, 1, 0, 0}, {-1, 1, 0, 0}}, // Same len. & `targets` orthon. 3 vec. : 4 dim.
		/*Incorrect arguments.*/
		{{}, {}, {}, {}},                             // All empty. 4 vec : 0 dim
		{{1, 3}, {3, 5}, {9, 0}, {0, 3}},             // Dependant	4 : 2
		{{200000, 200_000}, {0, 3, 0}, {1, 6, 0, 0}}, // Apparently different length. 3 : 2
		{{1, 2, 3}, {4, 5}},                          // Different length. 2 : 3
		{{4, 5}, {1, 2, 3}},                          // Diff. len. reversed. 2 : 3
		{{}, {3, 5}},                                 // One empty 2 : 2
		{{6e34, 9e+34}, {}},                          // One empty reversed. 2 : 2
	}

	PossibilityTab := [tQuantity]bool{true, true}
	// Algorithm
	for i, Args := range ArgTab {
		orthogonalized := NilableOrthogonalizedWith(Args[0], Args[1:]...)
		if PossibilityTab[i] {
			for _, other := range Args[1:] {
				if !AreOrthogonalAfterDimAlignment(orthogonalized, other, 0) {
					t.Errorf("\nTest case %d failed. %v is not orthogonal to %v.\nTheir dot product is %v .",
						i, orthogonalized, other, ForceDotProd(orthogonalized, other))
					break
				}
			}
		} else { // If impossible, we expect nil.
			if orthogonalized != nil {
				t.Errorf("\nTest case %d failed. Expected != Got"+
					"\nnil != %v", i, orthogonalized)
			}
		}

	}
}

// func TestOrthogonalizedWith(t *testing.T) ommited, since main subfunction has been tested.
