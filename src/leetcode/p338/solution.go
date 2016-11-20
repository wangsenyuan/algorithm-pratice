package main

import "fmt"

func main() {
	fmt.Println(countBits(8))
}

func countBits(num int) []int {
	ans := make([]int, num+1)

	for i := 0; i < num; i += 2 {
		ans[i] = ans[i>>1]
		ans[i+1] = ans[i] + 1
	}
	ans[num] = ans[num>>1] + (num & 1)

	return ans
}

func countBits1(num int) []int {
	xs := make([]int, 32)
	ans := []int{0}
	lastOneAt := -1
	lastCnt := 0
	for i := 1; i <= num; i++ {
		if lastOneAt != 0 {
			lastCnt++
			lastOneAt = 0
			xs[lastOneAt] = 1
		} else {
			for xs[lastOneAt] == 1 {
				xs[lastOneAt] = 0
				lastCnt--
				lastOneAt++
			}
			lastCnt++
			xs[lastOneAt] = 1
		}
		ans = append(ans, lastCnt)
	}

	return ans
}
