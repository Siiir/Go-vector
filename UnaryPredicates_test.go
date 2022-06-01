package vector

import (
	"math"
	"reflect"
	"testing"
)

func testUnaryPred[inT any](t *testing.T, pred func(vec inT) bool, argSlice []inT, expectSlice []bool) {
	if len(argSlice) != len(expectSlice) {
		panic("argument error: len(`argSlice`)!=len(`expectSlice`)")
	}

	for i, arg := range argSlice {
		ex := expectSlice[i]
		got := pred(arg)

		if !reflect.DeepEqual(ex, got) {
			t.Errorf("\nTc %d failed! ex != got\n\t%v != %v",
				i, ex, got)
		}
	}
}

func TestNanCheckingFuncs(t *testing.T) {
	const tcQuantity = 7
	argTab := [tcQuantity][]float64{
		{math.NaN()},                         //0. Is & Has
		{math.NaN(), math.NaN(), math.NaN()}, //1. Is & Has
		{0, math.NaN(), 0, 0},                //2. HasNan
		{math.NaN(), 7, 34, 5, 4e34, 0, 99},  //3. HasNan
		nil,                                  //4. IsNan
		{},                                   //5. IsNan
		{5, 3, -8, 0},                        //6.
	}

	t.Run("IsNan", func(t *testing.T) {
		expectTab := [tcQuantity]bool{true, true, false, false, true, true, false}
		testUnaryPred(t, IsNan, argTab[:], expectTab[:])
	})

	t.Run("HasNan", func(t *testing.T) {
		expectTab := [tcQuantity]bool{true, true, true, true}
		testUnaryPred(t, HasNan, argTab[:], expectTab[:])
	})
}
