package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"time"
)

type Hearts struct {
	canvas     Canvas
	hearts     []*Heart
	origin     Point
	maxX, maxY int
}

func NewHearts(canvas Canvas, origin Point, nbHearts int) *Hearts {
	max := canvas.Bounds().Max
	h := Hearts{
		canvas: canvas,
		origin: origin,
		maxX:   max.X - origin.X - HeartWidth,
		maxY:   max.Y - origin.Y - HeartHeight,
	}
	for i := 0; i < nbHearts; i++ {
		go h.addHeart()
	}
	return &h
}

func (h *Hearts) addHeart() {
	time.Sleep(time.Duration(Random.Int63n(2000)) * time.Millisecond)
	h.hearts = append(h.hearts, generateHeart(h.canvas, h.maxX, h.maxY))
}

func (h *Hearts) replaceHeart(i int) {
	h.hearts[i] = generateHeart(h.canvas, h.maxX, h.maxY)
}

func generateHeart(canvas Canvas, maxX int, maxY int) *Heart {
	return NewHeart(
		canvas,
		nil,
		Point{
			X: Random.Intn(maxX),
			Y: Random.Intn(maxY),
		},
		time.Duration(Int64Between(2000, 5000))*time.Millisecond,
		1,
		false)
}

func (h *Hearts) Update(elapsedBetweenUpdate time.Duration) bool {
	dirty := false
	for i, heart := range h.hearts {
		dirty = heart.Update(elapsedBetweenUpdate) || dirty
		if heart.IsFaded() {
			h.replaceHeart(i)
			dirty = true
		}
	}
	return dirty
}

func (h *Hearts) Draw(canvas Canvas) error {
	for _, heart := range h.hearts {
		heart.Draw(canvas)
	}
	return nil
}
