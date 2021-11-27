package game

import "time"

type (
	AnimationEngine struct{}

	AnimationComponent interface {
		Update(dt time.Duration)
	}
)

// play animation (ie: depending on time only)
func (ae AnimationEngine) CalculateIntermediatePoses(bucket *EntityBucket, dt time.Duration) {

}

// correct animation parameters with previous effects
func (ae AnimationEngine) FinalizePoseAndMatrixPalette(bucket *EntityBucket) {

}

func NewAnimationEngine() *AnimationEngine {
	return &AnimationEngine{}
}
