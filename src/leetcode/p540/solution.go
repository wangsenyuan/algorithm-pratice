package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3,3,7,7,10,11,11}))
	fmt.Println(singleNonDuplicate([]int{3,3,7,7,10,10,11}))
	fmt.Println(singleNonDuplicate([]int{2, 3,3,7,7,10,10}))



}

func singleNonDuplicate(nums []int) int {
	n := len(nums)

	i, j := 0, n

	for j-i > 1 {
		k := i + (j-i)/2
		a := k - i

		if a%2 == 1 {
			//a is odd
			//two situation here
			if nums[k-1] == nums[k] {
				//single num in the second part
				i = k + 1
			} else {
				//single num in the first part
				j = k
			}
		} else {
			if nums[k-1] == nums[k] {
				//single num in the first part
				j = k - 1
			} else {
				//single num in the second part
				i = k
			}
		}
	}

	return nums[i]
}
