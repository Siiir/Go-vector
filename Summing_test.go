package vector

import (
	"reflect"
	"testing"
)

func Test_AddingFunctions(t *testing.T) {
	const tQuantity = 7

	var ArgTab = [tQuantity][]float64{
		{0, 5, 7},           // len. = 3
		{-0, -500.78, 7},    // len. = 3
		{-0, 500.78, 7},     // len. = 3
		{1, 5, 7, 0},        // len. = apparent 4
		{-0, -500.78, 7, 9}, // len. = 4
		{0, 1},              // len. = 2
		{},                  // Empty
	}

	t.Run("ForceAdd", func(t *testing.T) {
		// Definitions
		vec := []float64{7, -9, 5.009}
		ExpectTab := [tQuantity][]float64{
			{7, -4, 12.009},
			{7, -504.78, 19.009},
			{7, -4, 26.009},
			{8, 1, 33.009},
			{8, -499.78, 40.009},
			{8, -498.78, 40.009},
			{8, -498.78, 40.009},
		}

		// Algorithm
		for i, vec2 := range ArgTab {
			ForceAdd(vec, vec2)
			ex := ExpectTab[i]
			if !reflect.DeepEqual(ex, vec) {
				t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
					i, ex, vec)
			}
		}
	})

	t.Run("Add", func(t *testing.T) {
		// Definitions
		vec := []float64{7, -9, 5.009}
		ExpectTab := [tQuantity][]float64{
			{7, -4, 12.009},
			{7, -504.78, 19.009},
			{7, -4, 26.009},
		}

		// Algorithm
		ex := Clone(vec)
		for i, vec2 := range ArgTab {
			if Add(vec, vec2) { //States that impossible.
				if ExpectTab[i] != nil { // Actually possible
					t.Fatalf("\nTest case %d failed!\n"+
						"Function should return false (and do perform addition), returned true.",
						i)
				}
				// If indeed impossible, than
				if !reflect.DeepEqual(ex, vec) { //Shouldn't have changed.
					t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
						i, ex, vec)
				}
			} else if ex = ExpectTab[i]; !reflect.DeepEqual(ex, vec) {
				if ex == nil { // Actually impossible
					t.Fatalf("\nTest case %d failed!\n"+
						"Function should return true (and not perform addition), returned false.",
						i)
				}
				t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
					i, ex, vec)
			}
		}
	})

	// t.Run("Sum",...) ommited, because main subfunction has been tested.

	t.Run("LongestSumSlice", func(t *testing.T) {
		// Definitions
		vec := []float64{1}
		ExpectTab := [tQuantity][]float64{
			{1, 5, 7},
			{1, -495.78, 14},
			{1, 5, 21},
			{2, 10, 28, 0},
			{2, -490.78, 35, 9},
			{2, -489.78, 35, 9},
			{2, -489.78, 35, 9},
		}

		// Algorithm
		for i, vec2 := range ArgTab {
			vec = LongestSumSlice(vec, vec2)
			ex := ExpectTab[i]
			if !reflect.DeepEqual(ex, vec) {
				t.Fatalf("\nTest case %d failed! expected != got\n%v != %v",
					i, ex, vec)
			}
		}
	})
}

func Test_SubtractingFunctions(t *testing.T) {
	const tQuantity = 7

	var ArgTab = [tQuantity][]float64{
		{0.0, -5.0, -7.0},         // len. = -3.0
		{0.0, 500.78, -7.0},       // len. = -3.0
		{0.0, -500.78, -7.0},      // len. = -3.0
		{-1.0, -5.0, -7.0, 0.0},   // len. = apparent -4.0
		{0.0, 500.78, -7.0, -9.0}, // len. = -4.0
		{0.0, -1.0},               // len. = -2.0
		{},                        // Empty
	}

	t.Run("ForceSub", func(t *testing.T) {
		// Definitions
		vec := []float64{7, -9, 5.009}
		ExpectTab := [tQuantity][]float64{
			{7, -4, 12.009},
			{7, -504.78, 19.009},
			{7, -4, 26.009},
			{8, 1, 33.009},
			{8, -499.78, 40.009},
			{8, -498.78, 40.009},
			{8, -498.78, 40.009},
		}

		// Algorithm
		for i, vec2 := range ArgTab {
			ForceSub(vec, vec2)
			ex := ExpectTab[i]
			if !DeepEqual(ex, vec) {
				t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
					i, ex, vec)
			}
		}
	})

	t.Run("Sub", func(t *testing.T) {
		// Definitions
		vec := []float64{7, -9, 5.009}
		ExpectTab := [tQuantity][]float64{
			{7, -4, 12.009},
			{7, -504.78, 19.009},
			{7, -4, 26.009},
		}

		// Algorithm
		ex := Clone(vec)
		for i, vec2 := range ArgTab {
			if Sub(vec, vec2) { //States that impossible.
				if ExpectTab[i] != nil { // Actually possible
					t.Fatalf("\nTest case %d failed!\n"+
						"Function should return false (and do perform addition), returned true.",
						i)
				}
				// If indeed impossible, than
				if !DeepEqual(ex, vec) { //Shouldn't have changed.
					t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
						i, ex, vec)
				}
			} else if ex = ExpectTab[i]; !DeepEqual(ex, vec) {
				if ex == nil { // Actually impossible
					t.Fatalf("\nTest case %d failed!\n"+
						"Function should return true (and not perform addition), returned false.",
						i)
				}
				t.Fatalf("\nTest case %d failed! expected != actual\n%v != %v",
					i, ex, vec)
			}
		}
	})

	// t.Run("Diff",...) ommited, because main subfunction has been tested.

	t.Run("LongestDiffSlice", func(t *testing.T) {
		// Definitions
		argTab := [tQuantity][2][]float64{
			{nil, nil},
			{{}, {}},
			{{}, {-5}},
			{{-9.7}, nil},
			{{1, 2, 3}, {1, -2, 3, 0}}, //4.
			{{4, 5}, {2, 3, 8, 7}},
			{{0, 1, 5e200}, {0, 20}}, //6.
		}
		expectTab := [tQuantity][]float64{
			{}, {}, {5}, {-9.7},
			{0, 4, 0, 0}, //4.
			{2, 2, -8, -7},
			{0, -19, 5e200}, //6.
		}

		// Algorithm
		for i, args := range argTab {
			got := LongestDiffSlice(args[0], args[1])
			ex := expectTab[i]
			if !DeepEqual(ex, got) {
				//:{println(len(got), cap(got))
				t.Fatalf("\nTest case %d failed! expected != got\n%v != %v",
					i, ex, got)
			}
		}
	})
}
