package p2697

func punishmentNumber(n int) int {

	var res int

	for i := 1; i <= n; i++ {
		if check(i) {
			res += i * i
		}
	}

	return res
}

func check(i int) bool {
	num := i * i

	var dfs func(sum int, num int) bool

	dfs = func(sum int, num int) bool {
		if num == 0 {
			return sum == i
		}
		var base = 1
		var rem int
		for num > 0 {
			rem = (num%10)*base + rem
			num /= 10
			if dfs(sum+rem, num) {
				return true
			}
			base *= 10
		}
		return false
	}

	return dfs(0, num)
}
