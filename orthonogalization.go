package vector

/*
Returns new vector being projection of `baseVec` on `targetVec`.

Function assumes that n-dimensional vector has trailing
zeros for all dimensions > n.
And so it can operate on vectors of apparently different number of dimensions.

Note: the resulting slice will have the same number dimensions as `targetVec`.
*/
func ForceProj(
	baseVec []float64,
	targetVec []float64,
	forceDotProd func(vec1, vec2 []float64) float64,
) []float64 {
	// Default arg. dotProd for cartesian cordinate system
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

/*Not examinated subfunction of `NilableOrthogonalizedWith`.*/
func forceOrthogonalizedWith(vec []float64, targets ...[]float64) []float64 {
	// Simple check of vector dependence.
	if len(targets) >= len(vec) {
		return nil
	}

	// Algorithm
	orthogonal := Clone(vec)
	for _, targ := range targets {
		ForceSub(orthogonal, ForceProj(vec, targ, nil))
	}
	return orthogonal
}

/*
`NilableOrthogonalizedWith` presumes that `targets`
is a set of already orthogonal (=> not-dependant) vectors.
This is prerequisite for a meaningful result.

If [for every `el` in `targets` len(`el`) == len(`vec`)] and [len(`targets`)>len(`vec`)], than returns
	new vector that is orthogonal to every `el` .
Otherwise returns nil .
*/
func NilableOrthogonalizedWith(vec []float64, targets ...[]float64) []float64 {
	// Len. check
	l := len(vec)
	for _, targ := range targets {
		if l != len(targ) {
			return nil
		}
	}

	return forceOrthogonalizedWith(vec, targets...)
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
