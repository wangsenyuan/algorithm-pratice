package p3265

import "sort"

func countPairs(nums []int) int {
	sort.Ints(nums)

	freq := make(map[int]int)

	base := make([]int, 8)

	base[0] = 1
	for i := 1; i < 8; i++ {
		base[i] = 10 * base[i-1]
	}

	process := func(num int) int {
		var arr []int
		for tmp := num; tmp > 0; tmp /= 10 {
			arr = append(arr, tmp%10)
		}
		mem := make(map[int]bool)
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				tmp := num
				tmp = tmp - arr[i]*base[i] + arr[j]*base[i]
				tmp = tmp - arr[j]*base[j] + arr[i]*base[j]
				mem[tmp] = true
			}
		}
		mem[num] = true
		var res int
		for k := range mem {
			res += freq[k]
		}
		return res
	}

	var res int
	for _, num := range nums {
		res += process(num)
		freq[num]++
	}

	return res
}
