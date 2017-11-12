package main

import (
	"strings"
	"strconv"
	"fmt"
)

func exclusiveTime(n int, logs []string) []int {
	fns := make([]int, n)

	for i := 0; i < len(logs); i++ {
		j, _ := process(logs, i, fns)
		i = j
	}

	return fns
}

func process(logs []string, start int, fns []int) (int, int) {
	startId, _, startTs := parseLog(logs[start])

	took := 0

	for i := start + 1; i < len(logs); i++ {
		_, tp, ts := parseLog(logs[i])

		if tp == "start" {
			j, callTs := process(logs, i, fns)
			i = j
			took += callTs
		} else {
			execute := ts - startTs + 1 - took
			fns[startId] += execute
			return i, ts - startTs + 1
		}
	}
	panic("should not get here")
}

func parseLog(log string) (int, string, int) {
	parts := strings.Split(log, ":")
	id, _ := strconv.Atoi(parts[0])
	ty := parts[1]
	ts, _ := strconv.Atoi(parts[2])
	return id, ty, ts
}

func main() {
	/*logs := []string{
		"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6",
	}
	n := 2*/

	logs := []string{
		"0:start:0", "0:end:0", "1:start:1", "1:end:1", "2:start:2", "2:end:2", "2:start:3", "2:end:3", }
	n := 3
	fmt.Println(exclusiveTime(n, logs))
}
