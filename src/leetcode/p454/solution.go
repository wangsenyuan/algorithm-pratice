package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}

	fmt.Println(fourSumCount(A, B, C, D))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	X := product(A, B)
	Y := product(C, D)

	sort.Ints(X)
	sort.Ints(Y)

	ans := 0
	last := 0
	for i, x := range X {
		if i == 0 || X[i] != X[i-1] {
			last = count(Y, -x)
		}
		ans += last
	}

	return ans
}

func count(nums []int, val int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] > val
	})

	j := sort.Search(len(nums), func(j int) bool {
		return nums[j] >= val
	})

	return i - j
}

func product(A, B []int) []int {
	m := len(A)
	n := len(B)
	X := make([]int, m*n)

	for i, a := range A {
		for j, b := range B {
			X[i*n+j] = a + b
		}
	}
	return X
}
