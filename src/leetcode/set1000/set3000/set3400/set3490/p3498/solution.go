package p3498

func reverseDegree(s string) int {
	var res int

	for i, x := range []byte(s) {
		j := int(x - 'a')
		j = 25 - j
		j++
		res += (i + 1) * j
	}
	return res
}
