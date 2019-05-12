package p1041

func isRobotBounded(instructions string) bool {

	var dir int
	var x, y int

	n := len(instructions)
	for i := 0; i < 100*n; i++ {
		cur := instructions[i%n]
		if cur == 'G' {
			move(&x, &y, dir)
		} else if cur == 'L' {
			turnLeft(&dir)
		} else {
			turnRight(&dir)
		}
		if x == 0 && y == 0 && dir == 0 && (i + 1) % n == 0 {
			return true
		}
	}

	return false
}

func move(x, y *int, dir int) {
	if dir == 0 {
		*y++
	} else if dir == 1 {
		*x++
	} else if dir == 2 {
		*y--
	} else {
		*x--
	}
}

func turnLeft(dir *int) {
	if *dir == 0 {
		// go west
		*dir = 3
	} else if *dir == 1 {
		*dir = 0
	} else if *dir == 2 {
		*dir = 1
	} else {
		*dir = 2
	}
}

func turnRight(dir *int) {
	if *dir == 0 {
		// go west
		*dir = 1
	} else if *dir == 1 {
		*dir = 2
	} else if *dir == 2 {
		*dir = 3
	} else {
		*dir = 0
	}
}
