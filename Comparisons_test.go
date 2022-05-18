package vector

import (
	"reflect"
	"testing"
)

const eqFuncsTcQuantity = 17

func CreateEqFuncsArgTab() [eqFuncsTcQuantity][2][]float64 {
	randSlice := WithRandVals(9)

	partialClone := Clone(randSlice[1:3])
	randSliceClone := Clone(randSlice)

	concreteSlice1, concreteSlice2 := []float64{1, 2, 3, 8.6}, []float64{1, 2, 3, 4, 5}
	argTab := [eqFuncsTcQuantity][2][]float64{
		{nil, nil},                                     // HeapSame 1.
		{{}, {}},                                       // HeapSame 2.
		{randSlice, randSlice},                         // HeapSame 3.
		{randSlice[3:6], randSlice[3:6]},               // HeapSame 3.
		{{1, -2}, {1, -2}},                             // HeapEqual
		{{8, 0e4, -7.8}, {8e0, 0.0, -7.8}},             // HeapEqual
		{randSlice, randSliceClone},                    // HeapEqual
		{randSlice[1:3], partialClone},                 // DeepEqual
		{concreteSlice1[:3], concreteSlice2[:3]},       // DeepEqual
		{randSlice, randSlice[:6]},                     // WeakEqual
		{concreteSlice1, concreteSlice2[:3]},           // WeakEqual
		{randSlice, append(randSlice, 1.8, 8934e-300)}, // WeakEqual
		{{}, nil},                            // WeakEqual
		{{}, {-1, 2}},                        // WeakEqual
		{randSlice, nil},                     // WeakEqual
		{concreteSlice1, concreteSlice1[1:]}, // Not equal
		{{1, 4}, {1, 7}},                     // Not equal
	}

	return argTab
}

var eqFuncsArgTab = CreateEqFuncsArgTab()

func reflect_DeepEqualOfF64SliceWrapper(vec1, vec2 []float64) bool {
	return reflect.DeepEqual(vec1, vec2)
}

func Test_EqualityFunctions(t *testing.T) {
	testCases := func(t *testing.T, f func(vec1, vec2 []float64) bool, expectSlice []bool) {
		for i, args := range eqFuncsArgTab {
			ex := expectSlice[i]
			got := f(args[0], args[1])
			if ex != got {
				t.Errorf("\n%s: case %d failed! expected!=got\n%v != %v\n",
					t.Name(), i, ex, got)
			}
		}
	}

	t.Run("HeapSame[float64]", func(t *testing.T) { //-=> HeapSame
		expectTab := [eqFuncsTcQuantity]bool{true, true, true, true}
		testCases(t, HeapSame[float64], expectTab[:])
	})
	t.Run("HeapEqual", func(t *testing.T) {
		expectTab := [eqFuncsTcQuantity]bool{true, true, true, true, true, true, true}
		testCases(t, HeapEqual[float64], expectTab[:])

	})
	t.Run("DeapEqual", func(t *testing.T) {
		expectTab := [eqFuncsTcQuantity]bool{
			true, true, true, true, true,
			true, true, true, true,
		}

		t.Run("DeepEqual[float64]",
			func(t *testing.T) { testCases(t, DeepEqual[float64], expectTab[:]) })
		t.Run("reflect.DeepEqual",
			func(t *testing.T) { testCases(t, reflect_DeepEqualOfF64SliceWrapper, expectTab[:]) },
		)
	})
	t.Run("WeakEqual[float64]", func(t *testing.T) {
		expectTab := [eqFuncsTcQuantity]bool{
			true, true, true, true, true,
			true, true, true, true, true,
			true, true, true, true, true,
		}
		testCases(t, WeakEqual[float64], expectTab[:])
	})
}

func createRelationFuncLauncher(
	relFunc func(sl1, sl2 []float64) bool,
	argSlice [][2][]float64,
	launchesAmount uint,
) (funcLauncher func(*testing.B)) {
	return func(_ *testing.B) {
		for ; launchesAmount > 0; launchesAmount-- {
			for _, argsPair := range argSlice {
				relFunc(argsPair[0], argsPair[1])
			}
		}
	}
}

func Benchmark_DeepEqualFuncs(b *testing.B) {
	b.Run("eqFuncsArgTab", func(b *testing.B) {
		const launchesAmount = 5000000
		b.Run(
			"DeepEqual[float64]",
			createRelationFuncLauncher(DeepEqual[float64], eqFuncsArgTab[:], launchesAmount),
		)
		b.Run(
			"reflect_DeepEqualOfF64SliceWrapper",
			createRelationFuncLauncher(reflect_DeepEqualOfF64SliceWrapper, eqFuncsArgTab[:], launchesAmount),
		)
		b.Run(
			"reflect.DeepEqual",
			func(_ *testing.B) {
				lA := launchesAmount
				for ; lA > 0; lA-- {
					for _, argsPair := range eqFuncsArgTab {
						reflect.DeepEqual(argsPair[0], argsPair[1])
					}
				}
			},
		)
	})

	b.Run("HeavyArgs 0,1,0,1,...", func(b *testing.B) {
		// Definitions
		const argQuantity = uint(10)
		const slicesLen = uint(1000000)
		var firstArgTab, secondArgTab [argQuantity][]float64
		// Setup
		for aQ := argQuantity; aQ > 0; {
			aQ--
			firstArgTab[aQ] = WithRandVals(slicesLen)
			if aQ&1 == 1 {
				secondArgTab[aQ] = WithRandVals(slicesLen)
			} else {
				secondArgTab[aQ] = Clone(firstArgTab[aQ])
			}
		}
		// Benchmarks
		b.Run("DeepEqual[float64]", func(_ *testing.B) {
			for aQ := argQuantity; aQ > 0; {
				aQ--
				DeepEqual(firstArgTab[aQ], secondArgTab[aQ])
			}
		})
		b.Run("reflect.DeepEqual", func(_ *testing.B) {
			for aQ := argQuantity; aQ > 0; {
				aQ--
				reflect.DeepEqual(firstArgTab[aQ], secondArgTab[aQ])
			}
		})
	})

}
