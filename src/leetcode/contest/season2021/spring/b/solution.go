package b

func orchestraLayout(n int, xPos int, yPos int) int {
	cycle := min(min(xPos, n-1-xPos), min(yPos, n-1-yPos))

	w := (4*int64(n) + 4*int64(n) - int64(cycle)*8) * int64(cycle) / 2

	if xPos != cycle {
		w += int64(n - cycle*2)
		if yPos != n-1-cycle {
			w += int64(n - cycle*2 - 1)
			if xPos != n-1-cycle {
				w += int64(n - cycle*2 - 1)
				w += int64(n - cycle - 1 - xPos)
			} else {
				w += int64(n - cycle - 1 - yPos)
			}
		} else {
			w += int64(xPos - cycle)
		}
	} else {
		w += int64(yPos - cycle + 1)
	}
	w %= 9
	if w == 0 {
		return 9
	}
	return int(w)
}

func mod9(num int) int {
	num %= 9
	if num == 0 {
		num = 9
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
