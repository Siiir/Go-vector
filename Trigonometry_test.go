package vector

import (
	"fmt"
	"math"
	"testing"
)

type func2VecsToF64 func(vec1, vec2 []float64) float64
type func2VecsToF64AndBool func(vec1, vec2 []float64) (float64, bool)

func testFunc2VecsToF64(t *testing.T, f func2VecsToF64, argSlice [][2][]float64, expectSlice []float64) {
	if l1, l2 := len(argSlice), len(expectSlice); l1 != l2 {
		panic(
			"\nWrong arguments to function `testFunc2VecsToF64`." +
				"\nlen(argSlice)!=len(expectSlice)" +
				fmt.Sprintf("\n%d != %d\n", l1, l2),
		)
	}
	for i, args := range argSlice {
		ex := expectSlice[i]
		got := f(args[0], args[1])
		if !F64ValEquivalence(ex, got) {
			t.Fatalf("\n%s %d. tc failed! !F64ValEquivalence(ex, got)\nex=%v\ngot=%v\n",
				t.Name(), i, ex, got,
			)
		}
	}
}

type f64AndBool struct {
	float64
	bool
}

func testFunc2VecsToF64AndBool(t *testing.T, f func2VecsToF64AndBool, argSlice [][2][]float64, expectSlice []f64AndBool) {
	if l1, l2 := len(argSlice), len(expectSlice); l1 != l2 {
		panic(
			"\nWrong arguments to function `testFunc2VecsToF64`." +
				"\nlen(argSlice)!=len(expectSlice)" +
				fmt.Sprintf("\n%d != %d\n", l1, l2),
		)
	}
	for i, args := range argSlice {
		ex := expectSlice[i]
		g1, g2 := f(args[0], args[1])
		if ex.float64 != g1 || ex.bool != g2 {
			t.Fatalf("\n%s %d. tc failed! expect!=got\n(%v, %v) != (%v, %v)\n",
				t.Name(), i, ex.bool, ex.float64, g1, g2,
			)
		}
	}
}

// Tests
const trigTcQuantity = 12

var trigArgTab = [trigTcQuantity][2][]float64{
	// Empty/absent
	{nil, nil},
	{nil, {}},
	{{}, {}}, //2.
	// Possible: dimNum>0
	{{0}, {7}},       //3. Parallel
	{{0.6}, {8}},     // Parallel
	{{0, 1}, {1, 0}}, //5. Orthogonal
	// Normal
	{{8, 8, 8}, {-90, -90, -90}},   // Parallel
	{{9, 5, 13}, {-7, 10, 1}},      //7. Orthogonal
	{{0, -6, 0, 0}, {0, -6, 6, 0}}, //45° = π/4; cos α= 1/√2 = √2/2
	{{8, -6, 60}, {5, 9, -19}},     //9. cos α= -0.877902946208735; α= 151.39042257843425
	// With different lens.
	{{8}, {-75, 0}},               // Parallel
	{{16, 1, 6, 7, -9}, {-2, 32}}, //11. Orthogonal
}

func Test_CosOfAngleFuncs(t *testing.T) { //!
	// Algorithm
	t.Run("ForceCosOfAngle", func(t *testing.T) {
		expectTab := [trigTcQuantity]float64{
			// Empty/abset
			math.NaN(), math.NaN(), math.NaN(), //0..2.
			// Possible
			math.NaN(), //3.
			1,
			0, //5.
			// Normal
			-1,                 //6. P.
			0,                  //7. O.
			math.Sqrt2 / 2,     //8. π/4
			-0.877902946208735, //9. F.
			-1,                 //10. P.
			0,                  //11. O.
		}
		testFunc2VecsToF64(t, ForceCosOfAngle, trigArgTab[:], expectTab[:])
	})
	t.Run("NanableCosOfAngle", func(t *testing.T) {
		expectTab := [trigTcQuantity]float64{
			// Empty/abset
			math.NaN(), math.NaN(), math.NaN(), //0..2.
			// Possible
			math.NaN(), //3.
			1,
			0, //5.
			// Normal
			-1,                 //6. P.
			0,                  //7. O.
			math.Sqrt2 / 2,     //8. π/4
			-0.877902946208735, //9. F.
			math.NaN(),         //10. P.
			math.NaN(),         //11. O.
		}
		testFunc2VecsToF64(t, NanableCosOfAngle, trigArgTab[:], expectTab[:])
	})
	t.Run("CosOfAngle", func(t *testing.T) {
		expectTab := [trigTcQuantity]f64AndBool{
			// Empty/abset
			{0, true}, {0, true}, {0, true}, //0..2.
			// Possible
			{0, true},  //3.
			{1, false}, //4.
			{0, false}, //5.
			// Normal
			{-1, false},                 //6. P.
			{0, false},                  //7. O.
			{math.Sqrt2 / 2, false},     //8. π/4
			{-0.877902946208735, false}, //9. F.
			{0, true},
			{0, true},
		}
		testFunc2VecsToF64AndBool(t, CosOfAngle, trigArgTab[:], expectTab[:])
	})
}

func Test_AngleFuncs(t *testing.T) { //!
	// Algorithm
	t.Run("ForceAngle", func(t *testing.T) {
		expectTab := [trigTcQuantity]float64{
			// Empty/abset
			math.NaN(), math.NaN(), math.NaN(), //0..2.
			// Possible
			math.NaN(), //3.
			0,
			math.Pi / 2, //5.
			// Normal
			math.Pi,           //6. P.
			math.Pi / 2,       //7. O.
			math.Pi / 4,       //8.
			2.642261329979241, //9.
			math.Pi,
			math.Pi / 2,
		}

		testFunc2VecsToF64(t, ForceAngle, trigArgTab[:], expectTab[:])
	})
	t.Run("NanableAngle", func(t *testing.T) {
		expectTab := [trigTcQuantity]float64{
			// Empty/abset
			math.NaN(), math.NaN(), math.NaN(), //0..2.
			// Possible
			math.NaN(), //3.
			0,
			math.Pi / 2, //5.
			// Normal
			math.Pi,           //6. P.
			math.Pi / 2,       //7. O.
			math.Pi / 4,       //8.
			2.642261329979241, //9.
			math.NaN(),
			math.NaN(),
		}

		testFunc2VecsToF64(t, NanableAngle, trigArgTab[:], expectTab[:])
	})
	t.Run("Angle", func(t *testing.T) {
		expectTab := [trigTcQuantity]f64AndBool{
			// Empty/abset
			{0, true}, {0, true}, {0, true}, //0..2.
			// Possible
			{0, true}, //3.
			{0, false},
			{math.Pi / 2, false}, //5.
			// Normal
			{math.Pi, false},           //6. P.
			{math.Pi / 2, false},       //7. O.
			{math.Pi / 4, false},       //8.
			{2.642261329979241, false}, //9.
			{0, true},
			{0, true},
		}
		testFunc2VecsToF64AndBool(t, Angle, trigArgTab[:], expectTab[:])
	})
}
