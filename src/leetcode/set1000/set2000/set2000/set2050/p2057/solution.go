package p2057

func smallestEqual(nums []int) int {

	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}

	return -1
}
