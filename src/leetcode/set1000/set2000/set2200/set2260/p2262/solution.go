package p2262

func appealSum(s string) int64 {
	// n := len(s)

	var res int

	pos := make([]int, 26)
	var sum int

	for i, x := range []byte(s) {
		j := int(x - 'a')
		sum -= pos[j]
		pos[j] = i + 1
		sum += pos[j]
		res += sum
	}

	return int64(res)
}
