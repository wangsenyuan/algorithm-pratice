package p3110

func scoreOfString(s string) int {
	var res int

	for i := 1; i < len(s); i++ {
		x := int(s[i] - 'a')
		y := int(s[i-1] - 'a')
		res += abs(x - y)
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
