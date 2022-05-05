package vector

func Normalize(vec []float64) {
	DivByScalar(vec, CalcNorm(vec))
}

func Normalized(vec []float64) (normalized []float64) {
	clone := Clone(vec)
	Normalize(clone)
	return clone
}

func Remap(sl []float64, f func(float64) float64) {
	for i := range sl {
		sl[i] = f(sl[i])
	}
}

func Remaped(sl []float64,
	f func(float64) float64,
) (new []float64) {
	new = make([]float64, len(sl))
	for i := range sl {
		new[i] = f(sl[i])
	}
	return
}

func Proj(
	baseVec []float64,
	targetVec []float64,
	dotProd func(vec1, vec2 []float64) float64,
) []float64 {
	return ProdWithScalar(
		targetVec,
		dotProd(targetVec, baseVec)/dotProd(targetVec, targetVec),
	)
}

func ProjForCartesianSpace(baseVec, targetVec []float64) []float64 {
	return Proj(baseVec, targetVec, DotProd)
}

func OrthogonalizedWith(vec []float64, targets ...[]float64) (
	orthogonal []float64,
) {
	orthogonal = Clone(vec)
	for _, targ := range targets {
		ForceSub(orthogonal, ProjForCartesianSpace(vec, targ))
	}
	return
}
