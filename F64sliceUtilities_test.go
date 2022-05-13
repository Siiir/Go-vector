package vector

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/huandu/go-clone"
)

func TestIsEmptyArray(t *testing.T) {
	// Definitions
	const tQuantity = 3
	ArgsTab := [tQuantity][2]uintptr{
		{90, 20},
		{200000, 200000},
		{200000, 10000000},
	}
	ExpectTab := [tQuantity]bool{true, false, false}

	// Algorithm
	for i := 0; i < tQuantity; i++ {
		// Converted args go to `a` array
		var a [2]*float64
		for j, uPtr := range ArgsTab[i] {
			a[j] = newFloat64Ptr(uPtr)
		}

		//Now we can run our function
		if ex, got := ExpectTab[i], IsEmptyArray(a[0], a[1]); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestAreArraysOverlapping(t *testing.T) {
	// Definitions
	const tQuantity = 5
	ArgsTab := [tQuantity][4]uintptr{
		{90, 20, 80, 30},                 // Both empty
		{200000, 200000, 200000, 200000}, // Small but identic
		{0, 10000000, 100, 500},          // Second in first
		{100, 500, 0, 10000000},          // First in second
		{3, 1, 0, 10000},                 // First empty
	}
	ExpectTab := [tQuantity]bool{false, true, true, true, false}

	// Algorithm
	for i := 0; i < tQuantity; i++ {
		// Converted args go to `a` array
		var a [4]*float64
		for j, uPtr := range ArgsTab[i] {
			a[j] = newFloat64Ptr(uPtr)
		}

		//Now we can run our function
		if ex, got := ExpectTab[i], AreArraysOverlapping(a[0], a[1], a[2], a[3]); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestAreSlicingSameArray(t *testing.T) {
	const tQuantity = 6
	ArgsTab := [tQuantity][2][]float64{
		{{}, {}},                             //0 Both empty
		{{200000, 200000}, {200000, 200000}}, //1 Small but identic
		{{-4, 8, -20.3434}},                  //2 Cloned
		{{100, 500, 0, 10000000}},            //3 Copied
		{{3, 1, 0, 10000}},                   //4 Containing
		{{-0, 7.4, 5e5, 0}},                  //5 Splitting <=> Apparently disjunctive
	}

	ArgsTab[2][1] = Clone(ArgsTab[2][0])
	ArgsTab[3][1] = ArgsTab[3][0]
	ArgsTab[4][1] = ArgsTab[4][0][:2]
	{
		ArgsTab[5][1] = ArgsTab[5][0][2:]
		ArgsTab[5][0] = ArgsTab[5][0][0:2]
	}
	ExpectTab := [tQuantity]bool{false, false, false, true, true, true}

	for i, Args := range ArgsTab {
		sl1, sl2 := Args[0], Args[1]
		if ex, got := ExpectTab[i], AreSlicingSameArray(sl1, sl2); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func TestAreClones(t *testing.T) {
	const tQuantity = 10
	ArgsTab := [tQuantity][2][]float64{
		// Prefabricated part of ArgsTab
		{{}, {}},                             //0 Both empty
		{{200000, 200000}, {200000, 200000}}, //1 Small but identic

		// Partialy generated part of ArgsTab
		{{-4, 8, -20.3434}},             //2 Cloned
		{{100, 500, 0, 10000000, 89.0}}, //3 Cloned with github.com/huandu/go-clone.clone
		{{100, 500, 0, 10000000}},       //4 Copied
		{{3, 1, 0, 10000}},              //5 Containing
		{{-0, 7.4, 5e5, 0}},             //6 Splitting <=> Apparently disjunctive

		// Prefabricated part of ArgsTab
		{{5, 6}, {0, -66e-80}},           //7 Same len, but different.
		{{}, {6, 73, 8}},                 //8 one empty
		{{5, 6}, {0, -66e-80, 6, 8, 90}}, //9 much different

	}

	ArgsTab[2][1] = Clone(ArgsTab[2][0])
	ArgsTab[3][1] = clone.Clone(ArgsTab[3][0]).([]float64)
	ArgsTab[4][1] = ArgsTab[4][0]
	ArgsTab[5][1] = ArgsTab[5][0][:2]
	{
		ArgsTab[6][1] = ArgsTab[6][0][2:]
		ArgsTab[6][0] = ArgsTab[6][0][0:2]
	}
	ExpectTab := [tQuantity]bool{true, true, true, true} //false,false...

	for i, Args := range ArgsTab {
		sl1, sl2 := Args[0], Args[1]
		if ex, got := ExpectTab[i], AreClones(sl1, sl2); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

// Omitted ```func ExampleFillWithRand64s()```, because contained in ```WithRand64s()```

func ExampleWithRandF64s() {
	for i := 10; i > 0; i-- {
		sl := WithRandF64s(uint(rand.Intn(7)))
		fmt.Printf("%v\n", sl)
	}
	//:-) Output:
}

// Omitted ```func ExampleFillWithRandVal()```, because contained in ```WithRandVal()```

func ExampleWithRandVals() {
	for i := 10; i > 0; i-- {
		sl := WithRandVals(uint(rand.Intn(7)))
		fmt.Printf("%v\n", sl)
	}
	//:-) Output:

}

func Benchmark_CloneAlgorithms_N_in0to1400(b *testing.B) {
	const bQuantity = uint(14_000)
	var slices [bQuantity][]float64
	// Setup
	for u := bQuantity; 0 < u; {
		u-- //`u` will reach 0 in last iteration
		slices[u] = WithRandF64s(u)
	}

	//0.1282 ns/op
	b.Run(
		"clone.Clone",
		func(_ *testing.B) {
			for u := bQuantity; 0 < u; {
				u-- //`u` will reach 0 in last iteration
				_ = clone.Clone(slices[u]).([]float64)
			}
		},
	)

	// 0.1199 ns/op
	b.Run(
		"vector.Clone",
		func(_ *testing.B) {
			for u := bQuantity; 0 < u; {
				u-- //`u` will reach 0 in last iteration
				_ = Clone(slices[u])
			}
		},
	)

	//0.1295 ns/op
	b.Run(
		"clone.Slowly",
		func(_ *testing.B) {
			for u := bQuantity; 0 < u; {
				u-- //`u` will reach 0 in last iteration
				_ = clone.Slowly(slices[u]).([]float64)
			}
		},
	)
}

func Benchmark_CloneAlgorithms_Neq10_000_000(b *testing.B) {
	const N = uint(10_000_000)
	var slice []float64 = WithRandF64s(N)

	//0.01502 ns/op
	b.Run(
		"clone.Clone",
		func(_ *testing.B) {
			_ = clone.Clone(slice).([]float64)
		},
	)
	//0.004998 ns/op
	b.Run(
		"vector.Clone",
		func(_ *testing.B) {
			_ = Clone(slice)
		},
	)
	//0.01599 ns/op
	b.Run(
		"clone.Slowly",
		func(_ *testing.B) {
			_ = clone.Slowly(slice).([]float64)
		},
	)
}
