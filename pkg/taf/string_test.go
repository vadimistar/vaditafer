package taf

import (
	"testing"
	"time"
)

func TestWind_String(t *testing.T) {
	type fields struct {
		Direction string
		Speed     int
		Gusts     int
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				Direction: "",
				Speed:     0,
				Gusts:     0,
			},
			want: "штиль",
		},
		{
			fields: fields{
				Direction: "переменный",
				Speed:     1,
				Gusts:     0,
			},
			want: "переменный 1 м/c",
		},
		{
			fields: fields{
				Direction: "переменный",
				Speed:     1,
				Gusts:     6,
			},
			want: "переменный 1 м/c (порывы 6 м/c)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			w := &Wind{
				Direction: tt.fields.Direction,
				Speed:     tt.fields.Speed,
				Gusts:     tt.fields.Gusts,
			}
			if got := w.String(); got != tt.want {
				t.Errorf("Wind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudLayer_String(t *testing.T) {
	type fields struct {
		Quantity string
		Height   int
		Kind     string
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				Quantity: "сплошная",
				Height:   300,
				Kind:     "",
			},
			want: "сплошная, 300 м",
		},
		{
			fields: fields{
				Quantity: "небольшая",
				Height:   1200,
				Kind:     "кучево-дождевая",
			},
			want: "небольшая кучево-дождевая, 1200 м",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			c := &CloudLayer{
				Quantity: tt.fields.Quantity,
				Height:   tt.fields.Height,
				Kind:     tt.fields.Kind,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("CloudLayer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeHeader_String(t *testing.T) {
	type fields struct {
		Probability int
		Kind        string
		Start       time.Time
		End         time.Time
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				Probability: 40,
				Kind:        "",
				Start:       time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
				End:         time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: "с вероятностью 40% с 06:00 до 12:00",
		},
		{
			fields: fields{
				Probability: 40,
				Kind:        "временами",
				Start:       time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
				End:         time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: "с вероятностью 40% временами с 06:00 до 12:00",
		},
		{
			fields: fields{
				Kind:  "временами",
				Start: time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: "временами с 06:00 до 12:00",
		},
		{
			fields: fields{
				Start: time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
			},
			want: "с 06:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			c := &ChangeHeader{
				Probability: tt.fields.Probability,
				Kind:        tt.fields.Kind,
				Start:       tt.fields.Start,
				End:         tt.fields.End,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("ChangeHeader.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
