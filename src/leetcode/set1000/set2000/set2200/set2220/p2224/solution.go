package p2224

func convertTime(current string, correct string) int {
	a := parseTime(current)
	c := parseTime(correct)

	x := (c - a) / 60
	c -= 60 * x

	y := (c - a) / 15
	c -= 15 * y

	u := (c - a) / 5
	c -= 5 * u

	v := c - a

	return x + y + u + v
}

func parseTime(s string) int {
	hour := int(s[0]-'0')*10 + int(s[1]-'0')
	min := int(s[3]-'0')*10 + int(s[4]-'0')
	return hour*60 + min
}
