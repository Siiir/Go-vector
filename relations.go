package vector

import "math"

/*
Returns true if vectors `vec1` & `vec2` are orthogonal.

The last argument is an acceptable distance of vectors' dot product from zero.
It must be (non-zero) taken into account when higher float64 values are in game
due to possible precision loss.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.

... These docs aren't complete.
*/
func AreOrthogonalAfterDimAlignment(vec1, vec2 []float64, fpErr float64) (
	are bool,
) {
	return math.Abs(ForceDotProd(vec1, vec2)) < fpErr
}

/*
Returns true if given slices are of equal length
and being vectors their are orthogonal.

The last argument is an acceptable distance of vectors' dot product from zero.
It must be (non-zero) taken into account when higher float64 values are in game

due to possible precision loss.
... These docs aren't complete.
*/
func AreAbsolutelyOrthogonal(vec1, vec2 []float64, fpErr float64) (
	are bool,
) {
	return (len(vec1) == len(vec2) &&
		AreOrthogonalAfterDimAlignment(vec1, vec2, fpErr))
}

/*
Returns true if given vectors are parallel.

The last argument is an acceptable distance between ratios of
the corresponding cordinates of two given vectors — which mathematically
ought to be equal.
It must be (non-zero) taken into account even if numbers in game are small.
If ratios are not whole numbers, there is high possibility of PRECISION ERROR.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.

... These docs aren't complete.
*/
func AreParallelAfterDimAlignment(v1, v2 []float64, fpDivErr float64) (are bool) {
	if l1, l2 := len(v1), len(v2); l1 != l2 {
		if l1-int(NumOfTrailingZeros(v1)) != l2-int(NumOfTrailingZeros(v2)) {
			//Still different len
			return false
		}
	}

	are = true
	defer recover()

	ratio := v1[0] / v2[0]
	for i := 1; ; i++ {
		if math.Abs(v1[i]/v2[i]-ratio) > fpDivErr {
			return false
		}
	}
}

/*Returns true if given slices are of equal length
and being vectors their are parallel.

The last argument is an acceptable distance between ratios of
the corresponding cordinates of two given vectors — which mathematically
ought to be equal.
It must be (non-zero) taken into account even if numbers in game are small.
If ratios are not whole numbers, there is high possibility of PRECISION ERROR.

... These docs aren't complete.
*/
func AreAbsolutelyParallel(vec1, vec2 []float64, fpDivErr float64) (are bool) {
	return (len(vec1) == len(vec2) &&
		AreParallelAfterDimAlignment(vec1, vec2, fpDivErr))
}
