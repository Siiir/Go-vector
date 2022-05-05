package vector

import "math"

func ForceCosOfAngle(vec1, vec2 []float64) float64 {
	return (ForceDotProd(vec1, vec2) /
		(CalcNorm(vec1) * CalcNorm(vec2)))
}

func CosOfAngle(vec1, vec2 []float64) (cos_x float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}

	cos_x = (ForceDotProd(vec1, vec2) /
		(CalcNorm(vec1) * CalcNorm(vec2)))

	return
}

func ForceAngle(vec1, vec2 []float64) float64 {
	/// My math knowledge says: "it should work"
	return math.Acos(ForceCosOfAngle(vec1, vec2))
}

func Angle(vec1, vec2 []float64) (angle float64, impossible bool) {
	impossible = len(vec1) != len(vec2)
	if impossible {
		return
	}
	angle = ForceAngle(vec1, vec2)
	return
}
