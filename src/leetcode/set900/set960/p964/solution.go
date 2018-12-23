package p964

func leastOpsExpressTarget(x int, y int) int {
	var pos int
	var neg int
	var k int

	for y > 0 {
		cur := y % x
		y /= x
		if k > 0 {
			pos2 := min(cur*k+pos, (cur+1)*k+neg)
			neg2 := min((x-cur)*k+pos, (x-cur-1)*k+neg)
			pos, neg = pos2, neg2
		} else {
			pos = cur * 2
			neg = (x - cur) * 2
		}
		k++
	}

	return min(pos, k+neg) - 1
}

func leastOpsExpressTarget1(x int, target int) int {
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
