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

	nums5 := nums2[:10]
	fmt.Printf("steps 5: nums5 = %v, len = %d, cap = %d\n", nums5, len(nums5), cap(nums5))

	nums6 := nums2[1:7]
	fmt.Printf("step 6: nums6 = %v, len = %d, cap = %d\n", nums6, len(nums6), cap(nums6))
	fmt.Printf("\ttry to get a number after the len: %d\n", nums6[:8][len(nums6)])
	// var x int
	// fmt.Printf("\ttry to get a number before the 0: %d\n", nums6[x-1 : 2][0])
	nums7 := nums2[1:7:8]
	fmt.Printf("step 7: nums7 = %v, len = %d, cap = %d\n", nums7, len(nums7), cap(nums7))
	// fmt.Printf("\ttry to get a number after the len: %v\n", nums7[:8])
	nums7[5] = 99
	fmt.Printf("\t: change nums7[5] to 99, nums7=%v, nums=%v\n", nums7, nums)

	nums[7] = 100
	nums7 = nums7[:7]
	fmt.Printf("\t: change nums[7] to 100, nums7=%v, nums=%v\n", nums7, nums)
	nums8 := nums7[:8]
	fmt.Printf("step 8: try to slice more than cap: nums8 = %v\n", nums8)
}
