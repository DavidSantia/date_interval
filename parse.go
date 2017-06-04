package date_interval

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(interval string) (years, months, days int, err error) {
	var dFound, mFound, yFound bool

	d := []string{"days", "day", "d"}
	m := []string{"months", "month", "mo"}
	y := []string{"years", "year", "yr", "y"}
	all := append(d, m...)
	all = append(all, y...)
	allStr := strings.Join(all, ", ")

	// match overall pattern
	overallRgx, _ := regexp.Compile("^(-?\\d+ ?\\w+ ?){1,3}$")

	// split into fields
	splitRgx, _ := regexp.Compile("(-?\\d+ ?\\w+ ?)\\b")

	// split fields into value and unit
	fieldRgx, _ := regexp.Compile("(-?\\d+) ?(\\w+)")

	dRgx, _ := regexp.Compile("^" + strings.Join(d, "|") + "$")
	mRgx, _ := regexp.Compile("^" + strings.Join(m, "|") + "$")
	yRgx, _ := regexp.Compile("^" + strings.Join(y, "|") + "$")

	// initialize
	years = 0
	yFound = false
	months = 0
	mFound = false
	days = 0
	dFound = false

	s := strings.ToLower(strings.TrimSpace(interval))

	if !overallRgx.MatchString(s) {
		err = fmt.Errorf("Invalid interval %q, must be up to 3 numbers, each followed by a unit: %s",
			interval, allStr)
		return
	}

	// split into up to 3 fields
	fields := splitRgx.FindAllString(s, 3)

	for i := range fields {
		sub := fieldRgx.FindStringSubmatch(fields[i])
		if len(sub) != 3 {
			err = fmt.Errorf("Invalid interval %q, failed to parse field %q",
				interval, fields[i])
			return
		}

		if dRgx.MatchString(sub[2]) {
			if dFound {
				err = fmt.Errorf("Invalid interval %q, duplicate unit %q",
					interval, sub[2])
				return
			}
			days, err = strconv.Atoi(sub[1])
			dFound = true
		} else if mRgx.MatchString(sub[2]) {
			if mFound {
				err = fmt.Errorf("Invalid interval %q, duplicate unit %q",
					interval, sub[2])
				return
			}
			months, err = strconv.Atoi(sub[1])
			mFound = true
		} else if yRgx.MatchString(sub[2]) {
			if yFound {
				err = fmt.Errorf("Invalid interval %q, duplicate unit %q",
					interval, sub[2])
				return
			}
			years, err = strconv.Atoi(sub[1])
			yFound = true
		} else {
			err = fmt.Errorf("Invalid interval %q, failed to parse unit %q, must be on of: %s",
				interval, sub[2], allStr)
			return
		}

		if err != nil {
			err = fmt.Errorf("Invalid interval %q, failed to convert value %s",
				interval, sub[1])
		}
	}
	return
}
