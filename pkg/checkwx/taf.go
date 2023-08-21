package checkwx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/vadimistar/vaditafer/pkg/taf"
)

func (c *Client) Taf(id string) (t *taf.Taf, err error) {
	defer func() {
		err = errors.Wrapf(err, "client has an endpoint: %s", c.ApiEndpoint)
	}()

	if !idRegex.MatchString(id) {
		return nil, fmt.Errorf("expect id to be 4 characters long and consist only of english letters: %s", id)
	}

	url, err := url.Parse(c.ApiEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "decode client endpoint")
	}
	url.Path = fmt.Sprintf("taf/%s/decoded", id)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create request")
	}

	req.Header.Add("X-API-Key", c.ApiKey)

	rs, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "cannot do request")
	}

	defer rs.Body.Close()

	var resp response

	if err := json.NewDecoder(rs.Body).Decode(&resp); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	if len(resp.Data) <= 0 {
		return nil, errors.Wrap(err, "no data")
	}

	t, err = parseResponse(&resp.Data[0])
	if err != nil {
		return nil, err
	}

	return t, nil
}

var idRegex = regexp.MustCompile("^[A-Z]{4}$")

func parseResponse(resp *data) (t *taf.Taf, err error) {
	defer func() {
		err = errors.Wrapf(err, "parse api response: %+v", resp)
	}()
	t = new(taf.Taf)

	t.CreatedAt, err = time.Parse(time.RFC3339, resp.Timestamp.Issued)
	if err != nil {
		return nil, errors.Wrap(err, "parse created at")
	}

	t.From, err = time.Parse(time.RFC3339, resp.Timestamp.From)
	if err != nil {
		return nil, errors.Wrap(err, "parse from")
	}

	t.To, err = time.Parse(time.RFC3339, resp.Timestamp.To)
	if err != nil {
		return nil, errors.Wrap(err, "parse to")
	}

	t.Forecasts = make([]*taf.Forecast, len(resp.Forecast))
	for i, forecast := range resp.Forecast {
		if forecast.Ceiling.Text != "" {
			cloudLayer := parseCloudLayer(forecast.Ceiling.BaseMetersAgl, forecast.Ceiling.Code)
			t.Forecasts[i].CloudLayers = append(t.Forecasts[i].CloudLayers, cloudLayer)
		}

		t.Forecasts[i].Header, err = parseHeader(forecast.Change.Indicator.Code, forecast.Timestamp.From, forecast.Timestamp.To)
		if err != nil {
			return nil, err
		}

		t.Forecasts[i].Weather = make([]string, len(forecast.Conditions))
		for j, cond := range forecast.Conditions {
			ww, err := taf.ParseWeather(cond.Prefix + cond.Code)
			if err != nil {
				return nil, err
			}
			t.Forecasts[i].Weather[j] = ww
		}

		if forecast.Visibility.Meters != "" {
			t.Forecasts[i].Visibility = int(forecast.Visibility.MetersFloat)
		}

		if forecast.Wind.Degrees != 0 || forecast.Wind.SpeedKph != 0 {
			t.Forecasts[i].Wind, err = parseWind(forecast.Wind.Degrees, forecast.Wind.SpeedMps, forecast.Wind.GustMps)
			if err != nil {
				return nil, err
			}
		}
	}

	return t, nil
}

func parseWind(dir int, speed int, gusts int) (w *taf.Wind, err error) {
	defer func() {
		err = errors.Wrapf(err, "parse wind dir = %d speed = %d gusts = %d", dir, speed, gusts)
	}()
	w = new(taf.Wind)

	if dir != 0 {
		w.Direction, err = taf.ParseWindDirection(fmt.Sprint(dir))
		if err != nil {
			return nil, err
		}
	}

	w.Speed = speed
	if gusts != 0 {
		w.Gusts = gusts
	}

	return w, nil
}

func parseHeader(code string, from string, to string) (header *taf.ChangeHeader, err error) {
	defer func() {
		err = errors.Wrapf(err, "parse change header: %s", code)
	}()

	switch code {
	case "FM":
		header.Kind = ""
	case "TEMPO":
		header.Kind = "временами"
	case "BECMG":
		header.Kind = "изменение погоды"
	}

	header.Start, err = time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, errors.Wrap(err, "parse from")
	}

	header.End, err = time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, errors.Wrap(err, "parse to")
	}

	return header, nil
}

func parseCloudLayer(base int, code string) *taf.CloudLayer {
	var quantity string

	switch code {
	case "FEW":
		quantity = "небольшая"
	case "SCT":
		quantity = "рассеянная"
	case "BKN":
		quantity = "значительная"
	case "OVC", "OVX":
		quantity = "сплошная"
	}

	return &taf.CloudLayer{
		Quantity: quantity,
		Height:   base,
	}
}

type response struct {
	Data    []data `json:"data"`
	Results int    `json:"results"`
}

type data struct {
	Forecast []struct {
		Ceiling struct {
			BaseFeetAgl   int    `json:"base_feet_agl"`
			BaseMetersAgl int    `json:"base_meters_agl"`
			Code          string `json:"code"`
			Feet          int    `json:"feet"`
			Meters        int    `json:"meters"`
			Text          string `json:"text"`
		} `json:"ceiling,omitempty"`
		Clouds []struct {
			BaseFeetAgl   int    `json:"base_feet_agl"`
			BaseMetersAgl int    `json:"base_meters_agl"`
			Code          string `json:"code"`
			Feet          int    `json:"feet"`
			Meters        int    `json:"meters"`
			Text          string `json:"text"`
		} `json:"clouds"`
		Conditions []struct {
			Code   string `json:"code"`
			Prefix string `json:"prefix"`
			Text   string `json:"text"`
		} `json:"conditions"`
		Timestamp struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"timestamp"`
		Visibility struct {
			Meters      string  `json:"meters"`
			MetersFloat float64 `json:"meters_float"`
			Miles       string  `json:"miles"`
			MilesFloat  float64 `json:"miles_float"`
		} `json:"visibility,omitempty"`
		Wind struct {
			Degrees  int `json:"degrees"`
			SpeedKph int `json:"speed_kph"`
			SpeedKts int `json:"speed_kts"`
			SpeedMph int `json:"speed_mph"`
			SpeedMps int `json:"speed_mps"`
			GustMps  int `json:"gust_mps"`
		} `json:"wind,omitempty"`
		Change struct {
			Indicator struct {
				Code string `json:"code"`
				Desc string `json:"desc"`
				Text string `json:"text"`
			} `json:"indicator"`
		} `json:"change,omitempty"`
	} `json:"forecast"`
	Icao    string `json:"icao"`
	RawText string `json:"raw_text"`
	Station struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Location string `json:"location"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	} `json:"station"`
	Timestamp struct {
		Bulletin string `json:"bulletin"`
		From     string `json:"from"`
		Issued   string `json:"issued"`
		To       string `json:"to"`
	} `json:"timestamp"`
}
