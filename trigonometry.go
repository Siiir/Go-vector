package vector

import "math"

/*
Returns a cosine of angle between `vec1` & `vec2`.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func ForceCosOfAngle(vec1, vec2 []float64) float64 {
	return (ForceDotProd(vec1, vec2) /
		(CalcNorm(vec1) * CalcNorm(vec2)))
}

/*
If both vectors have the same length
	Returns a cosine of angle between `vec1` & `vec2`
		and false as second value.
Otherwise returns (0.0, true) .
*/
func CosOfAngle(vec1, vec2 []float64) (cos_x float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	cos_x = (ForceDotProd(vec1, vec2) /
		(CalcNorm(vec1) * CalcNorm(vec2)))

	return
}

/*
Returns an angle between `vec1` & `vec2`.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
Therefore it can operate on vectors with apparently different number of dimensions.
*/
func ForceAngle(vec1, vec2 []float64) float64 {
	/// My math knowledge says: "it should work"
	return math.Acos(ForceCosOfAngle(vec1, vec2))
}

/*
If both vectors have the same length
	Returns the smaller angle between `vec1` & `vec2`
		and false as second value.
Otherwise returns (0.0, true) .
*/
func Angle(vec1, vec2 []float64) (angle float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	angle = ForceAngle(vec1, vec2)
	return
}
