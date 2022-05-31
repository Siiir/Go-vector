package vector

import "math"

// CosOfAngleFuncs

/*
Returns a cosine of angle between `vec1` & `vec2`.
If CalcNorm(`vec1`) * CalcNorm(`vec2`) == 0, the result will be math.NaN().

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func ForceCosOfAngle(vec1, vec2 []float64) float64 {
	return (ForceDotProd(vec1, vec2) /
		(CalcNorm(vec1) * CalcNorm(vec2)))
}

/*
If CalcNorm(`vec1`) * CalcNorm(`vec2`) == 0 or len(`vec1`)!=len(`vec2`),
	returns math.NaN() .
Otherwise returns a cosine of angle between `vec1` & `vec2`.
*/
func NanableCosOfAngle(vec1, vec2 []float64) float64 {
	if len(vec1) != len(vec2) {
		return math.NaN()
	}
	return ForceCosOfAngle(vec1, vec2)
}

/*
If both vectors have the same length and non-zero norm,
	returns a cosine of angle between `vec1` & `vec2`
		and false as second value.
Otherwise returns (0.0, true) .
*/
func CosOfAngle(vec1, vec2 []float64) (cos_x float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	denominator := CalcNorm(vec1) * CalcNorm(vec2)
	impossible = denominator == 0.0
	if impossible {
		return
	}

	cos_x = ForceDotProd(vec1, vec2) / denominator

	return
}

// AngleFuncs

/*
Returns a smaller angle between `vec1` & `vec2`.
If CalcNorm(`vec1`) * CalcNorm(`vec2`) == 0, the result will be math.NaN().

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func ForceAngle(vec1, vec2 []float64) float64 {
	return math.Acos(ForceCosOfAngle(vec1, vec2))
}

/*
If CalcNorm(`vec1`) * CalcNorm(`vec2`) == 0 or len(`vec1`)!=len(`vec2`),
	returns math.NaN() .
Otherwise returns a smaller angle between `vec1` & `vec2`.
*/
func NanableAngle(vec1, vec2 []float64) float64 {
	return math.Acos(NanableCosOfAngle(vec1, vec2))
}

/*
If both vectors have the same length and non-zero norm,
	returns a smaller angle between `vec1` & `vec2`
		and false as second value.
Otherwise returns (0.0, true) .
*/
func Angle(vec1, vec2 []float64) (angle float64, impossible bool) {
	cos_x, impossible := CosOfAngle(vec1, vec2)
	if impossible {
		return
	}

	return math.Acos(cos_x), impossible
}
