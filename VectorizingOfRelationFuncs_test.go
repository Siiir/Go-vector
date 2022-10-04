package vector

import (
	"math"
	"testing"
)

// Aux. functions

func haveSmallerCords(vecA, vecB []float64) bool {
	return _DeepRelatedBy(vecA, vecB,
		func(a, b float64) bool { return a < b })
}

func areSamellyMemorized(vecA, vecB []float64) bool {
	return _HeapRelatedBy(vecA, vecB,
		func(a, b float64) bool { return cmp64Bits(&a, &b) })
}

func haveGECords(vecA, vecB []float64) bool {
	return _DeepRelatedBy(vecA, vecB,
		func(a, b float64) bool { return a >= b })
}

func deepAbsEqual(vecA, vecB []float64) bool {
	return _DeepRelatedBy(vecA, vecB,
		func(a, b float64) bool { return math.Abs(a) == math.Abs(b) })
}

// Testing functions

// ! Not run.
func Test_VectorizingFuncs(t *testing.T) {
	const tcQuantity = 10

	var allFuncs = [4]func(vecA, vecB []float64) bool{
		haveSmallerCords,
		areSamellyMemorized,
		haveGECords,
		deepAbsEqual,
	}
	var allFuncsNames = [4]string{
		"haveSmallerCords",
		"areSamellyMemorized",
		"haveGECords",
		"deepAbsEqual",
	}

	argTab := [tcQuantity]struct {
		vecA, vecB []float64
		relation   func(vecA, vecB []float64) bool
	}{
		// 1.
		{[]float64{0, 1, 2, 3}, []float64{1, 2, 3, 4}, haveSmallerCords},
		{[]float64{5, 0.345, 65, 2, math.NaN(), 65}, []float64{5, 0.345, 65, 2, math.NaN(), 65}, areSamellyMemorized},
		{[]float64{0, 1, 2}, []float64{-1, -4e8, 2.0}, haveGECords},
		{[]float64{7, 5, -math.E}, []float64{7, -5, math.E}, deepAbsEqual},
		{[]float64{5, 7, 4}, []float64{6, 8, -9}, nil},
		// 2.
		{[]float64{-4, 90, math.Pi}, append([]float64{0, 800, 4}, 5, 6)[:3], haveSmallerCords},
		{append([]float64{}, 8, 9, math.NaN()), append([]float64{}, 8, 9, math.NaN()), areSamellyMemorized},
		{append([]float64{3, 4, 5}, 0), []float64{2, 4, 5, -700e300}, haveGECords},
		{[]float64{}, nil, deepAbsEqual},
		{[]float64{1}, []float64{}, nil},
	}

	// The main algortihm
	for tcInd, args := range argTab {
		for fInd, f := range allFuncs {
			ex := &f == &args.relation
			got := f(args.vecA, args.vecB)
			if ex != got {
				t.Fatalf("\nTC %d, %d.f.: \"%s\" failed! ex!=got\n\t%v != %v\n",
					tcInd, fInd, allFuncsNames[fInd], ex, got)
			}
		}
	}
}
