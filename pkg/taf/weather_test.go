package taf

import "testing"

func Test_weather(t *testing.T) {
	tests := []struct {
		code    string
		want    string
		wantErr bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "-",
			wantErr: true,
		},
		{
			code: "TS",
			want: "гроза",
		},
		{
			code: "-TSRA",
			want: "гроза, слабый дождь",
		},
		{
			code: "-TSRAGR",
			want: "гроза, слабый дождь, град",
		},
		{
			code: "FG",
			want: "туман",
		},
		{
			code: "BLSN",
			want: "метель(буря), снег",
		},
		{
			code: "DRSN",
			want: "поземок, снег",
		},
		{
			code: "+SHRA",
			want: "сильный ливневый дождь",
		},
		{
			code: "-FZDZ",
			want: "слабая замерзающая морось",
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			got, err := weather(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("weather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("weather() = %v, want %v", got, tt.want)
			}
		})
	}
}
