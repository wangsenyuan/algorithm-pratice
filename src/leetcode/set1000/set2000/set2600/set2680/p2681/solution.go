package p2681

import "sort"

const MOD = 1e9 + 7

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	// a[i] * a[i] * (a[1] * 2 ** (i - 2) + a[2] * 2 ** (i - 1) + .. + a[i] * 2 ** (i - (i-1))
	// a[i] * a[i] * (a[1] * 2 ** 1

	// i - (j + 1)
	// (i - 3)
	var sum int
	var res int

	for i := 1; i <= n; i++ {
		pw := pow(nums[i-1], 2)
		pw = mul(pw, pow(2, i))
		cur := mul(pw, sum)
		res = add(res, cur)
		res = add(res, pow(nums[i-1], 3))
		sum = add(sum, mul(nums[i-1], inverse(pow(2, i+1))))
	}

	return res
}
