package p950

func deckRevealedIncreasing(deck []int) []int {
	var loop func(nums []int) []int

	loop = func(nums []int) []int {
		n := len(nums)
		if n == 0 {
			return nil
		}
		if n == 1 {
			return nums
		}
		var j int
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[j] {
				j = i
			}
		}
		first := nums[j]
		copy(nums[j:], nums[j+1:])
		tmp := loop(nums[:n-1])
		last := tmp[n-2]
		tmp = tmp[:n-2]
		return append([]int{first, last}, tmp...)
	}
	return loop(deck)
}
