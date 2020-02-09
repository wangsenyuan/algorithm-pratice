package p1344

import "math"

func angleClock(hour int, minutes int) float64 {
	hour %= 12

	a := 6 * float64(minutes)

	b := 30*float64(hour) + a/12

	if a < b {
		return math.Min(b-a, a+360-b)
	}
	return math.Min(a-b, b+360-a)
}
