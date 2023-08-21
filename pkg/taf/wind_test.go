package taf

import (
	"reflect"
	"testing"
)

func Test_direction(t *testing.T) {
	tests := []struct {
		code    string
		wantDir string
		wantErr bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "000",
			wantDir: "",
			wantErr: false,
		},
		{
			code:    "-100",
			wantDir: "",
			wantErr: true,
		},
		{
			code:    "045",
			wantDir: "северо-восточный",
			wantErr: false,
		},
		{
			code:    "360",
			wantDir: "северный",
			wantErr: false,
		},
		{
			code:    "VRB",
			wantDir: "переменный",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotDir, err := ParseWindDirection(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("direction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDir != tt.wantDir {
				t.Errorf("direction() = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}

func Test_wind(t *testing.T) {
	tests := []struct {
		code    string
		wantW   *Wind
		wantErr bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "VRB01BBVB",
			wantErr: true,
		},
		{
			code:    "VRB01G14BBVB",
			wantErr: true,
		},
		{
			code:    "VRB-2MPS",
			wantErr: true,
		},
		{
			code:    "VRB12G-2MPS",
			wantErr: true,
		},
		{
			code: "VRB02MPS",
			wantW: &Wind{
				Direction: "переменный",
				Speed:     2,
				Gusts:     0,
			},
		},
		{
			code: "VRB02G20KT",
			wantW: &Wind{
				Direction: "переменный",
				Speed:     1,
				Gusts:     10,
			},
		},
		{
			code: "VRB02G20MPS",
			wantW: &Wind{
				Direction: "переменный",
				Speed:     2,
				Gusts:     20,
			},
		},
		{
			code: "VRB02G20KMH",
			wantW: &Wind{
				Direction: "переменный",
				Speed:     1,
				Gusts:     6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotW, err := wind(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("wind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotW, tt.wantW) {
				t.Errorf("wind() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
