package p1617

func minimumOneBitOperations(num int) int {

	mem := make(map[int]int)

	var loop func(num int) int

	loop = func(num int) int {
		if num == 0 {
			return 0
		}
		if num == 1 {
			return 1
		}
		if num == 2 {
			return 3
		}
		if num == 3 {
			return 2
		}
		if v, found := mem[num]; found {
			return v
		}
		i := 30
		for num&(1<<i) == 0 {
			i--
		}
		var res int
		// num & (1 << i) > 0
		if num&(1<<(i-1)) > 0 {
			// 11000
			res += loop(num & (1<<(i-1) - 1))
			res += loop(1 << (i - 1))
			res++
		} else {
			res += loop(1 << (i - 1))
			res -= loop(num & (1<<(i-1) - 1))
			res += loop(1<<i | 1<<(i-1))
		}
		mem[num] = res
		return res
	}

	return loop(num)
}
