package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
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
			"img/bean/bean0.gif",
			"img/bean/bean1.gif",
			"img/bean/bean2.gif",
			"img/bean/bean3.gif",
			"img/bean/bean4.gif",
			"img/bean/bean5.gif",
			"img/bean/bean6.gif",
			"img/bean/bean7.gif",
			"img/bean/bean8.gif",
			"img/bean/bean9.gif",
			"img/bean/bean10.gif",
			"img/bean/bean11.gif",
		),
		movingDot: NewMovingDot(
			c,
			Point{
				X: Random.Intn(128),
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
func (b *BeanNDot) Update(elapsedBetweenUpdate time.Duration) bool {
	b.beanGif.Update(elapsedBetweenUpdate)
	b.movingDot.Update(elapsedBetweenUpdate)
	return true
}

func (b *BeanNDot) Draw(canvas Canvas) error {
	err := b.movingDot.Draw(canvas)
	if err != nil {
		return err
	}
	return b.beanGif.Draw(canvas)
}
