package p2375

func smallestNumber(pattern string) string {
	n := len(pattern)
	buf := make([]int, n+1)

	var check func(s int, i int, used int) bool

	check = func(s int, i int, used int) bool {
		buf[i] = s
		if i == n {
			return true
		}
		if pattern[i] == 'I' {
			for x := s + 1; x <= 9; x++ {
				if (used>>x)&1 == 0 && check(x, i+1, used|(1<<x)) {
					return true
				}
			}
			return false
		}
		for x := 1; x < s; x++ {
			if (used>>x)&1 == 0 && check(x, i+1, used|(1<<x)) {
				return true
			}
		}
		return false
	}

	for i := 1; i <= 9; i++ {
		if check(i, 0, 1<<i) {
			return toString(buf)
		}
	}

	return ""
}

func toString(arr []int) string {
	buf := make([]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		buf[i] = byte('0' + arr[i])
	}
	return string(buf)
}
