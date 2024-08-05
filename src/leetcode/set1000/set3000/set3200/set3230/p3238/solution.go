package p3238

func winningPlayerCount(n int, pick [][]int) int {
	win := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		win[i] = make(map[int]int)
	}
	for _, cur := range pick {
		x, y := cur[0], cur[1]
		win[x][y]++
	}
	var res int
	for i := 0; i < n; i++ {
		for _, v := range win[i] {
			if v > i {
				res++
				break
			}
		}
	}
	return res
}
