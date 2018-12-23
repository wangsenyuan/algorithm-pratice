package p962

import "sort"

func maxWidthRamp(A []int) int {
	n := len(A)
	B := make([]Pair, n)
	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
	}
	sort.Sort(Pairs(B))

	m := n
	var ans int

	for _, p := range B {
		i := p.second
		ans = max(ans, i-m)
		m = min(m, i)
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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

func maxWidthRamp1(A []int) int {
	n := len(A)

	if n <= 1 {
		return 0
	}

	stack := make([]int, n)
	var p int
	var ans int
	for i := 0; i < n; i++ {
		if p > 0 {
			left, right := 0, p
			for left < right {
				mid := (left + right) / 2
				if A[stack[mid]] <= A[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			//A[left] <= A[i]
			if left < p {
				ans = max(ans, i-stack[left])
			}
		}

		if p == 0 || A[stack[p-1]] > A[i] {
			stack[p] = i
			p++
		}
	}

	return ans
}
