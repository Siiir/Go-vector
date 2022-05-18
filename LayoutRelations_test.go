package vector

import "testing"

func Test_OrthogonalityCheckingFuncs(t *testing.T) {
	// Definitions
	const tcQuantity = 8
	ArgTab := [tcQuantity][2][]float64{
		// Same len
		{{}, {}},
		{{1, 0, 0}, {0, 0, 7e200}},
		{{7e5, 0.009}, {-0.009, 7e5}},
		// Apparently different len.
		{{1, 2, 3}, {-12, 3, 2, 0, 0}},
		// Different len.
		{{1, -3}, {6, 2, 4, 787.45}},
		// Not orthogonal at all
		{{5, 6}, {9, 0}},
		{{2, 90, 1}, {2, 90, 1}},
		{{1, 2, 3}, {0, 4}},
	}
	testCases := func(f func([]float64, []float64, float64) bool, ExpectTab []bool) {
		for i, args := range ArgTab {
			ex := ExpectTab[i]
			got := f(args[0], args[1], 0)
			if ex != got {
				t.Errorf("\nTest case %d failed! expect != got\n%v != %v\n",
					i, ex, got)
			}
		}
	}
	// Subtest 1.
	t.Run("AreOrthogonalAfterDimAlignment", func(t *testing.T) {
		ExpectTab := [tcQuantity]bool{true, true, true, true, true}
		testCases(AreOrthogonalAfterDimAlignment, ExpectTab[:])
	})
	// Subtest 2.
	t.Run("AreAbsolutelyOrthogonal", func(t *testing.T) {
		ExpectTab := [tcQuantity]bool{true, true, true}
		testCases(AreAbsolutelyOrthogonal, ExpectTab[:])
	})
}

func Test_ParallelityCheckingFuncs(t *testing.T) {
	// Definitions
	const tcQuantity = 9
	ArgTab := [tcQuantity][2][]float64{
		// Same len
		{{}, {}},
		{{1, 0, 0}, {7e200, 0, 0}},
		{{-7e5, 0.009}, {7e5, -0.009}},
		{{6e-200, 1.8e-300}, {6e+300, 1.8e+200}}, // ! ratio=~0
		// Apparently different len.
		{{1, 2, 3}, {1, 2, 3, 0, 0}},
		// Different len. => not parallel
		{{1}, {1, 4, 787.45}},
		// Not parallel at all
		{{2, 90, 1}, {2, 90, 1, 1}},
		{{0, 9}, {9, 0}},
		{{1, 2, 3}, {0, 4}},
	}
	testCases := func(f func([]float64, []float64, float64) bool, ExpectTab []bool) {
		for i, args := range ArgTab {
			ex := ExpectTab[i]
			got := f(args[0], args[1], 0)
			if ex != got {
				t.Errorf("\nTest case %d failed! expect != got\n%v != %v\n",
					i, ex, got)
			}
		}
	}
	// Subtest 1.
	t.Run("AreParallelAfterDimAlignment", func(t *testing.T) {
		ExpectTab := [tcQuantity]bool{true, true, true, true, true}
		testCases(AreParallelAfterDimAlignment, ExpectTab[:])
	})
	// Subtest 2.
	t.Run("AreAbsolutelyParallel", func(t *testing.T) {
		ExpectTab := [tcQuantity]bool{true, true, true, true}
		testCases(AreAbsolutelyParallel, ExpectTab[:])
	})
}
