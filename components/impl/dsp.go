package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gordonklaus/portaudio"
	"github.com/mjibson/go-dsp/spectral"
	"image/color"
	"log"
	"math"
	"sync"
	"time"
)

var frequencies = []float64{
	10, 14, 20, 30, 50, 70,
	100, 140, 200, 300, 500, 700,
	1000, 1400, 2000, 3000, 5000, 7000,
	10000, 14000, 20000,
}

const NB_FREQUENCIES = 20 // len(frequencies) - 1
const MAX_HEIGHT = 64
const BAR_WIDTH = 4
const BAR_SPACING = 0
const TOTAL_DSP_WIDTH = NB_FREQUENCIES * (1 + BAR_WIDTH + 2*BAR_SPACING)

type Dsp struct {
	shape          *CompositeDrawable
	horizontalLine *shapes.Line
	verticalLines  []*shapes.Line
	frequencyBars  []*frequencyBar
}

func NewDsp(c Canvas) Component {
	graphic := NewOffsetGraphic(nil, nil, Point{
		X: (c.Bounds().Max.X - TOTAL_DSP_WIDTH) / 2,
		Y: 0,
	})
	dsp := Dsp{
		shape:          NewCompositeDrawable(graphic),
		horizontalLine: nil,
		verticalLines:  make([]*shapes.Line, NB_FREQUENCIES+1),
		frequencyBars:  make([]*frequencyBar, NB_FREQUENCIES),
	}

	//dsp.horizontalLine = dsp.buildHorizontalLine(graphic)
	//dsp.shape.AddDrawable(dsp.horizontalLine)

	barGraphic := NewGraphic(graphic, NewLayout(ColorWhite, ColorWhite))
	for i := 0; i <= NB_FREQUENCIES; i++ {
		//dsp.verticalLines[i] = dsp.buildVerticalLine(graphic, i)
		//dsp.shape.AddDrawable(dsp.verticalLines[i])
		if i < NB_FREQUENCIES {
			dsp.frequencyBars[i] = dsp.buildFrequencyBar(barGraphic, i)
			dsp.shape.AddDrawable(dsp.frequencyBars[i])
		}
	}
	go analyze()
	return &dsp
}

func (dsp *Dsp) buildHorizontalLine(graphic *Graphic) *shapes.Line {
	return shapes.NewLine(graphic,
		Point{
			X: 0,
			Y: MAX_HEIGHT,
		},
		Point{
			X: TOTAL_DSP_WIDTH,
			Y: MAX_HEIGHT,
		},
	)
}

func (dsp *Dsp) buildVerticalLine(graphic *Graphic, i int) *shapes.Line {
	return shapes.NewLine(graphic,
		Point{
			X: i * (1 + BAR_WIDTH + 2*BAR_SPACING),
			Y: 0,
		},
		Point{
			X: i * (1 + BAR_WIDTH + 2*BAR_SPACING),
			Y: MAX_HEIGHT,
		},
	)
}
func (dsp *Dsp) buildFrequencyBar(graphic *Graphic, i int) *frequencyBar {
	r, g, b := HsvToRgb(float64(i)/NB_FREQUENCIES*360, 1, 1)

	return newFrequencyBar(graphic, Point{
		X: i*(1+BAR_WIDTH+2*BAR_SPACING) + BAR_SPACING + 1,
		Y: 0,
	}, true, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff})
}

func (dsp *Dsp) Update(elapsedBetweenUpdate time.Duration) bool {
	updated := false
	valsMu.Lock()
	for i, bar := range dsp.frequencyBars {
		bar.SetRate(vals[i])
		updated = bar.Update(elapsedBetweenUpdate) || updated
	}
	valsMu.Unlock()
	return updated
}

func (dsp *Dsp) Draw(canvas Canvas) error {
	return dsp.shape.Draw(canvas)
}

// Analyzer

var (
	valsMu sync.Mutex
	vals   []float64
)

type minMax struct {
	Min, Max int
}

func analyze() {
	if err := portaudio.Initialize(); err != nil {
		log.Fatal(err)
	}
	defaultOutput, _ := portaudio.DefaultInputDevice()
	log.Print("in from ", defaultOutput.Name)
	in := make([]int32, 2048)
	inF := make([]float64, 2048)

	f2B := func(freq float64) int {
		return freq2Bin(freq, 44100, float64(len(in)))
	}
	var buckets = make([]minMax, len(frequencies)-1)
	for i := 0; i < len(frequencies)-1; i++ {
		buckets[i] = minMax{
			Min: f2B(frequencies[i]),
			Max: f2B(frequencies[i+1]),
		}
	}

	thresholds := make([]float64, len(buckets))
	tmpVals := make([]float64, len(buckets)) // temp buffer

	prevPxx := make([]float64, len(in))

	valsMu.Lock()
	vals = make([]float64, len(buckets))
	valsMu.Unlock()

	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	check(err)
	check(stream.Start())
	for {
		check(stream.Read())
		for i := range in {
			inF[i] = float64(in[i])
		}

		pxx, _ := spectral.Pwelch(inF, float64(44100), &spectral.PwelchOptions{
			NFFT: len(in),
		})

		const smoothing = .2

		// smooth with previous
		for i := range pxx {
			pxx[i] = prevPxx[i]*smoothing + pxx[i]*(1-smoothing)
		}
		prevPxx = pxx

		valsMu.Lock()
		for i := range vals {
			min := buckets[i].Min
			max := buckets[i].Max + 1
			tmpVals[i] = 0
			for n := min; n < max; n++ {
				tmpVals[i] += lerp(90, 200, math.Log10(pxx[n])*10) // db min/max
			}
			tmpVals[i] /= float64(max - min)
			if tmpVals[i] > thresholds[i] {
				thresholds[i] = tmpVals[i]

				// on beat
				vals[i] = tmpVals[i]
			} else {
				vals[i] = 0
			}
		}
		valsMu.Unlock()

		for i := range thresholds {
			//thresholds[i] *= .996 // decay
			thresholds[i] *= .2 // decay
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func lerp(min, max float64, val float64) float64 {
	if val < min {
		return 0
	}
	if val > max {
		return 1
	}
	return (val - min) / (max - min)
}

func freq2Bin(freq float64, sampleRate float64, fftSize float64) int {
	return int(freq / (sampleRate / fftSize))
}

func bin2Freq(bin int, sampleRate float64, fftSize float64) float64 {
	return float64(bin) * (sampleRate / fftSize)
}
