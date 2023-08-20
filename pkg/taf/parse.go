package taf

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Parse(code string) (taf *Taf, err error) {
	defer func() {
		if recErr := recover(); recErr != nil {
			err = fmt.Errorf("unexpected error while parsing: %v", recErr)
		}

		err = errors.Wrapf(err, "cannot parse taf code: %s", code)
	}()
	taf = new(Taf)

	groups := strings.Fields(code)

	var index int

	if groups[index] == "TAF" {
		index++
	}
	if groups[index] == "COR" {
		index++
	}
	if groups[index] == "AMD" {
		index++
	}

	index++ // icao

	taf.CreatedAt, err = createdAt(groups[index])
	if err != nil {
		return nil, err
	}
	index++

	if groups[index] == "NIL" {
		return nil, errors.New("empty forecast: NIL")
	}

	taf.From, taf.To, err = period(groups[index])
	if err != nil {
		return nil, err
	}
	index++

	if groups[index] == "CNL" {
		return nil, errors.New("empty forecast: CNL")
	}

	taf.Wind, err = wind(groups[index])
	if err != nil {
		return nil, err
	}
	index++

	if groups[index] == "CAVOK" {
		taf.Visibility = 9999
		index++
	} else {
		taf.Visibility, err = visibility(groups[index])
		if err != nil {
			return nil, err
		}
		index++

		for {
			w, err := weather(groups[index])
			if err != nil {
				break
			}
			taf.Weather = append(taf.Weather, w)
			index++
		}

		for {
			c, err := cloudLayer(groups[index])
			if err != nil {
				break
			}
			taf.CloudLayers = append(taf.CloudLayers, c)
			index++
		}
	}

	for index < len(code) {
	}

	return nil, errors.New("todo")
}
