package p2860

func sumIndicesWithKSetBits(nums []int, k int) int {
	var res int

	for i, num := range nums {
		if digitCount(i) == k {
			res += num
		}
	}
	return res
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
