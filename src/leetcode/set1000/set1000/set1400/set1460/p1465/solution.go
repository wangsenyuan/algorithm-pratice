package p1465

import "sort"

const MOD = 1000000007

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	var x = horizontalCuts[0]

	for i := 1; i < len(horizontalCuts); i++ {
		x = max(x, horizontalCuts[i]-horizontalCuts[i-1])
	}

	x = max(x, h-horizontalCuts[len(horizontalCuts)-1])

	var y = verticalCuts[0]

	for i := 1; i < len(verticalCuts); i++ {
		y = max(y, verticalCuts[i]-verticalCuts[i-1])
	}

	y = max(y, w-verticalCuts[len(verticalCuts)-1])

	X := int64(x)
	Y := int64(y)

	return int(X * Y % MOD)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
