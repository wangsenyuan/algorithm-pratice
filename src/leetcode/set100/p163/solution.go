package main

import "fmt"

func main() {
	//nums := []int{0, 1, 3, 50, 75}
	//var nums []int
	nums := []int{-1}
	fmt.Println(findMissingRanges(nums, -1, -1))
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	var res []string

	from := lower - 1

	for _, num := range nums {
		if num < from {
			continue
		}
		if num-from > 2 {
			res = append(res, fmt.Sprintf("%d->%d", from+1, num-1))
		} else if num-from == 2 {
			res = append(res, fmt.Sprintf("%d", num-1))
		}
		from = num
	}

	if upper-from >= 2 {
		res = append(res, fmt.Sprintf("%d->%d", from+1, upper))
	} else if upper-from == 1 {
		res = append(res, fmt.Sprintf("%d", from+1))
	}

	return res
}
