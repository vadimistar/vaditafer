package taf

import "testing"

func Test_visibility(t *testing.T) {
	tests := []struct {
		code    string
		wantVis int
		wantErr bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "0050",
			wantVis: 50,
		},
		{
			code:    "M0050",
			wantVis: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotVis, err := visibility(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("visibility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVis != tt.wantVis {
				t.Errorf("visibility() = %v, want %v", gotVis, tt.wantVis)
			}
		})
	}
}
