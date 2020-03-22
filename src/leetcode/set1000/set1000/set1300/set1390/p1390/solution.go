package p1390

func sumFourDivisors(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res += sumDiv(nums[i])
	}

	return res
}

func sumDiv(num int) int {
	var sum int
	var cnt int
	for x := 1; x*x <= num && cnt <= 4; x++ {
		if num%x == 0 {
			sum += x
			cnt++
			y := num / x
			if x != y {
				sum += y
				cnt++
			}
		}
	}
	if cnt != 4 {
		return 0
	}
	return sum
}
