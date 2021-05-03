package p1849

func splitString(s string) bool {
	n := len(s)

	var check func(i int, cur int) bool

	check = func(i int, cur int) bool {
		if i == n {
			return true
		}

		if cur == 1 {
			for i < n && s[i] == '0' {
				i++
			}
			return i == n
		}

		var tmp int
		for i < n && tmp != cur-1 {
			tmp = tmp*10 + int(s[i]-'0')
			i++
		}
		if i == n {
			return tmp == cur-1
		}
		return check(i, tmp)
	}

	var num int

	for i := 0; i < n-1; i++ {
		num = num*10 + int(s[i]-'0')
		if check(i+1, num) {
			return true
		}
	}
	return false
}
