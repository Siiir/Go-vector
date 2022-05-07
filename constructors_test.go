package vector

import (
	"reflect"
	"testing"
)

func TestNewFilled(t *testing.T) {
	var ex, got []float64

	ex = []float64{-18, -18, -18, -18}
	got = NewFilled(4, -18.0)
	if !reflect.DeepEqual(ex, got) {
		t.Logf("%v != %v", ex, got)
	}

	v := 1.78e220
	ex = []float64{v, v}
	got = NewFilled(2, v)
	if !reflect.DeepEqual(ex, got) {
		t.Logf("%v != %v", ex, got)
	}

	ex = []float64{}
	got = NewFilled(0, v)
	if !reflect.DeepEqual(ex, got) {
		t.Logf("%v != %v", ex, got)
	}
}
