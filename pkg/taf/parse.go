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

	for index < len(code) {
	}

	return nil, errors.New("todo")
}
