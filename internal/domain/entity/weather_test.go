package entity

import (
	"testing"
)

func TestNewWeather(t *testing.T) {
	tests := []struct {
		name  string
		tempC float64
		wantC float64
		wantF float64
		wantK float64
	}{
		{
			name:  "freezing point",
			tempC: 0,
			wantC: 0,
			wantF: 32,
			wantK: 273,
		},
		{
			name:  "boiling point",
			tempC: 100,
			wantC: 100,
			wantF: 212,
			wantK: 373,
		},
		{
			name:  "room temperature",
			tempC: 25,
			wantC: 25,
			wantF: 77,
			wantK: 298,
		},
		{
			name:  "negative temperature",
			tempC: -40,
			wantC: -40,
			wantF: -40,
			wantK: 233,
		},
		{
			name:  "decimal temperature",
			tempC: 28.5,
			wantC: 28.5,
			wantF: 83.3,
			wantK: 301.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weather := NewWeather(tt.tempC)

			if weather.TempC != tt.wantC {
				t.Errorf("TempC = %v, want %v", weather.TempC, tt.wantC)
			}
			if weather.TempF != tt.wantF {
				t.Errorf("TempF = %v, want %v", weather.TempF, tt.wantF)
			}
			if weather.TempK != tt.wantK {
				t.Errorf("TempK = %v, want %v", weather.TempK, tt.wantK)
			}
		})
	}
}
