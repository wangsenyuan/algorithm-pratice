package p2831

func longestEqualSubarray(nums []int, k int) int {
	// 对于[l..r] 删除最多k个元素后，剩余的数都是相同的
	// (r - l + 1) - cnt[nums[r]][r] - cnt[nums[r]][l]  <= k
	n := len(nums)
	cnt := make([]int, n+1)
	pos := make([][]Pair, n+1)
	id := make([]int, n+1)
	var ans int = 1
	for r, num := range nums {
		cnt[num]++
		for id[num] < len(pos[num]) {
			cur := pos[num][id[num]]
			l := cur.first
			if r-l+1-(cnt[num]-cur.second+1) <= k {
				ans = max(ans, cnt[num]-cur.second+1)
				break
			}
			id[num]++
		}
		if len(pos[num]) == 0 {
			pos[num] = make([]Pair, 0, 1)
		}
		pos[num] = append(pos[num], Pair{r, cnt[num]})
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
