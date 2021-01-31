package p1742

func countBalls(lowLimit int, highLimit int) int {
	var best int
	mem := make(map[int]int)
	for x := lowLimit; x <= highLimit; x++ {
		y := sumDigits(x)
		mem[y]++
		if mem[y] > best {
			best = mem[y]
		}
	}
	return best
}

func sumDigits(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
