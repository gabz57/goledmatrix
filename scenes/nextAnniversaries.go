package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"sort"
	"strconv"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

func (p Person) nbDaysBeforeAnniversary() int {
	today := time.Now()
	nextAnniversaryDay := p.Birthday
	for (nextAnniversaryDay.Sub(today).Hours() / 24) < 1 {
		nextAnniversaryDay = nextAnniversaryDay.AddDate(1, 0, 0)
	}
	return int(nextAnniversaryDay.Sub(today).Hours() / 24)
}

var persons = []Person{
	newPerson("Tatiana", 18, time.April, 1986),
	newPerson("Maman", 11, time.November, 1959),
	newPerson("Papa", 6, time.January, 1960),
	newPerson("Marion", 22, time.April, 1991),
	newPerson("Bloublou", 5, time.February, 2021),
	newPerson("Thomas", 2, time.April, 1988),
}

// ByNextAnniversary implements sort.Interface based on the Birthday field.
type ByNextAnniversary []Person

func (a ByNextAnniversary) Len() int { return len(a) }
func (a ByNextAnniversary) Less(i, j int) bool {
	return a[i].nbDaysBeforeAnniversary() < a[j].nbDaysBeforeAnniversary()
}
func (a ByNextAnniversary) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func newPerson(name string, day int, month time.Month, year int) Person {
	return Person{
		Name:     name,
		Birthday: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

const nbPersonsDisplayed = 5

type NextAnniversaries struct {
	shape            *CompositeDrawable
	persons          []Person
	titleText        *shapes.Text
	nameTexts        []*shapes.Text
	anniversaryTexts []*shapes.Text
	nbDaysTexts      []*shapes.Text
	currentDate      string
}

func NewNextAnniversariesComponent(canvas Canvas) *NextAnniversaries {
	graphic := NewGraphic(nil, NewLayout(ColorWhite, ColorBlack))
	var na = NextAnniversaries{
		shape:            NewCompositeDrawable(graphic),
		persons:          persons,
		nameTexts:        make([]*shapes.Text, nbDays),
		anniversaryTexts: make([]*shapes.Text, nbDays),
		nbDaysTexts:      make([]*shapes.Text, nbDays),
		currentDate:      "",
	}

	//font := fonts.Bdf5x7
	fontSmall := fonts.Bdf4x6
	na.titleText = shapes.NewText(graphic, Point{X: 2, Y: 5}, "Prochains anniversaires", fontSmall)
	var drawableTitleText Drawable = na.titleText
	na.shape.AddDrawable(&drawableTitleText)

	graphics := NewOffsetGraphic(graphic, nil, Point{Y: 10})
	var offset = 18
	var currentOffset = 0

	for i := 0; i < nbPersonsDisplayed; i++ {

		na.nameTexts[i] = shapes.NewText(graphics, Point{X: 2, Y: 1 + currentOffset}, "", fontSmall)
		var drawableNameText Drawable = na.nameTexts[i]
		na.shape.AddDrawable(&drawableNameText)

		na.anniversaryTexts[i] = shapes.NewText(graphics, Point{X: 2, Y: 7 + currentOffset}, "", fontSmall)
		var drawableAnniversaryText Drawable = na.anniversaryTexts[i]
		na.shape.AddDrawable(&drawableAnniversaryText)

		na.nbDaysTexts[i] = shapes.NewText(graphics, Point{X: 90, Y: 7 + currentOffset}, "", fontSmall)
		var drawableNbDaysText Drawable = na.nbDaysTexts[i]
		na.shape.AddDrawable(&drawableNbDaysText)

		currentOffset += offset
	}
	return &na
}

func (na NextAnniversaries) Update(elapsedBetweenUpdate time.Duration) bool {
	date := formatDate(time.Now())
	if na.currentDate != date {
		na.currentDate = date
		sort.Sort(ByNextAnniversary(na.persons))

		for i := 0; i < nbPersonsDisplayed; i++ {
			person := na.persons[i]
			age := time.Now().AddDate(0, 0, person.nbDaysBeforeAnniversary()).Year() - person.Birthday.Year()
			txt := person.Name + " " + strconv.Itoa(age) + " an"
			if age > 1 {
				txt += "s"
			}
			na.nameTexts[i].SetText(txt)
			na.anniversaryTexts[i].SetText(formatDate(person.Birthday))
			na.nbDaysTexts[i].SetText(strconv.Itoa(person.nbDaysBeforeAnniversary()) + " jours")
		}

		return true
	}
	return false
}

func (na NextAnniversaries) Draw(canvas Canvas) error {
	return na.shape.Draw(canvas)
}
