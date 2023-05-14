package p2676

func doesValidArrayExist(derived []int) bool {
	n := len(derived)
	// d[i] = o[i] ^ o[(i + 1) % n]
	// 如果 o[0]= 1 or  o[0] = 0
	// o[0] ^ o[1] = d[0]
	// o[1] = o[0] ^ d[0]
	check := func(x int) bool {
		cur := x

		for i := 0; i < n; i++ {
			cur = cur ^ derived[i]
		}

		return cur == x
	}

	return check(0) || check(1)
}
