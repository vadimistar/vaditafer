package taf

import "time"

type Taf struct {
	CreatedAt time.Time
	From      time.Time
	To        time.Time

	Forecasts []*Forecast
}

type Wind struct {
	Direction string
	Speed     int // in mps
	Gusts     int // in mps
}

type CloudLayer struct {
	Quantity string
	Height   int
	Kind     string
}

type Forecast struct {
	Header *ChangeHeader

	Wind        *Wind
	Visibility  int
	Weather     []string
	CloudLayers []*CloudLayer
}

type ChangeHeader struct {
	Probability int
	Kind        string
	Start       time.Time
	End         time.Time
}
