package p954

import "sort"

func canReorderDoubled(A []int) bool {
	n := len(A)
	if n == 0 {
		return true
	}
	B := make(map[int]int)
	for i := 0; i < n; i++ {
		B[A[i]]++
	}

	sort.Sort(XX(A))
	var cnt int
	for i := 0; i < n; i++ {
		if B[A[i]] == 0 {
			continue
		}
		B[A[i]]--
		y := 2 * A[i]
		if B[y] > 0 {
			B[y]--
			cnt += 2
		} else {
			return false
		}
	}
	return cnt == n
}

type XX []int

func (this XX) Len() int {
	return len(this)
}

func (this XX) Less(i, j int) bool {
	return abs(this[i]) < abs(this[j])
}

func (this XX) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
