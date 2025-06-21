package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func MinutesBetween(t1, t2 string) (int, error) {
	h1m := strings.Split(t1, ":")
	h2m := strings.Split(t2, ":")

	h1, err1 := strconv.Atoi(h1m[0])
	m1, err2 := strconv.Atoi(h1m[1])
	h2, err3 := strconv.Atoi(h2m[0])
	m2, err4 := strconv.Atoi(h2m[1])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return 0, fmt.Errorf("invalid time format '%s' or '%s'", t1, t2)
	}

	total1 := h1*60 + m1
	total2 := h2*60 + m2

	return total2 - total1, nil
}
