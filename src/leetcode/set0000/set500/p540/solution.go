package p540

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
