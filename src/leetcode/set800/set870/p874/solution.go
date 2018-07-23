package p874

func robotSim(commands []int, obstacles [][]int) int {
	grid := make(map[int]map[int]bool)

	for i := 0; i < len(obstacles); i++ {
		x, y := obstacles[i][0], obstacles[i][1]
		if grid[x] == nil {
			grid[x] = make(map[int]bool)
		}
		grid[x][y] = true
	}
	var best int
	// north
	var X, Y int
	dx, dy := 0, 1
	for _, cmd := range commands {
		if cmd == -2 {
			turnLeft(&dx, &dy)
		} else if cmd == -1 {
			turnRight(&dx, &dy)
		} else {
			for i := 0; i < cmd; i++ {
				x, y := X+dx, Y+dy
				var found bool
				if _, foundx := grid[x]; foundx {
					if _, foundy := grid[x][y]; foundy {
						found = true
					}
				}
				if found {
					break
				}
				X, Y = x, y
				tmp := dist(X, Y)
				if tmp > best {
					best = tmp
				}
			}
		}
	}

	return best
}

func dist(x, y int) int {
	return x*x + y*y
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func turnLeft(dx, dy *int) {
	if *dx == 0 && *dy == 1 {
		*dx = -1
		*dy = 0
	} else if *dx == -1 && *dy == 0 {
		*dx = 0
		*dy = -1
	} else if *dx == 0 && *dy == -1 {
		*dx = 1
		*dy = 0
	} else {
		*dx = 0
		*dy = 1
	}
}

func turnRight(dx, dy *int) {
	turnLeft(dx, dy)
	turnLeft(dx, dy)
	turnLeft(dx, dy)
}
