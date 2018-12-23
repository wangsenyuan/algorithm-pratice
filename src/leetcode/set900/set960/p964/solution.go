package p964

func leastOpsExpressTarget(x int, target int) int {
	mem := make(map[int64]int)
	X := int64(x)
	var loop func(target int64) int

	loop = func(target int64) int {
		if v, found := mem[target]; found {
			return v
		}
		if target == X {
			return 0
		}
		if target == 1 {
			return 1
		}

		y := int64(1)
		var b int

		for y < target {
			y *= X
			b++
		}
		// X ** b >= targget
		if y == target {
			return b - 1
		}
		a := loop(target - y/X)
		if b == 1 {
			a += 2
		} else {
			a += b - 1
		}
		if y-target < target {
			a = min(a, loop(y-target)+b)
		}

		mem[target] = a
		return a
	}

	return loop(int64(target))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
