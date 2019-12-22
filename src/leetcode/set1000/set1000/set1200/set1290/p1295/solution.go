package p1295

func findNumbers(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		cnt := countDigits(nums[i])
		if cnt%2 == 0 {
			res++
		}
	}
	return res
}

func countDigits(num int) int {
	var res int

	for num > 0 {
		res++
		num /= 10
	}
	return res
}
