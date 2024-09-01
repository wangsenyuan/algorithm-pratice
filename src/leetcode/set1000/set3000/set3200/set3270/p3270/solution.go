package p3270

func generateKey(num1 int, num2 int, num3 int) int {
	var res int
	base := 1
	for num1 > 0 || num2 > 0 || num3 > 0 {
		x := min(num1%10, num2%10, num3%10)
		res += x * base
		base *= 10
		num1 /= 10
		num2 /= 10
		num3 /= 10
	}
	return res
}
