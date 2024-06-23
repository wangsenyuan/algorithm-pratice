package p3193

func minimumOperations(nums []int) int {
	var res int

	for _, num := range nums {
		num %= 3
		if num != 0 {
			res++
		}
	}
	return res
}
