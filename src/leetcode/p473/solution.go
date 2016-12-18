package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{1, 1, 2, 2, 2
	//nums := []int{3, 3, 3, 3, 4}
	nums := []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	if makesquare(nums) {
		fmt.Println("yes")
	} else {
		fmt.Println("no!!!!!")
	}
}

func makesquare(nums []int) bool {
	sum := int64(0)
	for _, num := range nums {
		sum += int64(num)
	}

	if sum == 0 || sum%4 != 0 {
		return false
	}

	side := sum / 4

	sort.Ints(nums)

	var check func(i int, flag int, sum int64) int
	check = func(i int, flag int, sum int64) int {
		if sum == side {
			return flag
		}

		if sum > side || i < 0 {
			return -1
		}

		for j := i - 1; j >= 0; j-- {
			if flag&(1<<uint(j)) > 0 {
				continue
			}
			ans := check(j, flag|(1<<uint(j)), sum+int64(nums[j]))
			if ans > 0 {
				return ans
			}
		}

		return -1
	}

	flag := 0
	for i := 0; i < 3; i++ {
		flag = check(len(nums), flag, 0)
		if flag < 0 {
			return false
		}
	}

	return true
}
