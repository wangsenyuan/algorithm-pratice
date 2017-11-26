package p735

func asteroidCollision(asteroids []int) []int {
	n := len(asteroids)

	flag := make([]bool, n)

	stack := make([]int, n)
	p := 0
	for i := 0; i < n; i++ {
		if asteroids[i] > 0 {
			stack[p] = i
			p++
			continue
		}
		for p > 0 && asteroids[stack[p-1]] < -asteroids[i] {
			flag[stack[p-1]] = true
			p--
		}

		if p > 0 && asteroids[stack[p-1]] == -asteroids[i] {
			flag[stack[p-1]] = true
			p--
			flag[i] = true
			continue
		}

		if p > 0 {
			flag[i] = true
		}
	}

	res := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if !flag[i] {
			res = append(res, asteroids[i])
		}
	}
	return res
}
