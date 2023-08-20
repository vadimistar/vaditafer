package taf

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func createdAt(code string) (t time.Time, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid created at group: %s", code)
	}()

	if len(code) != 7 {
		return t, fmt.Errorf("invalid length of code group: %d", len(code))
	}

	if code[6] != 'Z' {
		return t, errors.New("no Z character at the end")
	}

	date, err := strconv.Atoi(code[:2])
	if err != nil {
		return t, fmt.Errorf("invalid date: %s", code[:2])
	}

	hour, err := strconv.Atoi(code[2:4])
	if err != nil {
		return t, fmt.Errorf("invalid hour: %s", code[2:4])
	}

	minute, err := strconv.Atoi(code[4:6])
	if err != nil {
		return t, fmt.Errorf("invalid minute: %s", code[4:6])
	}

	t = time.Now().UTC()
	return time.Date(t.Year(), t.Month(), date, hour, minute, 0, 0, time.UTC), nil
}

func period(code string) (from time.Time, to time.Time, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid forecast time group: %s", code)
	}()

	if len(code) != 9 {
		return from, to, fmt.Errorf("time group has to be 9 characters, got %d", len(code))
	}

	if code[4] != '/' {
		return from, to, errors.New("group[4] is not '/'")
	}

	from, err = datehour(code[:4])
	if err != nil {
		return from, to, fmt.Errorf("invalid from group: %s", from)
	}

	to, err = datehour(code[5:])
	if err != nil {
		return from, to, fmt.Errorf("invalid to group: %s", to)
	}

	return from, to, nil
}

func from(code string) (t time.Time, err error) {
	defer func() {
		err = errors.Wrapf(err, "invalid from time group: %s", code)
	}()

	if len(code) != 8 {
		return t, fmt.Errorf("invalid length, expect 8 characters, got %d", len(code))
	}

	if code[:2] != "FM" {
		return t, errors.New("first 2 characters are not FM")
	}

	dateAndHour, err := datehour(code[2:6])
	if err != nil {
		return t, errors.Wrap(err, "invalid date or(and) hour")
	}

	minutes, err := strconv.Atoi(code[6:])
	if err != nil {
		return t, fmt.Errorf("invalid minutes, expect integer: %s", code[6:])
	}

	t = time.Now().UTC()
	return time.Date(t.Year(), t.Month(), dateAndHour.Day(), dateAndHour.Hour(), minutes, 0, 0, time.UTC), nil
}

func datehour(code string) (t time.Time, err error) {
	if len(code) != 4 {
		return t, fmt.Errorf("invalid date hour (expect 4 characters): %s", code)
	}

	date, err := strconv.Atoi(code[:2])
	if err != nil {
		return t, errors.Wrapf(err, "invalid date: %s", code[:2])
	}

	hour, err := strconv.Atoi(code[2:])
	if err != nil {
		return t, errors.Wrapf(err, "invalid hour: %s", code[2:])
	}

	t = time.Now().UTC()
	return time.Date(t.Year(), t.Month(), date, hour, 0, 0, 0, time.UTC), nil
}
