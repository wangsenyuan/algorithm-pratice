package p1734

func decode(encoded []int) []int {
	n := len(encoded)
	var tot int

	for i := 1; i < n; i += 2 {
		tot ^= encoded[i]
	}

	n++

	all := ((n + 1) / 2) & 1

	res := make([]int, n)
	res[0] = all ^ tot
	for i := 1; i < n; i++ {
		res[i] = encoded[i-1] ^ res[i-1]
	}

	return res
}
