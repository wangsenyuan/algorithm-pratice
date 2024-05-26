package p3163

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	freq := make([]int, 1e6+1)

	for _, num := range nums2 {
		freq[num]++
	}

	var res int

	for _, num := range nums1 {
		if num%k != 0 {
			continue
		}
		num /= k
		for i := 1; i <= num/i; i++ {
			if num%i == 0 {
				res += freq[i]
				if num != i*i {
					res += freq[num/i]
				}
			}
		}
	}

	return int64(res)
}
