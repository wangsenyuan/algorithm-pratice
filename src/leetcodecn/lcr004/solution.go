package lcr004

func singleNumber(nums []int) int {
	cnt := make([]int, 3)
	for _, num := range nums {
		if num == 0 {
			cnt[0]++
		} else if num > 0 {
			cnt[1]++
		} else {
			cnt[2]++
		}
	}
	if cnt[0] == 1 {
		return 0
	}
	sign := 1
	var arr []int
	if cnt[1]%3 == 1 {
		for _, num := range nums {
			if num > 0 {
				arr = append(arr, num)
			}
		}
	} else {
		sign = -1
		for _, num := range nums {
			if num < 0 {
				arr = append(arr, -num)
			}
		}
	}

	return sign * solve(arr)
}

func solve(arr []int) int {
	freq := make([]int, 32)
	for _, num := range arr {
		for i := 0; i < 32; i++ {
			freq[i] += (num >> i) & 1
		}
	}
	var res int
	for i := 0; i < 32; i++ {
		if freq[i]%3 == 1 {
			res |= 1 << i
		}
	}
	return res
}
