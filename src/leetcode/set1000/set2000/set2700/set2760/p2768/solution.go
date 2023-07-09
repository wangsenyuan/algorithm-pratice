package p2768

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	cnt := make(map[int64]int)

	for _, coor := range coordinates {
		x, y := coor[0], coor[1]
		if x-1 >= 0 && y-1 >= 0 {
			k := int64(x-1)*int64(n) + int64(y) - 1
			cnt[k]++
		}
		if x-1 >= 0 && y+1 < n {
			k := int64(x-1)*int64(n) + int64(y)
			cnt[k]++
		}
		if x+1 < m && y-1 >= 0 {
			k := int64(x)*int64(n) + int64(y-1)
			cnt[k]++
		}
		if x+1 < m && y+1 < n {
			k := int64(x)*int64(n) + int64(y)
			cnt[k]++
		}
	}

	res := make([]int64, 5)
	var sum int64
	for _, v := range cnt {
		sum++
		res[v]++
	}
	res[0] = int64(m-1)*int64(n-1) - sum
	return res
}
