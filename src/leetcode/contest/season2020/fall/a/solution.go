package a

func calculate(s string) int {
	var x, y int64 = 1, 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}

	return int(x + y)
}
