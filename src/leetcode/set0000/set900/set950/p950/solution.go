package p950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}
	sort.Ints(deck)
	res := make([]int, n)
	front, end := 0, n-1
	for _, card := range deck {
		res[index[front]] = card
		front = (front + 1) % n
		if front != end {
			end = (end + 1) % n
			index[end] = index[front]
			front = (front + 1) % n
		}
	}
	return res
}

func deckRevealedIncreasing1(deck []int) []int {
	var loop func(nums []int) []int

	loop = func(nums []int) []int {
		n := len(nums)
		if n == 0 {
			return nil
		}
		if n == 1 {
			return nums
		}
		var j int
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[j] {
				j = i
			}
		}
		first := nums[j]
		copy(nums[j:], nums[j+1:])
		tmp := loop(nums[:n-1])
		last := tmp[n-2]
		tmp = tmp[:n-2]
		return append([]int{first, last}, tmp...)
	}
	return loop(deck)
}
