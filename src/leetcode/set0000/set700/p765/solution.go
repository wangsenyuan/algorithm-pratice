package p765

func minSwapsCouples(row []int) int {
	//row := make([]int, len(ROW))
	//copy(row, ROW)
	n := len(row) / 2
	idx := make([]int, len(row))
	for i := 0; i < n; i++ {
		idx[row[2*i]] = 2 * i
		idx[row[2*i+1]] = 2*i + 1
	}

	var res int

	for i := 0; i < len(row); i += 2 {
		j := row[i] ^ 1
		if row[i+1] != j {
			swap(row, idx, i+1, idx[j])
			res++
		}
	}

	return res
}

func swap(row []int, pos []int, x int, y int) {
	temp := row[x]
	pos[temp] = y
	pos[row[y]] = x
	row[x] = row[y]
	row[y] = temp
}
