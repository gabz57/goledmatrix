package meteo

import (
	"reflect"
	"testing"
	"time"
)

func TestMeteoConceptClient_ForecastDaily(t *testing.T) {
	type fields struct {
		insee                 string
		forecastHourResponse  *ForecastHourResponse
		forecastDayResponse   *ForecastDayResponse
		nextForecastHourFetch *time.Time
		nextForecastDayFetch  *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ForecastDayResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MeteoConceptClient{
				insee:                 tt.fields.insee,
				forecastHourResponse:  tt.fields.forecastHourResponse,
				forecastDayResponse:   tt.fields.forecastDayResponse,
				nextForecastHourFetch: tt.fields.nextForecastHourFetch,
				nextForecastDayFetch:  tt.fields.nextForecastDayFetch,
			}
			got, err := m.ForecastDaily()
			if (err != nil) != tt.wantErr {
				t.Errorf("ForecastDaily() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForecastDaily() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeteoConceptClient_ForecastNextHours(t *testing.T) {
	type fields struct {
		insee                 string
		forecastHourResponse  *ForecastHourResponse
		forecastDayResponse   *ForecastDayResponse
		nextForecastHourFetch *time.Time
		nextForecastDayFetch  *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ForecastHourResponse
		wantErr bool
	}{
		{"name", fields{"94016", nil, nil, nil, nil}, nil, false},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MeteoConceptClient{
				insee:                 tt.fields.insee,
				forecastHourResponse:  tt.fields.forecastHourResponse,
				forecastDayResponse:   tt.fields.forecastDayResponse,
				nextForecastHourFetch: tt.fields.nextForecastHourFetch,
				nextForecastDayFetch:  tt.fields.nextForecastDayFetch,
			}
			got, err := m.ForecastNextHours()
			if (err != nil) != tt.wantErr {
				t.Errorf("ForecastNextHours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForecastNextHours() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeteoConceptClient_getForecastDailyResponse(t *testing.T) {
	type fields struct {
		insee                 string
		forecastHourResponse  *ForecastHourResponse
		forecastDayResponse   *ForecastDayResponse
		nextForecastHourFetch *time.Time
		nextForecastDayFetch  *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ForecastDayResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MeteoConceptClient{
				insee:                 tt.fields.insee,
				forecastHourResponse:  tt.fields.forecastHourResponse,
				forecastDayResponse:   tt.fields.forecastDayResponse,
				nextForecastHourFetch: tt.fields.nextForecastHourFetch,
				nextForecastDayFetch:  tt.fields.nextForecastDayFetch,
			}
			got, err := m.getForecastDailyResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("getForecastDailyResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getForecastDailyResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeteoConceptClient_getForecastNextHoursResponse(t *testing.T) {
	type fields struct {
		insee                 string
		forecastHourResponse  *ForecastHourResponse
		forecastDayResponse   *ForecastDayResponse
		nextForecastHourFetch *time.Time
		nextForecastDayFetch  *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ForecastHourResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MeteoConceptClient{
				insee:                 tt.fields.insee,
				forecastHourResponse:  tt.fields.forecastHourResponse,
				forecastDayResponse:   tt.fields.forecastDayResponse,
				nextForecastHourFetch: tt.fields.nextForecastHourFetch,
				nextForecastDayFetch:  tt.fields.nextForecastDayFetch,
			}
			got, err := m.getForecastNextHoursResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("getForecastNextHoursResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getForecastNextHoursResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMeteoConceptClient(t *testing.T) {
	type args struct {
		insee string
	}
	tests := []struct {
		name string
		args args
		want *MeteoConceptClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMeteoConceptClient(tt.args.insee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMeteoConceptClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
