package p3142

func findPermutation(nums []int) []int {
	n := len(nums)
	m := n - 1
	T := 1 << m
	dp := make([][]*Item, T)
	for i := 0; i < T; i++ {
		dp[i] = make([]*Item, m)
		for j := 0; j < m; j++ {
			dp[i][j] = NewItem(m)
		}
	}

	for i := 0; i < m; i++ {
		dp[1<<i][i].val = nums[i+1]
		dp[1<<i][i].arr[0] = i + 1
	}

	for state := 1; state < T; state++ {
		var cnt int
		for i := 0; i < m; i++ {
			if (state>>i)&1 == 1 {
				cnt++
			}
		}
		if cnt == 1 {
			continue
		}
		for i := 0; i < m; i++ {
			if (state>>i)&1 == 0 {
				continue
			}

			for j := 0; j < m; j++ {
				if (state>>j)&1 == 0 {
					continue
				}
				it := dp[state^(1<<i)][j]
				tmp := it.val + abs(j+1-nums[i+1])

				if cnt == m {
					tmp += abs(i + 1 - nums[0])
				}
				if tmp >= oo {
					continue
				}
				if tmp < dp[state][i].val || tmp == dp[state][i].val && comp(it.arr, dp[state][i].arr, cnt) < 0 {
					dp[state][i].val = tmp
					copy(dp[state][i].arr, it.arr)
				}
			}

			dp[state][i].arr[cnt-1] = i + 1
		}
	}

	best := oo
	res := make([]int, n)

	for i := 0; i < m; i++ {
		if dp[T-1][i].val < best || dp[T-1][i].val == best && comp(dp[T-1][i].arr, res[1:], m) < 0 {
			best = dp[T-1][i].val
			copy(res[1:], dp[T-1][i].arr)
		}
	}
	return res
}

func comp(a, b []int, n int) int {
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	return 0
}

type Item struct {
	val int
	arr []int
}

func NewItem(n int) *Item {
	item := new(Item)
	item.val = oo
	item.arr = make([]int, n)
	return item
}

const oo = 1 << 60

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
