package p3168

const H = 30

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	cnt := make([][]int, H)
	for i := 0; i < H; i++ {
		cnt[i] = make([]int, n+1)
	}

	getValue := func(l int, r int) int {
		var res int
		for i := 0; i < H; i++ {
			if cnt[i][r]-cnt[i][l] == r-l {
				res |= 1 << i
			}
		}
		return res
	}

	res := 1 << 50

	for i := 0; i < n; i++ {
		for j := 0; j < H; j++ {
			cnt[j][i+1] = cnt[j][i] + (nums[i]>>j)&1
		}

		l := search(i+1, func(j int) bool {
			return getValue(j, i+1) >= k
		})
		tmp := getValue(l, i+1)
		if tmp >= k {
			res = min(res, abs(tmp-k))
		}
		if l > 0 {
			tmp = getValue(l-1, i+1)
			res = min(res, abs(tmp-k))
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
