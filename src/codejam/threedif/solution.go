package main

import (
	"fmt"
)

var M int64 = 1000000007

func main() {
	t := 0
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		t --
		nums := make([]int64, 3)

		fmt.Scanf("%d %d %d\n", &nums[0], &nums[1], &nums[2])

		sort(nums)

		x := nums[0] % M
		y := (nums[1] - 1) % M
		z := (nums[2] - 2) % M

		res := (((x * y) % M) * z) % M

		println(res)
	}
}

func sort(nums []int64) {
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		j := i;
		for j > 0 && nums[j - 1] > num {
			nums[j] = nums[j - 1]
			j--
		}
		nums[j] = num
	}
}
