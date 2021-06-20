package p1904

import (
	"strconv"
	"strings"
)

func numberOfRounds(startTime string, finishTime string) int {
	start := parse(startTime)
	finish := parse(finishTime)

	if finish < start {
		finish += 24 * 60
	}
	start = (start + 14) / 15 * 15
	duration := finish - start
	return duration / 15
}

func parse(s string) int {
	ss := strings.Split(s, ":")
	h, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])
	return h*60 + m
}
