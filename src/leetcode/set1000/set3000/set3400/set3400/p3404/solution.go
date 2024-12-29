package p3404

type pair struct {
	first  int
	second int
}

func numberOfSubsequences(nums []int) int64 {
	// n <= 1000
	// p * r = q * s
	// p + 1 < q, q + 1 < r, r + 1 < s
	freq := make(map[pair]int)
	n := len(nums)
	for q := 2; q < n-1; q++ {
		for p := 0; p+1 < q; p++ {
			a, b := nums[p], nums[q]
			c := gcd(a, b)
			a /= c
			b /= c
			freq[pair{a, b}]++
		}
	}

	var res int

	for r := n - 1; r > 0; r-- {
		// 先把和r-1上的删除掉
		for p := 0; p+1 < r-1; p++ {
			a, b := nums[p], nums[r-1]
			c := gcd(a, b)
			a /= c
			b /= c
			freq[pair{a, b}]--
		}
		// 考虑r的问题
		for s := r + 2; s < n; s++ {
			// a / b = d / c
			// c = nums[r]
			c := nums[r]
			d := nums[s]
			g := gcd(c, d)
			c /= g
			d /= g
			res += freq[pair{d, c}]
		}
	}

	return int64(res)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
