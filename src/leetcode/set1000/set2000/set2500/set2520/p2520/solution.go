package p2520

func countDigits(num int) int {
	var res int
	x := num

	var r int

	for x > 0 {
		x, r = x/10, x%10
		if num%r == 0 {
			res++
		}
	}

	return res
}
