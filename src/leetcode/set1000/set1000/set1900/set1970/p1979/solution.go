package p1979

func findGCD(nums []int) int {
	min := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
