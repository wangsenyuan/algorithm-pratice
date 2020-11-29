package p5615


func minMoves(nums []int, limit int) int {
	n := len(nums)
	bit := make(BIT, 2*limit+2)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		tmp0 := 2
		tmp1 := min(nums[i], nums[j])
		tmp2 := nums[i] + nums[j]
		tmp3 := max(nums[i], nums[j]) + limit
		tmp4 := 2 * limit
		// cnt[tmp] = 0
		// when sum is between res > tmp && res <= tmp
		bit.updateRange(tmp0, tmp1, 2)
		bit.updateRange(tmp1+1, tmp2-1, 1)
		bit.updateRange(tmp2+1, tmp3, 1)
		bit.updateRange(tmp3+1, tmp4, 2)
	}

	var best = n

	for i := 2; i <= 2*limit; i++ {
		best = min(best, bit.get(i))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type BIT []int

func (bit BIT) update(p int, v int) {
	n := len(bit)
	p++
	for p < n {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) updateRange(l, r int, v int) {
	if l > r {
		return
	}
	bit.update(l, v)
	bit.update(r+1, -v)
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}
