package p2193

func mostFrequent(nums []int, key int) int {
	freq := make(map[int]int)

	for i := 0; i+1 < len(nums); i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}
	var ans int
	var cnt int
	for k, v := range freq {
		if v > cnt {
			ans = k
			cnt = v
		}
	}
	return ans
}
