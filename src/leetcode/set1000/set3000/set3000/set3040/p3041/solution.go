package p3041

func maxOperations(nums []int) int {
	var res int

	expect := nums[0] + nums[1]

	for i := 0; i+2 <= len(nums); i += 2 {
		if nums[i]+nums[i+1] != expect {
			break
		}
		res++
	}
	return res
}
