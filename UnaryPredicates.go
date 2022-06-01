package vector

import "math"

/*Returns false if there's no no-NaN number in `f64Slice`.
Otherwise true.*/
func IsNan(f64Slice []float64) bool {
	for _, f64 := range f64Slice {
		if !math.IsNaN(f64) {
			return false
		}
	}
	return true
}

// Returns true if there's atleast one NaN in `f64Slice`.
func HasNan(f64Slice []float64) bool {
	for _, f64 := range f64Slice {
		if math.IsNaN(f64) {
			return true
		}
	}
	return false
}
