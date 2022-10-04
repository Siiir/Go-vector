package vector

import (
	"testing"
)

func _TestNormalized(t *testing.T) {
	// Definitions
	const tcQuantity = 7
	argTab := [tcQuantity][]float64{
		// Not normalizable
		nil,
		{},
		{0, 0, 0},
		//
		{1, 0, 0},
		{4e45},
		{},
	}
	expectTab := [tcQuantity][]float64{
		// NewFilled(math.NaN()), NewFilled(math.NaN()), NewFilled(math.NaN()), //0..=2

	}

	// Algorithm
	for tcInd, vec := range argTab {
		ex := expectTab[tcInd]
		got := Normalized(vec)
		if !_HeapEquivalent(ex, got) {
			t.Errorf("\nTC %d failed! Not [expected<==>got]\n\texpected= %v\n\tgot= %v\n",
				tcInd, ex, got)
		}
	}
}

func TestRemap(t *testing.T) { panic("NIY") }

func TestRemapped(t *testing.T) { panic("NIY") }
