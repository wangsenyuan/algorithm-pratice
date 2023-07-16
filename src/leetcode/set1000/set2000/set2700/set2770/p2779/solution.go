package p2779

func minimumIndex(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	ans := nums[0]
	for k, v := range freq {
		if v > freq[ans] {
			ans = k
		}
	}
	prev := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		prev[nums[i]]++
		freq[nums[i]]--
		if 2*prev[ans] > i+1 && 2*freq[ans] > (n-i-1) {
			return i
		}
	}
	return -1
}
