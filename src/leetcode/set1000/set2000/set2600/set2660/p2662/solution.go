package p2662

const INF = 1 << 30

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	n := len(specialRoads)
	// dp[i][j] = 使用j+1条边叨叨speicalRoads末端时的cost
	dp := make([]int, n)

	distanceToRoad := func(a []int, b []int) int {
		x := distance(a, b[:2]) + b[4]
		y := distance(a, b[2:4])
		return min(x, y)
	}
	fp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = distanceToRoad(start, specialRoads[i])
		fp[i] = INF
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			// if connect i & k to get a new path
			for k := 0; k < n; k++ {
				fp[k] = min(fp[k], dp[i]+distanceToRoad(specialRoads[i][2:4], specialRoads[k]))
			}
		}
		for i := 0; i < n; i++ {
			dp[i] = min(dp[i], fp[i])
			fp[i] = INF
		}
	}

	res := distance(start, target)

	for i := 0; i < n; i++ {
		res = min(res, distance(target, specialRoads[i][2:4])+dp[i])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func distance(a, b []int) int {
	dx := abs(a[0] - b[0])
	dy := abs(a[1] - b[1])
	return dx + dy
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
