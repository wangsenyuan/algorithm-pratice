package p2894

const H = 30

const mod = 1_000_000_000 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func maxSum(nums []int, k int) int {
	//sort.Ints(nums)
	// a, b => a & b, a | b
	// x = a & b, y = a | b
	// a2 + b2 => x2 + y2 if a * b > x * y
	// 3, 6 = 2, 7 => 3 * 3 + 6 * 6 = 9 + 36 = 45
	// 4 + 49 = 54
	// 似乎是只要 a & b != a, 就应该合并
	// 如果 a & b
	cnt := make([]int, H)

	for _, num := range nums {
		for i := 0; i < H; i++ {
			cnt[i] += (num >> i) & 1
		}
	}

	var res int

	for k > 0 {
		k--
		var x int
		for i := 0; i < H; i++ {
			if cnt[i] > 0 {
				cnt[i]--
				x |= 1 << i
			}
		}
		res = add(res, mul(x, x))
	}
	return res
}
