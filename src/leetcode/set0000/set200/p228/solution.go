package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var process func(i int, from, end int)

	var res []string

	process = func(i int, from, end int) {
		if i == len(nums) {
			res = append(res, format(from, end))
			return
		}

		curr := nums[i]
		if curr-end > 1 {
			res = append(res, format(from, end))
			from = curr
		}

		process(i+1, from, curr)
	}

	process(1, nums[0], nums[0])

	return res
}

func format(from, end int) string {
	if from == end {
		return fmt.Sprintf("%d", from)
	}

	return fmt.Sprintf("%d->%d", from, end)
}
