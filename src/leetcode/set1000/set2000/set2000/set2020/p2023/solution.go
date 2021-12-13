package p2023

import "strings"

func numOfPairs(nums []string, target string) int {

	var res int

	for i := 0; i < len(nums); i++ {
		if !strings.HasPrefix(target, nums[i]) {
			continue
		}
		for j := 0; j < len(nums); j++ {
			if i == j || len(nums[i])+len(nums[j]) != len(target) {
				continue
			}
			if strings.HasSuffix(target, nums[j]) {
				res++
			}
		}
	}
	return res
}
