package p2446

func haveConflict(event1 []string, event2 []string) bool {
	a, b := parseTime(event1[0]), parseTime(event1[1])
	c, d := parseTime(event2[0]), parseTime(event2[1])

	return max(a, c) <= min(b, d)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func parseTime(s string) int {
	hour := int(s[0]-'0')*10 + int(s[1]-'0')
	min := int(s[3]-'0')*10 + int(s[4]-'0')
	return hour*60 + min
}
