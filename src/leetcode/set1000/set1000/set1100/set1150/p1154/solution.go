package p1154

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	ss := strings.Split(date, "-")
	year, _ := strconv.Atoi(ss[0])
	month, _ := strconv.Atoi(ss[1])
	day, _ := strconv.Atoi(ss[2])
	if month == 1 {
		return day
	}

	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeap(year) {
		months[1] = 29
	}

	var res int

	for i := 1; i < month; i++ {
		res += months[i-1]
	}

	return res + day
}

func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%4 == 0 {
		return year%100 != 0
	}

	return false
}
