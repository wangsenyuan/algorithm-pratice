package p1745

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	preSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + int64(candiesCount[i])
	}

	ans := make([]bool, len(queries))

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		fav, day, can := cur[0], cur[1], cur[2]
		tot := preSum[fav]
		//
		earliest := tot / int64(can)

		latest := tot + int64(candiesCount[fav]) - 1

		ans[i] = earliest <= int64(day) && int64(day) <= latest
	}

	return ans
}
