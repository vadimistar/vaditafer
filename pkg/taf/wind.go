package taf

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func wind(code string) (w *Wind, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid wind group: %s", code)
	}()

	w = new(Wind)

	if len(code) < 8 {
		return nil, fmt.Errorf("expect input size to be > 8: %s", code)
	}

	ddd := code[:3]
	ff := code[3:5]
	gg := ""
	if code[5] == 'G' {
		gg = code[6:8]
	}
	fmfm := code[len(code)-3:]
	if strings.HasSuffix(code, "KT") {
		fmfm = "KT"
	}

	w.Direction, err = ParseWindDirection(ddd)
	if err != nil {
		return nil, err
	}

	w.Speed, err = strconv.Atoi(ff)
	if err != nil {
		return nil, fmt.Errorf("invalid wind speed group: %s", ff)
	}
	if w.Speed < 0 {
		return nil, fmt.Errorf("invalid wind speed (negative): %d", w.Speed)
	}

	if gg != "" {
		w.Gusts, err = strconv.Atoi(gg)
		if err != nil {
			return nil, fmt.Errorf("invalid wind gusts group: %s", gg)
		}
		if w.Gusts < 0 {
			return nil, fmt.Errorf("invalid wind gusts (negative): %d", w.Gusts)
		}
	}

	switch fmfm {
	case "MPS":
		{
		}
	case "KT":
		w.Speed = int(math.Round(float64(w.Speed) * ktToMps))
		w.Gusts = int(math.Round(float64(w.Gusts) * ktToMps))
	case "KMH":
		w.Speed = int(math.Round(float64(w.Speed) * kmhToMps))
		w.Gusts = int(math.Round(float64(w.Gusts) * kmhToMps))
	default:
		return nil, fmt.Errorf("invalid wind speed/gusts units: %s", fmfm)
	}

	return w, nil
}

const ktToMps = 0.51
const kmhToMps = 0.28

func ParseWindDirection(code string) (dir string, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid wind direction group: %s", code)
	}()

	if code == "VRB" {
		return "переменный", nil
	}

	d, err := strconv.Atoi(code)
	if err != nil {
		return "", errors.Wrapf(err, "invalid wind direction: %s", code)
	}

	switch {
	case d == 0:
		return "", nil
	case (d > 315+23 && d <= 360) || (d > 0 && d <= 23):
		return "северный", nil
	case d > 23 && d <= 45+23:
		return "северо-восточный", nil
	case d > 45+23 && d <= 90+23:
		return "восточный", nil
	case d > 90+23 && d <= 135+23:
		return "юго-восточный", nil
	case d > 135+23 && d <= 180+23:
		return "южный", nil
	case d > 180+23 && d <= 225+23:
		return "юго-западный", nil
	case d > 225+23 && d <= 270+23:
		return "западный", nil
	case d > 270+23 && d <= 315+23:
		return "северо-западный", nil
	}

	return "", fmt.Errorf("invalid wind direction: %s (%d)", code, d)
}
