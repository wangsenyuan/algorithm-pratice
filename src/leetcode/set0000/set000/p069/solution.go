package p069

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var y = float64(x)
	for i := 0; i < 20; i++ {
		y = (y + float64(x)/y) / 2
	}
	return int(y)
}
