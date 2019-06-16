package main

func topKFrequent1(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums)+1)

	for num, f := range freq {
		if len(bucket[f]) == 0 {
			bucket[f] = make([]int, 0, 1)
		}
		bucket[f] = append(bucket[f], num)
	}

	ans := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		for j := 0; j < len(bucket[i]) && k > 0; j++ {
			ans = append(ans, bucket[i][j])
			k--
		}
	}
	return ans
}
