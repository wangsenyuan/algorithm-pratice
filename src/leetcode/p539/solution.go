package main

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}

func findMinDifference(timePoints []string) int {
	if len(timePoints) <= 1 {
		return 0
	}

	parse := func(str string) int {
		hm := strings.Split(str, ":")
		hour, _ := strconv.Atoi(hm[0])
		min, _ := strconv.Atoi(hm[1])
		return hour*60 + min
	}

	time := make([]int, len(timePoints))

	for i := 0; i < len(timePoints); i++ {
		time[i] = parse(timePoints[i])
	}

	sort.Ints(time)

	var ans = -1

	for i := 1; i < len(time); i++ {
		if ans == -1 || time[i]-time[i-1] < ans {
			ans = time[i] - time[i-1]
		}
	}

	var x = time[0]
	for x < time[len(time)-1] {
		x += 24 * 60
	}
	if ans == -1 || x-time[len(time)-1] < ans {
		ans = x - time[len(time)-1]
	}

	return ans
}
