package vector

// The sense of this function is yet UNDEFINED.
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

// If vectors have the same number of dimensions, then
//		returns new vector being projection of `baseVec` on `targetVec`.
// Otherwise returns nil .
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

/*
If vectors have the same number of dimensions, then returns
	I. new vector being projection of `baseVec` on `targetVec`.
	II. false
Otherwise returns
	I. nil
	II. true
*/
func Proj(
	baseVec []float64,
	targetVec []float64,
	forceDotProd func(vec1, vec2 []float64) float64,
) (projection []float64, impossible bool) {

	projection = NilableProj(baseVec, targetVec, forceDotProd)
	impossible = projection == nil

	return
}

/*
`NilableOrthogonalizedWith` presumes that `targets` is a set of already orthogonal vectors.
This is prerequisite for a meaningful result.

If for every `el` in `targets` len(`el`) == len(`vec`), than returns
	new vector that is orthogonal to every `el` .
Otherwise returns nil .
*/
func NilableOrthogonalizedWith(vec []float64, targets ...[]float64) []float64 {
	// Check
	l := len(vec)
	for _, targ := range targets {
		if l != len(targ) {
			return nil
		}
	}

	// Algorithm
	orthogonal := Clone(vec)
	for _, targ := range targets {
		ForceSub(orthogonal, ForceProj(vec, targ, nil))
	}
	return orthogonal
}

/*
`OrthogonalizedWith` presumes that `targets` is a set of already orthogonal vectors.
This is prerequisite for a meaningful result.

If for every `el` in `targets` len(`el`) == len(`vec`), than returns
	I. new vector that is orthogonal to every `el`
	II. false
Otherwise returns (nil,true)
*/
func OrthogonalizedWith(vec []float64, targets ...[]float64) (
	orthogonal []float64, impossible bool,
) {
	orthogonal = NilableOrthogonalizedWith(vec, targets...)
	impossible = orthogonal == nil
	return
}
