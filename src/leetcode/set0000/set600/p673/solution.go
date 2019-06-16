package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	tails := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		tails[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if tails[j]+1 > tails[i] {
					tails[i] = tails[j] + 1
					cnt[i] = cnt[j]
				} else if tails[j]+1 == tails[i] {
					cnt[i] += cnt[j]
				}
			}
		}

	}
	mls, ans := 0, 0

	for i, ls := range tails {
		if ls > mls {
			mls = ls
			ans = cnt[i]
		} else if ls == mls {
			ans += cnt[i]
		}
	}

	return ans
}

func main() {
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 3, 5}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))

}
