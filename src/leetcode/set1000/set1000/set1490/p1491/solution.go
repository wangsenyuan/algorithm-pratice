package p1491

import "math"

func average(salary []int) float64 {
	var sum int
	var min = math.MaxInt32
	var max = math.MinInt32

	for _, num := range salary {
		sum += num

		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}

	sum -= min
	sum -= max

	return float64(sum) / float64(len(salary)-2)
}
