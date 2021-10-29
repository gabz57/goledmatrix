package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"github.com/gabz57/goledmatrix/meteo"
	"image"
	"strconv"
	"strings"
	"time"
)

type MeteoForecastDailyData struct {
	date         time.Time
	weather      int
	weatherLabel string
	min          int
	max          int
}

type MeteoForecastData struct {
	city string
	days []MeteoForecastDailyData
}

type MeteoForecast struct {
	shape              *CompositeDrawable
	meteoConceptClient *meteo.MeteoConceptClient
	data               *MeteoForecastData
	dateTimeTextValue  string

	cityText     *shapes.ScrollingText
	dateTimeText *shapes.ScrollingText

	dateText    []*shapes.ScrollingText
	weatherText []*shapes.ScrollingText
	weatherIcon []*shapes.Img
	tempMinText []*shapes.Text
	tempMaxText []*shapes.Text
	separator   []*shapes.Line
}

var nbDays = 6

const Day = "day"
const Night = "night"
const Rainy = "rainy"
const CloudyDay1 = "cloudy-day-1"
const CloudyDay2 = "cloudy-day-2"
const CloudyDay3 = "cloudy-day-3"

var imgNames = []string{
	Day,
	Night,
	Rainy,
	CloudyDay1,
	CloudyDay2,
	CloudyDay3,
}

func NewMeteoForecastComponent(canvas Canvas, insee string) *MeteoForecast {

	var meteoForecastGraphic = NewOffsetGraphic(nil, NewLayout(ColorWhite, ColorBlack), Point{Y: 6})
	font := fonts.Bdf5x7
	fontSmall := fonts.Bdf4x6

	var m = MeteoForecast{
		shape:              NewCompositeDrawable(meteoForecastGraphic),
		meteoConceptClient: meteo.NewMeteoConceptClient(insee),
		data:               nil,
		dateTimeTextValue:  "",
	}

	m.cityText = shapes.NewScrollingText(meteoForecastGraphic, canvas, "Prévisions à ...", font, Point{X: 2, Y: 0}, image.Rect(0, 0, 88, 7), 12*time.Second)
	var drawableCityText Drawable = m.cityText
	m.shape.AddDrawable(&drawableCityText)

	m.dateTimeText = shapes.NewScrollingText(meteoForecastGraphic, canvas, m.dateTimeTextValue, font, Point{X: 95, Y: 0}, image.Rect(0, 0, 31, 7), 12*time.Second)
	var drawableDateTimeText Drawable = m.dateTimeText
	m.shape.AddDrawable(&drawableDateTimeText)

	m.dateText = make([]*shapes.ScrollingText, nbDays)
	m.weatherText = make([]*shapes.ScrollingText, nbDays)
	m.weatherIcon = make([]*shapes.Img, nbDays)
	m.tempMinText = make([]*shapes.Text, nbDays)
	m.tempMaxText = make([]*shapes.Text, nbDays)
	m.separator = make([]*shapes.Line, nbDays-1)

	graphics := NewOffsetGraphic(meteoForecastGraphic, nil, Point{Y: 10})
	var offset = 18
	var currentOffset = 0
	for i := 0; i < nbDays; i++ {
		m.dateText[i] = shapes.NewScrollingText(graphics, canvas, "", font, Point{X: 2, Y: currentOffset}, image.Rect(0, 0, 78, 7), 12*time.Second)
		var drawableDateText Drawable = m.dateText[i]
		m.shape.AddDrawable(&drawableDateText)
		m.weatherText[i] = shapes.NewScrollingText(graphics, canvas, "", font, Point{X: 2, Y: 7 + currentOffset}, image.Rect(0, 0, 78, 7), 12*time.Second)
		var drawableWeatherText Drawable = m.weatherText[i]
		m.shape.AddDrawable(&drawableWeatherText)

		//m.weatherIcon[i] = shapes.NewRectangle(graphics, Point{81, currentOffset}, Point{12,12}, false)
		var imgPaths []string
		for _, name := range imgNames {
			imgPaths = append(imgPaths, "img/meteo/"+name+".png")
		}
		m.weatherIcon[i] = shapes.NewPngFromPaths(NewOffsetGraphic(graphics, nil, Point{X: 80, Y: currentOffset - 1}), Point{X: 16, Y: 16}, imgPaths...)
		var drawableIcon Drawable = m.weatherIcon[i]
		m.shape.AddDrawable(&drawableIcon)

		m.tempMinText[i] = shapes.NewText(graphics, Point{X: 95, Y: 1 + currentOffset}, m.dateTimeTextValue, fontSmall)
		var drawableTempMinText Drawable = m.tempMinText[i]
		m.shape.AddDrawable(&drawableTempMinText)
		m.tempMaxText[i] = shapes.NewText(graphics, Point{X: 95, Y: 7 + currentOffset}, m.dateTimeTextValue, fontSmall)
		var drawableTempMaxText Drawable = m.tempMaxText[i]
		m.shape.AddDrawable(&drawableTempMaxText)

		if i != nbDays-1 {
			m.separator[i] = shapes.NewLine(graphics, Point{X: 3, Y: currentOffset + offset - 3}, Point{X: 124, Y: currentOffset + offset - 3})
			var drawableLine Drawable = m.separator[i]
			m.shape.AddDrawable(&drawableLine)
		}

		currentOffset += offset
	}

	return &m
}

func (m *MeteoForecast) Update(elapsedBetweenUpdate time.Duration) bool {
	updatedDateTime := m.updateDatetime()
	updatedMeteo := m.updateMeteoData()
	updated := updatedDateTime || updatedMeteo

	if updatedDateTime {
		m.updateDateTextContent()
	}

	if updatedMeteo {
		m.updateMeteoIconContent()
		m.updateMeteoTextContent()
	}

	updated = m.cityText.Update(elapsedBetweenUpdate) || updated
	updated = m.dateTimeText.Update(elapsedBetweenUpdate) || updated
	for _, text := range m.dateText {
		updated = text.Update(elapsedBetweenUpdate) || updated
	}
	for _, text := range m.weatherText {
		updated = text.Update(elapsedBetweenUpdate) || updated
	}

	return updated
}

func (m *MeteoForecast) updateMeteoData() bool {
	var meteoForecastData, err = m.getForecastData()
	if err != nil {
		return false
	}
	if m.data == nil || m.data.differs(meteoForecastData) {
		m.data = meteoForecastData
		return true
	}
	return false
}

func (m *MeteoForecast) updateDateTextContent() {
	m.dateTimeText.SetText(m.dateTimeTextValue)
}

func (m *MeteoForecast) updateMeteoTextContent() {
	if m.data != nil {
		m.cityText.SetText("Prévisions à " + strings.ReplaceAll(m.data.city, "Arrondissement", "Arr."))
		var nbMaxDays = nbDays
		if len(m.data.days) < nbDays {
			nbMaxDays = len(m.data.days)
		}
		for i := 0; i < nbMaxDays; i++ {
			if i == 0 {
				m.dateText[i].SetText("Aujourd'hui")
			} else {
				m.dateText[i].SetText(formatDate(m.data.days[i].date))
			}
			m.weatherText[i].SetText(m.data.days[i].weatherLabel)
			m.tempMinText[i].SetText("Min:" + strconv.Itoa(m.data.days[i].min) + "°C")
			m.tempMaxText[i].SetText("Max:" + strconv.Itoa(m.data.days[i].max) + "°C")
		}
	}
}

func (m *MeteoForecast) updateDatetime() bool {
	var formatted = formatDateTime(time.Now())
	if formatted != m.dateTimeTextValue {
		m.dateTimeTextValue = formatted
		return true
	}
	return false
}

func (m *MeteoForecast) getForecastData() (*MeteoForecastData, error) {
	forecastDaily, err := m.meteoConceptClient.ForecastDaily()
	if err != nil {
		return nil, err

	}
	var days = make([]MeteoForecastDailyData, len(forecastDaily.Forecast))

	for index, dailyForecast := range forecastDaily.Forecast {
		parsedDate, err := time.Parse("2006-01-02T15:04:05-0700", dailyForecast.Datetime)
		if err != nil {
			return nil, err
		}
		days[index] = MeteoForecastDailyData{
			date:         parsedDate,
			min:          dailyForecast.Tmin,
			max:          dailyForecast.Tmax,
			weather:      dailyForecast.Weather,
			weatherLabel: meteo.WeatherLabels[dailyForecast.Weather],
		}
	}
	return &MeteoForecastData{
		city: forecastDaily.City.Name,
		days: days,
	}, nil
}

func (mf *MeteoForecastData) differs(data *MeteoForecastData) bool {
	var differs = len(mf.days) != len(data.days)
	if differs {
		return true
	}
	for i, day := range data.days {
		if day.differs(mf.days[i]) {
			return true
		}
	}
	return false
}

func (mfdd MeteoForecastDailyData) differs(data MeteoForecastDailyData) bool {
	return mfdd.date != data.date ||
		mfdd.max != data.max ||
		mfdd.min != data.min ||
		mfdd.weather != data.weather ||
		mfdd.weatherLabel != data.weatherLabel
}

func (m *MeteoForecast) Draw(canvas Canvas) error {
	return m.shape.Draw(canvas)
}

func (m *MeteoForecast) updateMeteoIconContent() {
	var nbMaxDays = nbDays
	if len(m.data.days) < nbDays {
		nbMaxDays = len(m.data.days)
	}
	for i := 0; i < nbMaxDays; i++ {
		weather := m.data.days[i].weather
		var index = weatherToMeteoImageIndex(weather)
		if index >= 0 {
			m.weatherIcon[i].SetActiveImage(&(*m.weatherIcon[i].Images())[index])
		}
	}
}

// 0 Day = "day"
// 1 Night = "night"
// 2 Rainy = "rainy"
// 3 CloudyDay1 = "cloudy-day-1"
// 4 CloudyDay2 = "cloudy-day-2"
// 5 CloudyDay3 = "cloudy-day-3"
func weatherToMeteoImageIndex(weather int) int {
	switch weather {
	case 0:
		return 0
	case 1, 2:
		return 3
	case 3, 5:
		return 4
	case 4, 6, 7:
		return 5
	default:
		return 2
	}
}
