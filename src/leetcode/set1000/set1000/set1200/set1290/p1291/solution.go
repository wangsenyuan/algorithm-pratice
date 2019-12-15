package p1291

import (
	"sort"
)

var nums []int

func init() {

	var gen func(prev int, i int, num int)

	gen = func(prev int, i int, num int) {
		if num < 0 {
			return
		}
		if i == 10 {
			nums = append(nums, num)
			return
		}

		if prev == 0 {
			// gen(0, i+1, num)
			for x := 0; x <= 9; x++ {
				gen(x, i+1, x)
			}
		} else if prev < 9 {
			prev++
			gen(prev, i+1, num*10+prev)
		}
	}

	gen(0, 0, 0)

	sort.Ints(nums)
	// fmt.Fprintf(os.Stderr, "%v", nums)
}

func sequentialDigits(low int, high int) []int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= low
	})

	j := sort.Search(len(nums), func(j int) bool {
		return nums[j] > high
	})

	return nums[i:j]
}
