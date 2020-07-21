package p477

func totalHammingDistance(nums []int) int {
	n := len(nums)
	total := 0

	for i := 0; i < 32; i++ {
		x := 0
		for _, num := range nums {
			if num&(1<<uint(i)) != 0 {
				x++
			}
		}
		total += x * (n - x)
	}

	return total
}
