package taf

import (
	"errors"
	"fmt"
	"strings"
)

func ParseWeather(code string) (string, error) {
	if code == "NSW" {
		return "нет погодных явлений", nil
	}

	var (
		lightIntensity bool
		denseIntensity bool
		index          int
		showers        bool
		freezing       bool
	)

	if code == "" {
		return "", fmt.Errorf("invalid weather group (empty): %s", code)
	}

	switch code[0] {
	case '+':
		denseIntensity = true
		index = 1
	case '-':
		lightIntensity = true
		index = 1
	}

	var ww strings.Builder

	if index == len(code) {
		return "", errors.New("no weather")
	}

	for ; index < len(code); index += 2 {
		if index+2 > len(code) {
			return "", fmt.Errorf("invalid weather group (at index %d expect 2 characters): %s", index, code)
		}

		switch code[index : index+2] {
		case "VC":
			ww.WriteString("вблизи ")
		case "MI":
			ww.WriteString("поземный ")
		case "BC":
			ww.WriteString("обрывками ")
		case "PR":
			ww.WriteString("частичный ")
		case "DR":
			ww.WriteString("поземок, ")
		case "BL":
			ww.WriteString("метель(буря), ")
		case "SH":
			showers = true
		case "TS":
			ww.WriteString("гроза, ")
		case "FZ":
			freezing = true
		case "DZ":
			if lightIntensity {
				ww.WriteString("слабая ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильная ")
				denseIntensity = false
			}
			if freezing {
				ww.WriteString("замерзающая ")
				freezing = false
			}
			ww.WriteString("морось, ")
		case "RA":
			if lightIntensity {
				ww.WriteString("слабый ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильный ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневый ")
				showers = false
			}
			if freezing {
				ww.WriteString("замерзающий ")
				freezing = false
			}
			ww.WriteString("дождь, ")
		case "SN":
			if lightIntensity {
				ww.WriteString("слабый ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильный ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневый ")
				showers = false
			}
			ww.WriteString("снег, ")
		case "SG":
			if lightIntensity {
				ww.WriteString("слабые ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильные ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневые ")
				showers = false
			}
			ww.WriteString("снежные зерна, ")
		case "IC":
			ww.WriteString("ледяные иглы, ")
		case "PL":
			if lightIntensity {
				ww.WriteString("слабая ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильная ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневая ")
				showers = false
			}
			ww.WriteString("ледяная крупа, ")
		case "GR":
			if lightIntensity {
				ww.WriteString("слабый ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильный ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневый ")
				showers = false
			}
			ww.WriteString("град, ")
		case "GS":
			if lightIntensity {
				ww.WriteString("слабый ")
				lightIntensity = false
			}
			if denseIntensity {
				ww.WriteString("сильный ")
				denseIntensity = false
			}
			if showers {
				ww.WriteString("ливневый ")
				showers = false
			}
			ww.WriteString("мелкий град (снежная крупа), ")
		case "BR":
			ww.WriteString("дымка, ")
		case "FG":
			ww.WriteString("туман, ")
		case "FU":
			ww.WriteString("дым, ")
		case "VA":
			ww.WriteString("вукланический пепел, ")
		case "DU":
			ww.WriteString("пыль, ")
		case "SA":
			ww.WriteString("песок, ")
		case "HZ":
			ww.WriteString("мгла, ")
		case "PO":
			ww.WriteString("пыльные(песчаные) вихри, ")
		case "SQ":
			ww.WriteString("шквалы, ")
		case "FC":
			ww.WriteString("торнадо(смерч), ")
		case "SS":
			ww.WriteString("песчаная буря, ")
		case "DS":
			ww.WriteString("пыльная буря, ")
		}
	}

	return strings.TrimSuffix(ww.String(), ", "), nil
}
