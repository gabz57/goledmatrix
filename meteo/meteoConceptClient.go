package meteo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// I know this is a bad practice, but this token is limited
//to 500 call/day and only helps to retrieve weather data
var API_TOKEN = "7dd633b15368426444bba6c2d798fa3fe132358e5e4a1f905a37fdfd9a5f06e5"

const meteoDataValidDuration = 10 * time.Minute

var WeatherLabels = map[int]string{
	0:   "Soleil",
	1:   "Peu nuageux",
	2:   "Ciel voilé",
	3:   "Nuageux",
	4:   "Très nuageux",
	5:   "Couvert",
	6:   "Brouillard",
	7:   "Brouillard givrant",
	10:  "Pluie faible",
	11:  "Pluie modérée",
	12:  "Pluie forte",
	13:  "Pluie faible verglaçante",
	14:  "Pluie modérée verglaçante",
	15:  "Pluie forte verglaçante",
	16:  "Bruine",
	20:  "Neige faible",
	21:  "Neige modérée",
	22:  "Neige forte",
	30:  "Pluie et neige mêlées faibles",
	31:  "Pluie et neige mêlées modérées",
	32:  "Pluie et neige mêlées fortes",
	40:  "Averses de pluie locales et faibles",
	41:  "Averses de pluie locales",
	42:  "Averses locales et fortes",
	43:  "Averses de pluie faibles",
	44:  "Averses de pluie",
	45:  "Averses de pluie fortes",
	46:  "Averses de pluie faibles et fréquentes",
	47:  "Averses de pluie fréquentes",
	48:  "Averses de pluie fortes et fréquentes",
	60:  "Averses de neige localisées et faibles",
	61:  "Averses de neige localisées",
	62:  "Averses de neige localisées et fortes",
	63:  "Averses de neige faibles",
	64:  "Averses de neige",
	65:  "Averses de neige fortes",
	66:  "Averses de neige faibles et fréquentes",
	67:  "Averses de neige fréquentes",
	68:  "Averses de neige fortes et fréquentes",
	70:  "Averses de pluie et neige mêlées localisées et faibles",
	71:  "Averses de pluie et neige mêlées localisées",
	72:  "Averses de pluie et neige mêlées localisées et fortes",
	73:  "Averses de pluie et neige mêlées faibles",
	74:  "Averses de pluie et neige mêlées",
	75:  "Averses de pluie et neige mêlées fortes",
	76:  "Averses de pluie et neige mêlées faibles et nombreuses",
	77:  "Averses de pluie et neige mêlées fréquentes",
	78:  "Averses de pluie et neige mêlées fortes et fréquentes",
	100: "Orages faibles et locaux",
	101: "Orages locaux",
	102: "Orages fort et locaux",
	103: "Orages faibles",
	104: "Orages",
	105: "Orages forts",
	106: "Orages faibles et fréquents",
	107: "Orages fréquents",
	108: "Orages forts et fréquents",
	120: "Orages faibles et locaux de neige ou grésil",
	121: "Orages locaux de neige ou grésil",
	122: "Orages locaux de neige ou grésil",
	123: "Orages faibles de neige ou grésil",
	124: "Orages de neige ou grésil",
	125: "Orages de neige ou grésil",
	126: "Orages faibles et fréquents de neige ou grésil",
	127: "Orages fréquents de neige ou grésil",
	128: "Orages fréquents de neige ou grésil",
	130: "Orages faibles et locaux de pluie et neige mêlées ou grésil",
	131: "Orages locaux de pluie et neige mêlées ou grésil",
	132: "Orages fort et locaux de pluie et neige mêlées ou grésil",
	133: "Orages faibles de pluie et neige mêlées ou grésil",
	134: "Orages de pluie et neige mêlées ou grésil",
	135: "Orages forts de pluie et neige mêlées ou grésil",
	136: "Orages faibles et fréquents de pluie et neige mêlées ou grésil",
	137: "Orages fréquents de pluie et neige mêlées ou grésil",
	138: "Orages forts et fréquents de pluie et neige mêlées ou grésil",
	140: "Pluies orageuses",
	141: "Pluie et neige mêlées à caractère orageux",
	142: "Neige à caractère orageux",
	210: "Pluie faible intermittente",
	211: "Pluie modérée intermittente",
	212: "Pluie forte intermittente",
	220: "Neige faible intermittente",
	221: "Neige modérée intermittente",
	222: "Neige forte intermittente",
	230: "Pluie et neige mêlées",
	231: "Pluie et neige mêlées",
	232: "Pluie et neige mêlées",
	235: "Averses de grêle",
}

type (
	ForecastHourResponse struct {
		City     City           `json:"city"`     // Code Insee de la commune
		Update   string         `json:"update"`   // Code postal de la commune
		Forecast []ForecastHour `json:"forecast"` // Latitude décimale de la commune
	}

	ForecastDayResponse struct {
		City     City          `json:"city"`     // Code Insee de la commune
		Update   string        `json:"update"`   // Code postal de la commune
		Forecast []ForecastDay `json:"forecast"` // Latitude décimale de la commune
	}

	City struct {
		Country   string  `json:"city"`      // Code Insee de la commune
		Insee     string  `json:"insee"`     // Code Insee de la commune
		Cp        int     `json:"cp"`        // Code postal de la commune
		Name      string  `json:"name"`      // Nom de la commune
		Latitude  float32 `json:"latitude"`  // Latitude décimale de la commune
		Longitude float32 `json:"longitude"` // Longitude décimale de la commune
		Altitude  int     `json:"altitude"`  // Altitude de la commune en mètres
	}

	ForecastHour struct {
		//Insee        string  `json:"insee"`        // Code Insee de la commune
		//Cp           int     `json:"cp"`           // Code postal de la commune
		//Latitude     float32 `json:"latitude"`     // Latitude décimale de la commune
		//Longitude    float32 `json:"longitude"`    // Longitude décimale de la commune
		Datetime string `json:"datetime"` // Date en heure locale, format ISO8601
		Temp2m   int    `json:"temp2m"`   // Température à 2 mètres en °C
		//Rh2m         int     `json:"rh2m"`         // Humidité à 2 mètres en %
		//Wind10m      int     `json:"wind10m"`      // Vent moyen à 10 mètres en km/h
		//Gust10m      int     `json:"gust10m"`      // Rafales de vent à 10 mètres en km/h
		//Dirwind10m   int     `json:"dirwind10m"`   // Direction du vent en degrés (0 à 360°)
		//Rr10         float32 `json:"rr10"`         // Cumul de pluie sur la tranche horaire ou tri-horaire en mm
		//Rr1          float32 `json:"rr1"`          // Cumul de pluie maximal sur la tranche horaire ou tri-horaire en mm
		Probarain    int `json:"probarain"`    // Probabilité de pluie entre 0 et 100%
		Weather      int `json:"weather"`      // Temps sensible (Code temps) – Voir Annexes
		Probafrost   int `json:"probafrost"`   // Probabilité de gel entre 0 et 100%
		Probafog     int `json:"probafog"`     // Probabilité de brouillard entre 0 et 100%
		Probawind70  int `json:"probawind70"`  // Probabilité de vent >70 km/h entre 0 et 100%
		Probawind100 int `json:"probawind100"` // Probabilité de vent >100 km/h entre 0 et 100%
		//Tsoil1       int     `json:"tsoil1"`       // Température du sol entre 0 et 10 cm en °C
		//Tsoil2       int     `json:"tsoil2"`       // Température du sol entre 10 et 40 cm en °C.
		Gustx int `json:"gustx"` // Rafale de vent potentielle sous orage ou grain en km/h
		//Iso0         int     `json:"iso0"`         // Altitude du isotherme 0°C en mètres
	}

	ForecastDay struct {
		Insee string `json:"insee"` // Code Insee de la commune
		Cp    int    `json:"cp"`    // Code postal de la commune
		//Latitude     float32 `json:"latitude"`     // Latitude décimale de la commune
		//Longitude    float32 `json:"longitude"`    // Longitude décimale de la commune
		Day      int    `json:"day"`      // Jour entre 0 et 13 (Pour le jour même : 0, pour le lendemain : 1, etc.)
		Datetime string `json:"datetime"` // Date en heure locale, format ISO8601
		//Wind10m      int     `json:"wind10m"`      // Vent moyen à 10 mètres en km/h
		//Gust10m      int     `json:"gust10m"`      // Rafales de vent à 10 mètres en km/h
		//Dirwind10m   int     `json:"dirwind10m"`   // Direction du vent en degrés (0 à 360°)
		//Rr10         float32 `json:"rr10"`         // Cumul de pluie sur la journée en mm
		//Rr1          float32 `json:"rr1"`          // Cumul de pluie maximal sur la journée en mm
		Probarain int `json:"probarain"` // Probabilité de pluie entre 0 et 100%
		Weather   int `json:"weather"`   // Temps sensible (Code temps) – Voir Annexes
		Tmin      int `json:"tmin"`      // Température minimale à 2 mètres en °C
		Tmax      int `json:"tmax"`      // Température maximale à 2 mètres en °C
		SunHours  int `json:"sunHours"`  // Ensoleillement en heures
		//Etp          float32 `json:"etp"`          // Cumul d'évapotranspiration en mm
		Probafrost   int `json:"probafrost"`   // Probabilité de gel entre 0 et 100%
		Probafog     int `json:"probafog"`     // Probabilité de brouillard entre 0 et 100%
		Probawind70  int `json:"probawind70"`  // Probabilité de vent >70 km/h entre 0 et 100%
		Probawind100 int `json:"probawind100"` // Probabilité de vent >100 km/h entre 0 et 100%
		Gustx        int `json:"gustx"`        // Rafale de vent potentielle sous orage ou grain en km/h
		//Iso0         int     `json:"iso0"`         // Altitude du isotherme 0°C en mètres
	}
)

type MeteoConceptClient struct {
	insee                 string
	forecastHourResponse  *ForecastHourResponse
	forecastDayResponse   *ForecastDayResponse
	nextForecastHourFetch *time.Time
	nextForecastDayFetch  *time.Time
}

func NewMeteoConceptClient(insee string) *MeteoConceptClient {
	return &MeteoConceptClient{
		insee: insee,
	}
}
func (m *MeteoConceptClient) ForecastNextHours() (*ForecastHourResponse, error) {

	if m.forecastHourResponse == nil || m.nextForecastHourFetch == nil || time.Now().After(*m.nextForecastHourFetch) {

		var err error
		m.forecastHourResponse, err = m.getForecastNextHoursResponse()
		if err != nil {
			return nil, err
		}

		nextForecastHourFetch := time.Now().Add(meteoDataValidDuration)
		m.nextForecastHourFetch = &nextForecastHourFetch
	}
	return m.forecastHourResponse, nil
}

func (m *MeteoConceptClient) getForecastNextHoursResponse() (*ForecastHourResponse, error) {
	log.Println("getForecastNextHours ...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.meteo-concept.com/api/forecast/nextHours?token="+API_TOKEN+"&insee="+m.insee+"&hourly=true", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("API Response as string %+v\n", string(bodyBytes))

	if err != nil {
		return nil, err
	}
	var responseObject ForecastHourResponse
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	return &responseObject, nil
}

func (m *MeteoConceptClient) ForecastDaily() (*ForecastDayResponse, error) {

	if m.forecastDayResponse == nil || m.nextForecastDayFetch == nil || time.Now().After(*m.nextForecastDayFetch) {
		var err error
		m.forecastDayResponse, err = m.getForecastDailyResponse()
		if err != nil {
			return nil, err
		}
		nextForecastDayFetch := time.Now().Add(meteoDataValidDuration)
		m.nextForecastDayFetch = &nextForecastDayFetch
	}

	return m.forecastDayResponse, nil
}

func (m *MeteoConceptClient) getForecastDailyResponse() (*ForecastDayResponse, error) {
	log.Println("getForecastDaily ", m.insee)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.meteo-concept.com/api/forecast/daily?token="+API_TOKEN+"&insee="+m.insee, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("API Response as string %+v\n", string(bodyBytes))

	if err != nil {
		return nil, err
	}
	var responseObject ForecastDayResponse
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	return &responseObject, nil
}
