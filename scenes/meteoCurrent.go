package scenes

import (
	"errors"
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/fonts"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/meteo"
	"image/color"
	"strconv"
	"strings"
	"time"
)

type MeteoCurrentData struct {
	city    string
	current int
	min     int
	//minTendancy  int
	max int
	//maxTendancy  int
	weather      int
	weatherLabel string
	riskType     string
	riskLabel    string
	riskPercent  int
}

type MeteoCurrent struct {
	shape              *CompositeDrawable
	meteoConceptClient *meteo.MeteoConceptClient
	data               *MeteoCurrentData
	dateTimeTextValue  string

	cityText        *shapes.Text
	dateTimeText    *shapes.Text
	tempCurrentText *shapes.Text
	tempMinText     *shapes.Text
	tempMaxText     *shapes.Text
	weatherText     *shapes.Text
	weatherIcon     *shapes.Img
	riskText        *shapes.Text
}

func NewMeteoCurrentComponent(_ Canvas, insee string) *MeteoCurrent {
	font := fonts.Bdf5x7
	m := MeteoCurrent{
		shape:              NewCompositeDrawable(NewGraphic(nil, NewLayout(ColorWhite, color.Transparent))),
		meteoConceptClient: meteo.NewMeteoConceptClient(insee),
		data:               nil,
		dateTimeTextValue:  "",
	}
	m.cityText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 4, Y: 10}, "Météo à ...", font)
	m.shape.AddDrawable(m.cityText)
	m.dateTimeText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 94, Y: 10}, m.dateTimeTextValue, font)
	m.shape.AddDrawable(m.dateTimeText)
	m.tempCurrentText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 10, Y: 20}, m.dateTimeTextValue, fonts.Bdf8x13)
	m.shape.AddDrawable(m.tempCurrentText)
	m.tempMinText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 4, Y: 39}, m.dateTimeTextValue, font)
	m.shape.AddDrawable(m.tempMinText)
	m.tempMaxText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 4, Y: 49}, m.dateTimeTextValue, font)
	m.shape.AddDrawable(m.tempMaxText)
	m.weatherText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 68, Y: 49}, m.dateTimeTextValue, font)
	m.shape.AddDrawable(m.weatherText)
	var imgPaths []string
	for _, name := range imgNames {
		imgPaths = append(imgPaths, "img/meteo/"+name+".png")
	}
	m.weatherIcon = shapes.NewPngFromPaths(NewOffsetGraphic(m.shape.Graphic, nil, Point{X: 82, Y: 16}), Point{X: 32, Y: 32}, imgPaths...)
	m.shape.AddDrawable(m.weatherIcon)
	m.riskText = shapes.NewText(NewOffsetGraphic(m.shape.Graphic, nil, Point{}), Point{X: 4, Y: 89}, m.dateTimeTextValue, font)
	m.shape.AddDrawable(m.riskText)
	return &m
}

func (m *MeteoCurrent) Update(_ time.Duration) bool {
	dataUpdated := m.updateData()
	datetimeUpdated := m.updateDatetime()
	dataUpdated = datetimeUpdated || dataUpdated

	if datetimeUpdated {
		m.dateTimeText.SetText(m.dateTimeTextValue)
	}

	if dataUpdated {
		m.updateTextContent()
	}

	return dataUpdated
}

func (m *MeteoCurrent) updateData() bool {
	var meteoCurrentData, err = m.getMeteoData()
	if err != nil {
		return false
	}
	if m.data == nil || m.data.differs(meteoCurrentData) {
		m.data = meteoCurrentData
		return true
	}
	return false
}

func (m *MeteoCurrent) updateTextContent() {
	if m.data != nil {
		m.cityText.SetText("Météo à " + m.data.city)
		m.tempCurrentText.SetText(strconv.Itoa(m.data.current) + "°C")
		m.tempMinText.SetText("Min: " + strconv.Itoa(m.data.min) + "°C")
		m.tempMaxText.SetText("Max: " + strconv.Itoa(m.data.max) + "°C")
		m.weatherText.SetText(m.data.weatherLabel)
		if m.data.riskPercent > 0 {
			m.riskText.SetText("Risque de " + m.data.riskLabel + ": " + strconv.Itoa(m.data.riskPercent) + "%")
		} else {
			m.riskText.SetText("")
		}
		var index = weatherToMeteoImageIndex(m.data.weather)
		if index >= 0 {
			m.weatherIcon.SetActiveImage(&(*m.weatherIcon.Images())[index])
		}
	}
}

var monthReplacerFR = strings.NewReplacer(
	"January", "Janvier",
	"February", "Février",
	"March", "Mars",
	"April", "Avril",
	"May", "Mai",
	"June", "Juin",
	"July", "Juillet",
	"August", "Août",
	"September", "Septembre",
	"October", "Octobre",
	"November", "Novembre",
	"December", "Décembre")
var shortMonthReplacerFR = strings.NewReplacer(
	"January", "Jan",
	"February", "Fév",
	"March", "Mar",
	"April", "Avr",
	"May", "Mai",
	"June", "Jui",
	"July", "Jui",
	"August", "Aoû",
	"September", "Sep",
	"October", "Oct",
	"November", "Nov",
	"December", "Déc")

var dayReplacerFR = strings.NewReplacer(
	"Monday", "Lundi",
	"Tuesday", "Mardi",
	"Wednesday", "Mercredi",
	"Thursday", "Jeudi",
	"Friday", "Vendredi",
	"Saturday", "Samedi",
	"Sunday", "Dimanche")

func localizeFR(dateTime string) string {
	return monthReplacerFR.Replace(dayReplacerFR.Replace(dateTime))
}

func localizeShortFR(dateTime string) string {
	return shortMonthReplacerFR.Replace(dayReplacerFR.Replace(dateTime))
}
func formatDateTime(dateTime time.Time) string {
	return localizeShortFR(dateTime.Format("_2 January 06 15:04"))
}

func formatDate(dateTime time.Time) string {
	return localizeShortFR(dateTime.Format("Monday 2 January"))
}

func (m *MeteoCurrent) updateDatetime() bool {
	var formatted = formatDateTime(time.Now())
	if formatted != m.dateTimeTextValue {
		m.dateTimeTextValue = formatted
		return true
	}
	return false
}

func (m *MeteoCurrent) getMeteoData() (*MeteoCurrentData, error) {
	forecastNextHours, err := m.meteoConceptClient.ForecastNextHours()
	if err != nil {
		return nil, err

	}
	forecastDaily, err := m.meteoConceptClient.ForecastDaily()
	if err != nil {
		return nil, err

	}
	currentHourForecast := forecastNextHours.Forecast[0]
	var riskPercent = 0
	var riskType, riskLabel string
	for _, dailyForecast := range forecastNextHours.Forecast {
		dailyRiskPercent, dailyRiskType, dailyRiskLabel := m.dailyRisk(dailyForecast)
		if dailyRiskPercent > riskPercent {
			riskPercent = dailyRiskPercent
			riskType = dailyRiskType
			riskLabel = dailyRiskLabel
		}
	}
	for _, dailyForecast := range forecastDaily.Forecast {
		if dailyForecast.Day == 0 {
			return &MeteoCurrentData{
				city:    forecastNextHours.City.Name,
				current: currentHourForecast.Temp2m,
				min:     dailyForecast.Tmin,
				//minTendancy:  0,
				max: dailyForecast.Tmax,
				//maxTendancy:  0,
				weather:      dailyForecast.Weather,
				weatherLabel: meteo.WeatherLabels[dailyForecast.Weather],
				riskType:     riskType,
				riskLabel:    riskLabel,
				riskPercent:  riskPercent,
			}, nil
		}
	}
	return nil, errors.New("no forecast")
}

func (m *MeteoCurrent) dailyRisk(dailyForecast meteo.ForecastHour) (int, string, string) {
	var riskPercent = 0
	var riskType = ""
	var riskLabel = ""
	if dailyForecast.Probawind100 > riskPercent {
		riskPercent = dailyForecast.Probawind100
		riskLabel = "Vent >100 km/h"
		riskType = "wind100"
	}
	if dailyForecast.Probawind70 > riskPercent {
		riskPercent = dailyForecast.Probawind70
		riskLabel = "Vent >70 km/h"
		riskType = "wind70"
	}
	if dailyForecast.Probafrost > riskPercent {
		riskPercent = dailyForecast.Probafrost
		riskLabel = "Gel"
		riskType = "frost"
	}
	if dailyForecast.Probafog > riskPercent {
		riskPercent = dailyForecast.Probafog
		riskLabel = "Brouillard"
		riskType = "fog"
	}
	if dailyForecast.Probarain > riskPercent {
		riskPercent = dailyForecast.Probarain
		riskLabel = "Pluie"
		riskType = "rain"
	}
	return riskPercent, riskType, riskLabel
}

func (md *MeteoCurrentData) differs(data *MeteoCurrentData) bool {
	return md.weather != data.weather ||
		md.max != data.max ||
		//md.maxTendancy != data.maxTendancy ||
		md.min != data.min ||
		//md.minTendancy != data.minTendancy ||
		md.riskType != data.riskType ||
		md.riskLabel != data.riskLabel ||
		md.riskPercent != data.riskPercent
}

func (m *MeteoCurrent) Draw(canvas Canvas) error {
	return m.shape.Draw(canvas)
}
