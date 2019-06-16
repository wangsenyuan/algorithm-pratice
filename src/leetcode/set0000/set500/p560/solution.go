package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	n := len(nums)

	as := make(map[int]int)
	as[0] = 1
	sum := 0
	res := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		a := sum - k
		if b, found := as[a]; found {
			res += b
		}
		as[sum]++
	}

	return res

}
