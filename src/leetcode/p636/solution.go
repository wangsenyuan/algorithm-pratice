package main

import (
	"strings"
	"strconv"
)

func exclusiveTime(n int, logs []string) []int {

}

func process(logs []string, fns []int) {
	id, tp, ts := parseLog(logs[0])

	if tp != "start" {
		panic("wrong assumption")
	}


}

func parseLog(log string) (int, string, int) {
	parts := strings.Split(log, ":")
	id, _ := strconv.Atoi(parts[0])
	ty := parts[1]
	ts, _ := strconv.Atoi(parts[2])
	return id, ty, ts
}
