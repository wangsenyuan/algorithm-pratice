package p2751

func countBeautifulPairs(nums []int) int {
	cnt := make([]int, 10)
	var res int

	for _, num := range nums {
		x := num % 10
		for i := 1; i < 10; i++ {
			if gcd(i, x) == 1 {
				res += cnt[i]
			}
		}
		for num >= 10 {
			num /= 10
		}
		cnt[num]++
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
