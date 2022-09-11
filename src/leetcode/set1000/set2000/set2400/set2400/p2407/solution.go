package p2407

const INF = 1 << 30
const N_INF = -INF

func lengthOfLIS(nums []int, k int) int {
	x := max_of_array(nums)
	x++

	arr := make([]int, 2*x)

	set := func(p int, v int) {
		p += x
		arr[p] = max(arr[p], v)
		for p > 0 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		res := N_INF

		l += x
		r += x

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for i := 1; i < len(arr); i++ {
		arr[i] = N_INF
	}

	set(0, 0)

	var best int

	n := len(nums)

	for i := 0; i < n; i++ {
		num := nums[i]
		tmp := get(max(num-k, 0), num)
		cur := max(1, tmp+1)
		best = max(best, cur)
		set(num, cur)
	}

	return best
}

func max_of_array(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
