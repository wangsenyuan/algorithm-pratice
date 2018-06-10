package p848

func shiftingLetters(S string, shifts []int) string {
	for i := 0; i < len(shifts); i++ {
		shifts[i] = shifts[i] % 26
	}
	n := len(S)

	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		x := int(S[i] - 'a')
		x = (x + shifts[i]) % 26
		res[i] = byte(x + 'a')
	}

	return string(res)
}
