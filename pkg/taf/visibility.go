package taf

import (
	"strconv"

	"github.com/pkg/errors"
)

func visibility(code string) (vis int, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid visibility group: %s", code)
	}()

	if code == "" {
		return 0, errors.New("input is empty")
	}

	if code[0] == 'M' {
		vis, err = strconv.Atoi(code[1:])
		if err != nil {
			return 0, errors.Wrapf(err, "invalid visibility: %s", code[1:])
		}
		return vis, nil
	}

	vis, err = strconv.Atoi(code)
	if err != nil {
		return 0, errors.Wrapf(err, "invalid visibility: %s", code[1:])
	}
	return vis, nil
}
