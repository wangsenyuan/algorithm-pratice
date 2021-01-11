package p1719

func decode(encoded []int, first int) []int {
	n := len(encoded)
	res := make([]int, n+1)

	res[0] = first

	for i := 1; i <= n; i++ {
		res[i] = encoded[i-1] ^ res[i-1]
	}

	return res
}
