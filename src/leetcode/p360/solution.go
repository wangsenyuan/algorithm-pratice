package main

import "fmt"

func main() {
	nums := []int{-4, -2, 2, 4}
	a, b, c := 0, 3, 5
	result := sortTransformedArray(nums, a, b, c)
	fmt.Printf("%v", result)
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	if a == 0 && b == 0 {
		return transform(nums, 0, len(nums), a, b, c, true)
	}

	if a == 0 && b > 0 {
		return transform(nums, 0, len(nums), a, b, c, true)
	}

	if a == 0 && b < 0 {
		return transform(nums, 0, len(nums), a, b, c, false)
	}

	i, j := 0, len(nums)-1
	x := -1.0 * b / 2 / a
	for i <= j {
		k := (i + j) / 2
		if nums[k] < x {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	asc := a < 0

	as := transform(nums, 0, i, a, b, c, asc)
	bs := transform(nums, i, len(nums), a, b, c, !asc)
	return merge(as, bs)
}

func transform(nums []int, start, end, a, b, c int, asc bool) []int {
	result := make([]int, end-start, end-start)

	for i := 0; i < end-start; i++ {
		j := start + i
		if !asc {
			j = end - 1 - i
		}
		result[i] = nums[j]*(nums[j]*a+b) + c
	}

	return result
}

func merge(as, bs []int) []int {
	cs := make([]int, len(as)+len(bs), len(as)+len(bs))

	i, j, k := 0, 0, 0
	for i < len(as) && j < len(bs) {
		if as[i] < bs[j] {
			cs[k] = as[i]
			k, i = k+1, i+1
		} else {
			cs[k] = bs[j]
			k, j = k+1, j+1
		}
	}

	for i < len(as) {
		cs[k] = as[i]
		k, i = k+1, i+1
	}

	for j < len(bs) {
		cs[k] = bs[j]
		k, j = k+1, j+1
	}
	return cs
}
