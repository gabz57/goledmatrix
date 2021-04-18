package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"math/rand"
	"time"
)

type BeanNDot struct {
	beanGif   *shapes.Img
	movingDot *MovingDot
}

func NewBeanNDot(c Canvas) *BeanNDot {
	graphic := NewGraphic(nil, nil)
	return &BeanNDot{
		beanGif: shapes.NewGifFromFiles(
			graphic,
			Point{X: 128, Y: 78},
			100*time.Millisecond,
			"img/bean0.gif",
			"img/bean1.gif",
			"img/bean2.gif",
			"img/bean3.gif",
			"img/bean4.gif",
			"img/bean5.gif",
			"img/bean6.gif",
			"img/bean7.gif",
			"img/bean8.gif",
			"img/bean9.gif",
			"img/bean10.gif",
			"img/bean11.gif",
		),
		movingDot: NewMovingDot(
			c,
			Point{
				X: rand.Intn(128),
				Y: int(Int64Between(79, 128)),
			},
			FloatingPoint{
				X: Float64Between(32, 64),
				Y: Float64Between(32, 64),
			},
			image.Rectangle{
				Min: image.Point{X: 0, Y: 79},
				Max: c.Bounds().Max,
			}),
	}
}
func (b *BeanNDot) Update(elapsedBetweenUpdate time.Duration) {
	b.beanGif.Update(elapsedBetweenUpdate)
	b.movingDot.Update(elapsedBetweenUpdate)
}

func (b *BeanNDot) Draw(canvas Canvas) error {
	err := b.movingDot.Draw(canvas)
	if err != nil {
		return err
	}
	return b.beanGif.Draw(canvas)
}
