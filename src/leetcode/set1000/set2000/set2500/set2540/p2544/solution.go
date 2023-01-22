package p2544

func alternateDigitSum(n int) int {
	var a int
	var b int
	var i int
	for n > 0 {
		if i%2 == 0 {
			a += n % 10
		} else {
			b += n % 10
		}
		n /= 10
		i++
	}
	i--
	if i%2 == 0 {
		return a - b
	}
	return b - a
}
