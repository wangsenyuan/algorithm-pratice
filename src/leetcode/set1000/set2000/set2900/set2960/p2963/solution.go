package p2963

func numberOfGoodPartitions(nums []int) int {
	// 对于一个数字x来说，它的最开始的数，和最后一个数，要在同一个分组内
	// 通过这个限制，可以把数组分成很多段，然后这些段中间，可以放隔板或者不放
	last := make(map[int]int)
	for i, x := range nums {
		last[x] = i
	}
	n := len(nums)
	var cnt int
	for i := 0; i < n; {
		r := last[nums[i]]
		for j := i; j < r; j++ {
			r = max(r, last[nums[j]])
		}
		cnt++
		i = r + 1
	}
	return pow(2, cnt-1)
}

const mod = 1000000007

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

func sub(a, b int) int {
	return add(a, mod-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}
