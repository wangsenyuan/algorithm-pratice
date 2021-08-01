package p1954

const MOD = 1000000007

func countSpecialSubsequences(nums []int) int {
	n := len(nums)
	cnt2 := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		cnt2[i] = cnt2[i+1]
		if nums[i] == 2 {
			cnt2[i]++
		}
	}

	var ones int64
	var cnt0 int
	var res int64
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			// make this is the last 1,
			// no ones before
			tmp1 := modAdd(pow(2, cnt0), -1)
			// has some ones before
			tmp2 := ones
			cur := tmp1 + tmp2
			ones = modAdd(ones, cur)
			tmp3 := modAdd(pow(2, cnt2[i]), -1)
			res = modAdd(res, cur*tmp3%MOD)
		} else if nums[i] == 0 {
			cnt0++
		}
	}
	return int(res)
}

func modAdd(a, b int64) int64 {
	a += b
	a %= MOD
	if a < 0 {
		a += MOD
	}
	return a
}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}
