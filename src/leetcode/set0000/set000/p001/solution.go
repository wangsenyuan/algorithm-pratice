package p001

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		b := target - a
		if j, found := cache[b]; found {
			return []int{i, j}
		}
		cache[a] = i
	}
	return nil
}
