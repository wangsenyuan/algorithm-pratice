package p2408

import "strconv"

var month_days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	a := getDays(arriveAlice)
	b := getDays(leaveAlice)
	c := getDays(arriveBob)
	d := getDays(leaveBob)

	x := max(a, c)
	y := min(b, d)
	if x > y {
		return 0
	}
	return y - x + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func getDays(s string) int {
	m := s[:2]
	d := s[3:]

	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)

	var days int

	for i := 0; i < month-1; i++ {
		days += month_days[i]
	}

	return days + day
}
