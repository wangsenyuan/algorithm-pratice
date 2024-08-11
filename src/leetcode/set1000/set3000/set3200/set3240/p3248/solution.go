package p3248

func finalPositionOfSnake(n int, commands []string) int {
	var x, y int
	for _, cmd := range commands {
		if cmd == "RIGHT" {
			y++
		} else if cmd == "DOWN" {
			x++
		} else if cmd == "LEFT" {
			y--
		} else {
			x--
		}
	}
	return x*n + y
}
