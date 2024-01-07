package p2999

func missingInteger(nums []int) int {
	n := len(nums)

	mem := make(map[int]int)

	for _, num := range nums {
		mem[num]++
	}

	sum := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1]+1 {
			break
		}
		sum += nums[i]
	}
	for mem[sum] > 0 {
		sum++
	}
	return sum
}
