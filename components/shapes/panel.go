package shapes

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"time"
)

type Panel struct {
	shape        *CompositeDrawable
	cornerRadius int
	dimensions   Point
	fill         bool
	border       bool
}

func NewPanel(parent *Graphic, layout *Layout, initialPosition Point, dimensions Point, cornerRadius int, fill, border bool) *Panel {
	p := Panel{
		shape: NewCompositeDrawable(
			NewOffsetGraphic(
				parent,
				layout,
				initialPosition),
		),
		cornerRadius: cornerRadius,
		dimensions:   dimensions,
		fill:         fill,
		border:       border,
	}

	if border {
		p.shape.AddDrawable(p.buildTopLine())
		p.shape.AddDrawable(p.buildBottomLine())
		p.shape.AddDrawable(p.buildLeftLine())
		p.shape.AddDrawable(p.buildRightLine())
	}

	if fill {
		// FIXME: adapt size when no border ??
		var fillLayout = NewLayout(*p.shape.Graphic.Layout().BackgroundColor(), nil)
		p.shape.AddDrawable(p.buildFillingCenter(fillLayout))
		if cornerRadius > 0 {
			p.shape.AddDrawable(p.buildFillingLeft(fillLayout))
			p.shape.AddDrawable(p.buildFillingRight(fillLayout))
		}
	}

	if cornerRadius > 0 {
		// FIXME: adapt circle size when no border ??
		p.shape.AddDrawable(p.buildTopLeftCorner())
		p.shape.AddDrawable(p.buildTopRightCorner())
		p.shape.AddDrawable(p.buildBottomLeftCorner())
		p.shape.AddDrawable(p.buildBottomRightCorner())
	}

	return &p
}

func (p *Panel) Update(elapsedBetweenUpdate time.Duration) {
}

func (p *Panel) Draw(canvas Canvas) error {
	return p.shape.Draw(canvas)
}

func (p *Panel) buildTopLine() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var line Drawable
	line = NewLine(graphic,
		Point{
			X: p.cornerRadius,
			Y: 0,
		},
		Point{
			X: p.dimensions.X - p.cornerRadius,
			Y: 0,
		},
	)
	return &line
}

func (p *Panel) buildBottomLine() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var line Drawable
	line = NewLine(graphic,
		Point{
			X: p.cornerRadius,
			Y: p.dimensions.Y,
		},
		Point{
			X: p.dimensions.X - p.cornerRadius,
			Y: p.dimensions.Y,
		},
	)
	return &line
}

func (p *Panel) buildLeftLine() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var line Drawable
	line = NewLine(graphic,
		Point{
			X: 0,
			Y: p.cornerRadius,
		},
		Point{
			X: 0,
			Y: p.dimensions.Y - p.cornerRadius,
		},
	)
	return &line
}

func (p *Panel) buildRightLine() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var line Drawable
	line = NewLine(graphic,
		Point{
			X: p.dimensions.X,
			Y: p.cornerRadius,
		},
		Point{
			X: p.dimensions.X,
			Y: p.dimensions.Y - p.cornerRadius,
		},
	)
	return &line
}

func (p *Panel) buildTopLeftCorner() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var circle Drawable
	circle = NewCircleTL(graphic, Point{
		X: p.cornerRadius,
		Y: p.cornerRadius,
	}, p.cornerRadius, p.fill)
	return &circle
}

func (p *Panel) buildTopRightCorner() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var circle Drawable
	circle = NewCircleTR(graphic, Point{
		X: p.dimensions.X - p.cornerRadius,
		Y: p.cornerRadius,
	}, p.cornerRadius, p.fill)
	return &circle
}

func (p *Panel) buildBottomLeftCorner() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var circle Drawable
	circle = NewCircleBL(graphic, Point{
		X: p.cornerRadius,
		Y: p.dimensions.Y - p.cornerRadius,
	}, p.cornerRadius, p.fill)
	return &circle
}

func (p *Panel) buildBottomRightCorner() *Drawable {
	graphic := NewGraphic(p.shape.Graphic, nil)
	var circle Drawable
	circle = NewCircleBR(graphic, Point{
		X: p.dimensions.X - p.cornerRadius,
		Y: p.dimensions.Y - p.cornerRadius,
	}, p.cornerRadius, p.fill)
	return &circle
}

func (p *Panel) buildFillingCenter(layout *Layout) *Drawable {

	graphic := NewGraphic(p.shape.Graphic, layout)
	var rectangle Drawable
	rectangle = NewRectangle(graphic, Point{
		X: p.cornerRadius + 1,
		Y: 1,
	}, Point{
		X: p.dimensions.X - 2*p.cornerRadius - 2,
		Y: p.dimensions.Y - 2,
	}, true)
	return &rectangle
}

func (p *Panel) buildFillingLeft(layout *Layout) *Drawable {
	graphic := NewGraphic(p.shape.Graphic, layout)
	var rectangle Drawable
	rectangle = NewRectangle(graphic, Point{
		X: 1,
		Y: p.cornerRadius + 1,
	}, Point{
		X: p.cornerRadius - 1,
		Y: p.dimensions.Y - 2*p.cornerRadius - 2,
	}, true)
	return &rectangle
}

func (p *Panel) buildFillingRight(layout *Layout) *Drawable {
	graphic := NewGraphic(p.shape.Graphic, layout)
	var rectangle Drawable
	rectangle = NewRectangle(graphic, Point{
		X: p.dimensions.X - p.cornerRadius,
		Y: p.cornerRadius + 1,
	}, Point{
		X: p.cornerRadius - 1,
		Y: p.dimensions.Y - 2*p.cornerRadius - 2,
	}, true)
	return &rectangle
}
