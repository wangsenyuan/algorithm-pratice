package p3296

import (
	"sort"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	check := func(expect int) bool {
		// 花费了这么多秒，减低了多少高度
		// w[i] * (1 + x) * x / 2 <= expect
		var sum int
		for _, w := range workerTimes {
			x := sort.Search(mountainHeight, func(i int) bool {
				return (1+i)*i/2 > expect/w
			})
			if (1+x)*x/2*w > expect {
				x--
			}
			sum += x
		}

		return sum >= mountainHeight
	}

	res := sort.Search(5000050000000001, check)
	return int64(res)
}
