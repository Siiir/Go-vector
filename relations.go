package vector

import "math"

func AreOrthogonalAfterDimAlignment(vec1, vec2 []float64, fpErr float64) (
	are bool,
) {
	return math.Abs(ForceDotProd(vec1, vec2)) < fpErr
}
func AreAbsolutelyOrthogonal(vec1, vec2 []float64, fpErr float64) (
	are bool,
) {
	return (len(vec1) == len(vec2) &&
		AreOrthogonalAfterDimAlignment(vec1, vec2, fpErr))
}

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

func AreAbsolutelyParallel(vec1, vec2 []float64, fpDivErr float64) (are bool) {
	return (len(vec1) == len(vec2) &&
		AreParallelAfterDimAlignment(vec1, vec2, fpDivErr))
}
