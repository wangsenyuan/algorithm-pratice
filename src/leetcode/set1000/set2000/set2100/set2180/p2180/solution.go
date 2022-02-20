package p2180

func countEven(num int) int {
	var res int

	for i := 1; i <= num; i++ {
		tmp := i
		var sum int
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
		if sum%2 == 0 {
			res++
		}
	}
	return res
}
