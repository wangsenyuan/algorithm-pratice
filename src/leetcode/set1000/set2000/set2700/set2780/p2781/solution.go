package p2781

func sumOfSquares(nums []int) int {
	var res int
	n := len(nums)
	for i, num := range nums {
		if n%(i+1) == 0 {
			res += num * num
		}
	}
	return res
}
