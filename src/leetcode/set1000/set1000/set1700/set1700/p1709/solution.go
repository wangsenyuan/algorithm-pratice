package p1709

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		a, b := boxTypes[i], boxTypes[j]
		return a[1] > b[1]
	})

	var res int

	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		cur := boxTypes[i]
		a, b := cur[0], cur[1]
		x := min(a, truckSize)
		res += b * x
		truckSize -= x
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
