package p1930

func sumGame(num string) bool {
	// ??19
	var a, b int
	var left, right int
	n := len(num)
	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			a++
		} else {
			left += int(num[i] - '0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			b++
		} else {
			right += int(num[i] - '0')
		}
	}

	if (a+b)&1 == 1 {
		return true
	}

	left -= right
	b -= a
	if 2*left != 9*b {
		return true
	}
	return false
}
