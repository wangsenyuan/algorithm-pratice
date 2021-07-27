package p1945

func getLucky(s string, k int) int {
	var cur int
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a' + 1)
		if x < 10 {
			cur += x
		} else {
			cur += x/10 + x%10
		}
	}

	for k > 1 {
		var tmp int
		for cur > 0 {
			tmp += cur % 10
			cur /= 10
		}
		cur = tmp
		k--
	}

	return cur
}
