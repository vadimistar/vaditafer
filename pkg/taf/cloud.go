package taf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func cloudLayer(code string) (layer *CloudLayer, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid cloud layer group: %s", code)
	}()
	layer = new(CloudLayer)

	if len(code) < 5 {
		return nil, fmt.Errorf("invalid length of group (expect >= 6, got %d)", len(code))
	}

	if len(code) > 9 {
		return nil, fmt.Errorf("invalid length of group (expect <= 9, got %d)", len(code))
	}

	switch code[:3] {
	case "FEW":
		layer.Quantity = "небольшая"
	case "SCT":
		layer.Quantity = "рассеянная"
	case "BKN":
		layer.Quantity = "значительная"
	case "OVC", "VV0", "VV/":
		layer.Quantity = "сплошная"
	}

	var hh int
	if code[:2] == "VV" {
		if code[2:5] != "///" {
			hh, err = strconv.Atoi(code[2:5])
			if err != nil {
				return nil, fmt.Errorf("invalid vertical visibility: %s", code[2:5])
			}
		}
	} else {
		hh, err = strconv.Atoi(code[3:6])
		if err != nil {
			return nil, fmt.Errorf("invalid cloud layer height: %s", code[3:6])
		}
	}
	layer.Height = hh * 30

	if strings.HasSuffix(code, "CB") {
		layer.Kind = "кучево-дождевая"
	} else if strings.HasSuffix(code, "TCU") {
		layer.Kind = "кучевая"
	}

	return layer, nil
}
