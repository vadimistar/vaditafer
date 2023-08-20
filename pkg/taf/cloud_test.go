package taf

import (
	"reflect"
	"testing"
)

func Test_cloudLayer(t *testing.T) {
	tests := []struct {
		code      string
		wantLayer *CloudLayer
		wantErr   bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code: "VV003",
			wantLayer: &CloudLayer{
				Quantity: "сплошная",
				Height:   90,
			},
		},
		{
			code: "VV///",
			wantLayer: &CloudLayer{
				Quantity: "сплошная",
				Height:   0,
			},
		},
		{
			code: "FEW020CB",
			wantLayer: &CloudLayer{
				Quantity: "небольшая",
				Height:   600,
				Kind:     "кучево-дождевая",
			},
		},
		{
			code: "SCT030",
			wantLayer: &CloudLayer{
				Quantity: "рассеянная",
				Height:   900,
			},
		},
		{
			code: "BKN040TCU",
			wantLayer: &CloudLayer{
				Quantity: "значительная",
				Height:   1200,
				Kind:     "кучевая",
			},
		},
		{
			code: "OVC010",
			wantLayer: &CloudLayer{
				Quantity: "сплошная",
				Height:   300,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotLayer, err := cloudLayer(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("cloudLayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLayer, tt.wantLayer) {
				t.Errorf("cloudLayer() = %v, want %v", gotLayer, tt.wantLayer)
			}
		})
	}
}
