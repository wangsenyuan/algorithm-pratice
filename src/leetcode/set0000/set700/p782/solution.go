package p782

func movesToChessboard(board [][]int) int {
	count := make(map[int]int)

	for _, row := range board {
		var tmp int
		for _, cell := range row {
			tmp = tmp*2 + cell
		}
		count[tmp]++
	}

	k1 := solve(count, len(board))

	if k1 < 0 {
		return -1
	}

	count = make(map[int]int)

	for j := 0; j < len(board); j++ {
		var tmp int
		for i := 0; i < len(board); i++ {
			tmp = tmp*2 + board[i][j]
		}
		count[tmp]++
	}

	k2 := solve(count, len(board))
	if k2 < 0 {
		return -1
	}

	return k1 + k2
}

func solve(count map[int]int, n int) int {
	if len(count) != 2 {
		return -1
	}

	var k1, k2 int
	for k := range count {
		if k1 == 0 {
			k1 = k
		} else {
			k2 = k
		}
	}

	if n%2 == 0 {
		if count[k1] != count[k2] {
			return -1
		}
	} else {
		if count[k1] != count[k2]+1 && count[k2] != count[k1]+1 {
			return -1
		}
	}

	N := 1<<uint(n) - 1
	if k1^k2 != N {
		return -1
	}

	ones := bitCount(k1)

	res := N + 1

	if n%2 == 0 || ones*2 < n {
		res = min(res, bitCount(k1^0xAAAAAAAA&N)/2)
	}

	if n%2 == 0 || ones*2 > n {
		res = min(res, bitCount(k1^0x55555555&N)/2)
	}

	return res
}

func bitCount(i int) int {
	// HD, Figure 5-2
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	return i & 0x3f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
