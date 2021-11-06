package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"time"
)

type MovingHearts struct {
	canvas     Canvas
	hearts     []*MovingHeart
	origin     Point
	maxX, maxY int
}

func NewMovingHearts(canvas Canvas, origin Point, nbMovingHearts int) *MovingHearts {
	max := canvas.Bounds().Max
	h := MovingHearts{
		canvas: canvas,
		origin: origin,
		maxX:   max.X - origin.X - HeartWidth,
		maxY:   max.Y - origin.Y - HeartHeight,
	}
	for i := 0; i < nbMovingHearts; i++ {
		go h.addHeart()
	}
	return &h
}

func (h *MovingHearts) addHeart() {
	time.Sleep(time.Duration(components.Random.Int63n(2000)) * time.Millisecond)
	h.hearts = append(h.hearts, generateMovingHeart(h.canvas, h.maxX, h.maxY))
}

func (h *MovingHearts) replaceHeart(i int) {
	h.hearts[i] = generateMovingHeart(h.canvas, h.maxX, h.maxY)
}

func generateMovingHeart(canvas Canvas, maxX int, maxY int) *MovingHeart {
	return NewMovingHeart(
		canvas,
		Point{
			X: components.Random.Intn(maxX),
			Y: components.Random.Intn(maxY),
		},
		time.Duration(components.Int64Between(2000, 5000))*time.Millisecond,
		1,
		false)
}

func (h *MovingHearts) Update(elapsedBetweenUpdate time.Duration) bool {
	dirty := false
	for i, heart := range h.hearts {
		dirty = heart.Update(elapsedBetweenUpdate) || dirty
		if heart.IsOut() {
			h.replaceHeart(i)
			dirty = true
		}
	}
	return dirty
}

func (h *MovingHearts) Draw(canvas Canvas) error {
	for _, heart := range h.hearts {
		err := heart.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
}
