package p1620

import "sort"

func trimMean(arr []int) float64 {
	n := len(arr)
	cnt := n / 20

	sort.Ints(arr)

	tmp := arr[cnt : n-cnt]

	var sum float64

	for i := 0; i < len(tmp); i++ {
		sum += float64(tmp[i])
	}

	return sum / float64(len(tmp))
}
