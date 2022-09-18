package p2410

const H = 30

func smallestSubarrays(nums []int) []int {
	// i, j 最大， 且len[i, j] 最小
	// 需要知道每一位的最近的位置
	n := len(nums)

	cnt := make([][]int, n)

	for i := 0; i < n; i++ {
		cnt[i] = make([]int, H)
		if i > 0 {
			copy(cnt[i], cnt[i-1])
		}
		for j := 0; j < H; j++ {
			cnt[i][j] += (nums[i] >> j) & 1
		}
	}

	pos := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		pos[i] = make([]int, H)
		for j := 0; j < H; j++ {
			pos[i][j] = n
		}
	}

	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		res[i] = 1
		p := i
		num := nums[i]
		for j := H - 1; j >= 0 && p < n; j-- {
			x := cnt[p][j]
			if i > 0 {
				x -= cnt[i-1][j]
			}
			// we want x to be > 0
			if x == 0 {
				if pos[p+1][j] < n {
					p = pos[p+1][j]
					res[i] = p - i + 1
				}
			}
			// else  then it is ok to keep p no change
		}
		pos[i] = make([]int, H)

		for j := H - 1; j >= 0; j-- {
			x := (num >> j) & 1

			if x == 1 {
				pos[i][j] = i
			} else {
				pos[i][j] = pos[i+1][j]
			}
		}
	}

	return res
}
