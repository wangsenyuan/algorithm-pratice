package p2562

import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
	// 交换后，使得它们相等，
	// 假设目前i, A[i] != B[i], 设 A[i] < B[i]
	// 并交换它们, 交换后， x, y, x需要迁移， y需要后移
	// 在之前也是B中的最小值，目前x是最小值了， 并且因为x的加入，
	// 但是 x 肯定只需要匹配新的A[i] 即可
	sort.Ints(basket1)
	sort.Ints(basket2)

	x := min(basket1[0], basket2[0])

	A := diff(basket1, basket2)
	B := diff(basket2, basket1)
	// A是在1中，且不在2中的
	var res int64
	var a int
	for _, p := range A {
		if p.second%2 == 1 {
			return -1
		}
		a += p.second
	}
	var b int
	for _, p := range B {
		if p.second%2 == 1 {
			return -1
		}
		b += p.second
	}

	if a != b {
		return -1
	}

	// 需要交换 a/2次
	cnt := a / 2

	A = append(A, B...)

	sort.Slice(A, func(i, j int) bool {
		return A[i].first < A[j].first
	})

	for i := 0; i < len(A) && cnt > 0; i++ {
		if A[i].first > 2*x {
			break
		}
		res += int64(A[i].first) * int64(min(cnt, A[i].second/2))
		cnt -= min(cnt, A[i].second/2)
	}

	res += int64(cnt) * int64(2*x)
	return res
}

type Pair struct {
	first  int
	second int
}

func diff(a, b []int) []Pair {
	cnt := make(map[int]int)
	for _, num := range a {
		cnt[num]++
	}
	for _, num := range b {
		cnt[num]--
	}
	var res []Pair
	for k, v := range cnt {
		if v > 0 {
			res = append(res, Pair{k, v})
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
