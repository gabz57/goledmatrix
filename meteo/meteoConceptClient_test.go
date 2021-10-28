package meteo

import (
	"reflect"
	"testing"
)

func Test_getForecastNextHours(t *testing.T) {
	type args struct {
		insee string
	}
	tests := []struct {
		name    string
		args    args
		want    *ForecastHourResponse
		wantErr bool
	}{
		{"name", args{"94016"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getForecastNextHours(tt.args.insee)
			if (err != nil) != tt.wantErr {
				t.Errorf("getForecastNextHours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getForecastNextHours() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getForecastNextDaily(t *testing.T) {
	type args struct {
		insee string
	}
	tests := []struct {
		name    string
		args    args
		want    *ForecastDayResponse
		wantErr bool
	}{
		{"name", args{"94016"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getForecastDaily(tt.args.insee)
			if (err != nil) != tt.wantErr {
				t.Errorf("getForecastDaily() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getForecastDaily() got = %v, want %v", got, tt.want)
			}
		})
	}
}
