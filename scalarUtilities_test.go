package vector

import (
	"math/rand"
	"testing"
)

func Test_ptrInRange(t *testing.T) {
	// Definitions
	const tQuantity = 5
	ArgsTab := [tQuantity][3]uintptr{
		{90, 20, 1000},
		{200000, 200000, 200000},
		{200000, 10000000, 0},
		{200001, 200000, 200000},
		{200001, 200001, 200000},
	}
	ExpectTab := [tQuantity]bool{true, true, false, false, false}

	// Algorithm
	for i := 0; i < tQuantity; i++ {
		// Converted args go to `a` array
		var a [3]*float64
		for j, uPtr := range ArgsTab[i] {
			a[j] = newFloat64Ptr(uPtr)
		}

		//Now we can run our function
		if ex, got := ExpectTab[i], ptrInRange(a[0], a[1], a[2]); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}

func Example_giveRandSign() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(giveRandSign(rand.Float64()), "\t")
		}
		println()
	}
	//:-) Output:
}

func Example_randF64() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(randF64(), "\t")
		}
		println()
	}
	//:-) Output:
}

func Example_randNormalF64() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(randNormalF64(), "\t")
		}
		println()
	}
	//Output:
}
