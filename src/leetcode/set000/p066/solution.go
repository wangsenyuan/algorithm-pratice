package p066

func plusOne(digits []int) []int {
	n := len(digits)
	res := make([]int, n+1)
	carry := 1
	j := n
	for i := n - 1; i >= 0; i-- {
		x := digits[i] + carry
		if x < 10 {
			res[j] = x
			copy(res[1:j], digits[:i])
			return res[1:]
		}
		res[j] = x - 10
		carry = 1
		j--
	}
	if carry == 1 {
		res[0] = 1
		return res
	}
	return res[1:]
}
