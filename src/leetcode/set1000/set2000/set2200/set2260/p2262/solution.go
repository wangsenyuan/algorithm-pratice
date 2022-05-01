package p2262

func appealSum(s string) int64 {
	n := len(s)
	pos := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
	}

	L := make([]int, n)

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		L[i] = pos[x]
		pos[x] = i
	}
	for i := 0; i < 26; i++ {
		pos[i] = n
	}

	var res int64

	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		R := pos[x]
		tmp := int64(R-i) * int64(i-L[i])
		res += tmp
	}

	return res
}
