package p3267

import (
	"slices"
	"strconv"
)

func countPairs(nums []int) int {

	slices.SortFunc(nums, func(a, b int) int {
		x := strconv.Itoa(a)
		y := strconv.Itoa(b)
		return len(x) - len(y)
	})

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
		mem[num] = true

		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				tmp := num
				tmp = tmp - arr[i]*base[i] + arr[j]*base[i]
				tmp = tmp - arr[j]*base[j] + arr[i]*base[j]
				mem[tmp] = true
				arr[i], arr[j] = arr[j], arr[i]
				for u := i + 1; u < len(arr); u++ {
					for v := u + 1; v < len(arr); v++ {
						tmp2 := tmp
						tmp2 = tmp2 - arr[u]*base[u] + arr[v]*base[u]
						tmp2 = tmp2 - arr[v]*base[v] + arr[u]*base[v]
						mem[tmp2] = true
					}
				}
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
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
