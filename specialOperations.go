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

func ForceProj( /// Works with dim-alignment
	/// But is it still mathematical (no hidden contradiction?)? I must talk to algebra teacher.
	baseVec []float64,
	targetVec []float64,
	forceDotProd func(vec1, vec2 []float64) float64,
) []float64 {
	/// Default arg. dotProd for cartesian cordinate
	if forceDotProd == nil {
		forceDotProd = ForceDotProd
	}

	return ProdWithScalar(
		targetVec, // Enforceses space of `targetVec``
		forceDotProd(targetVec, baseVec)/forceDotProd(targetVec, targetVec), // Here automatic alignment will be made
	)
}

func NilableProj(
	baseVec []float64,
	targetVec []float64,
	forceDotProd func(vec1, vec2 []float64) float64,
) []float64 {
	impossible := len(baseVec) != len(targetVec)
	if impossible {
		return nil
	}
	return ForceProj(baseVec, targetVec, forceDotProd)
}

func Proj(
	baseVec []float64,
	targetVec []float64,
	forceDotProd func(vec1, vec2 []float64) float64,
) (projection []float64, impossible bool) {

	projection = NilableProj(baseVec, targetVec, forceDotProd)
	impossible = projection == nil

	return
}

func OrthogonalizedWith(vec []float64, targets ...[]float64) (
	orthogonal []float64, impossible bool,
) {
	// Check
	l := len(vec)
	for _, targ := range targets {
		if l != len(targ) {
			impossible = true
			return
		}
	}

	// Algorithm
	orthogonal = Clone(vec)
	for _, targ := range targets {
		ForceSub(orthogonal, ForceProj(vec, targ, nil))
	}
	return
}
