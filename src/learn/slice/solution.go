package main

import "fmt"

func main() {
	nums := make([]int, 10)
	for i := 1; i <= 10; i++ {
		nums[i-1] = i
	}

	fmt.Printf("step 1: nums = %v, len = %d, cap = %d\n", nums, len(nums), cap(nums))

	nums2 := nums[:5]

	fmt.Printf("steps 2: nums2 = %v, len = %d, cap = %d\n", nums2, len(nums2), cap(nums2))

	nums3 := nums[3:8]

	fmt.Printf("steps 3: nums3 = %v, len = %d, cap = %d\n", nums3, len(nums3), cap(nums3))
	nums3[0] = 20

	for i := 0; i < 5; i++ {
		nums3 = append(nums3, i+13)
	}

	fmt.Printf("steps 4, extends nums3\n")
	fmt.Printf("			nums = %v, len = %d, cap = %d\n", nums, len(nums), cap(nums))
	fmt.Printf("			nums2 = %v, len = %d, cap = %d\n", nums2, len(nums2), cap(nums2))
	fmt.Printf("			nums3 = %v, len = %d, cap = %d\n", nums3, len(nums3), cap(nums3))
}
