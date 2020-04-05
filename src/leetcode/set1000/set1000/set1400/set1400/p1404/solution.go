package p1404

func numSteps(s string) int {
	var res int

	n := len(s)

	buf := []byte(s)

	for i := n - 1; i > 0; i-- {
		if buf[i] == '0' {
			res++
			continue
		}
		var j = i

		for j > 0 && buf[j] == '1' {
			j--
		}

		//011
		res++
		res += (i - j)
		if j == 0 {
			// buf[0] = 1 and carry 1 here
			res++
			break
		}
		buf[j] = '1'
		i = j + 1
	}

	return res
}
