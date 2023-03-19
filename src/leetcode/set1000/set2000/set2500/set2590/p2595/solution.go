package p2595

func evenOddBit(n int) []int {
	res := make([]int, 2)
	var i int
	for n > 0 {
		res[i%2] += n & 1
		n >>= 1
		i++
	}

	return res
}
