package main

import (
	"strconv"
	"time"
)

func getPaddedString(str string, padding int) (out string) {
	switch {
	case len(str) > padding:
		// Truncate everything after a length of 12
		out = str[:padding]
	case len(str) == padding:
		// String is perfect as-is, nothing needs to be done
		out = str
	default:
		delta := padding - len([]rune(str))
		prefix := delta / 2
		suffix := delta - prefix
		out = getSpaces(prefix) + str + getSpaces(suffix)
	}

	return
}

func getSpaces(v int) (out string) {
	buf := make([]rune, 0, v)
	for i := 0; i < v; i++ {
		buf = append(buf, ' ')
	}

	out = string(buf)
	return
}

func getDurationString(dur time.Duration) (out string) {
	str := getPaddedString(dur.String(), 12)

	switch {
	case dur < time.Millisecond*100:
		return cyan.Sprint(str)

	case dur < time.Millisecond*500:
		return green.Sprint(str)

	case dur < time.Second:
		return yellow.Sprint(str)

	default:
		return red.Sprint(str)
	}
}

func getStatusString(status int) (out string) {
	str := " " + strconv.Itoa(status) + " "

	switch {
	case status < 400:
		return bgGreen.Sprint(str)
	case status < 500:
		return bgYellow.Sprint(str)
	default:
		return bgRed.Sprint(str)
	}
}
