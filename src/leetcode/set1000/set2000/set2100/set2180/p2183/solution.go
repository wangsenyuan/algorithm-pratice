package p2183

func coutPairs(nums []int, k int) int64 {
	n := len(nums)
	// n <= 100000
	// a * b % k == 0
	// a % k == 0
	// 假设num用k的除数表示的向量是, x, y, z
	X := nums[0]

	for i := 1; i < n; i++ {
		X = max(X, nums[i])
	}
	cnt := make([]int, X+1)
	for _, x := range nums {
		cnt[x]++
	}

	for i := 1; i <= X; i++ {
		for j := 2 * i; j <= X; j += i {
			cnt[i] += cnt[j]
		}
	}

	var res int64

	for _, x := range nums {
		y := k / gcd(x, k)
		if y <= X {
			res += int64(cnt[y])
			if int64(x)*int64(x)%int64(k) == 0 {
				res--
			}
		}
	}

	return res / 2
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
