package p169

func majorityElement(nums []int) int {
	x := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == x {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			x = nums[i]
			cnt = 1
		}
	}
	return x
}
