package main

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}

	avg := sum / k

	for _, num := range nums {
		if num > avg {
			return false
		}
	}

	n := len(nums)

	var check func(i int, k int, left int, flag int) bool

	check = func(i int, k int, left int, flag int) bool {
		if left < 0 {
			return false
		}

		if flag == 0 {
			return true
		}

		if i == n {
			return false
		}

		for j := k; j < n; j++ {
			if nums[j] > left || flag>>uint(j)&1 == 0 {
				continue
			}
			tmp := false
			if nums[j] == left {
				tmp = check(i+1, 0, avg, flag^(1<<uint(j)))
			} else {
				tmp = check(i+1, k+1, left-nums[j], flag^(1<<uint(j)))
			}
			if tmp {
				return true
			}
		}
		return false
	}

	return check(0, 0, avg, 1<<uint(n)-1)
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{2, 2, 10, 5, 2, 7, 2, 2, 13}, 3))
}
