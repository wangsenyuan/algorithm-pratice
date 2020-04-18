package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)

		grid := make([][]int, m)

		for i := 0; i < m; i++ {
			grid[i] = readNNums(reader, n)
		}

		fmt.Println(solve(m, n, grid))
	}
}

func solve(m, n int, grid [][]int) int64 {

	row := make([][]int64, m)
	for i := 0; i < m; i++ {
		row[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			row[i][j] = int64(grid[i][j])
			if j > 0 {
				row[i][j] += row[i][j-1]
			}
		}
	}

	col := make([][]int64, n)

	for j := 0; j < n; j++ {
		col[j] = make([]int64, m)
		for i := 0; i < m; i++ {
			col[j][i] = int64(grid[i][j])
			if i > 0 {
				col[j][i] += col[j][i-1]
			}
		}
	}

	bob := make([][][][]int64, m)

	for i := 0; i < m; i++ {
		bob[i] = make([][][]int64, n)
		for j := 0; j < n; j++ {
			bob[i][j] = make([][]int64, m)
			for k := 0; k < m; k++ {
				bob[i][j][k] = make([]int64, n)
			}
		}
	}

	var total int64

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += int64(grid[i][j])
		}
	}

	rowSum := func(r, c0, c1 int) int64 {
		if c0 > c1 {
			return 0
		}
		res := row[r][c1]
		if c0 > 0 {
			res -= row[r][c0-1]
		}
		return res
	}

	colSum := func(c, r0, r1 int) int64 {
		if r0 > r1 {
			return 0
		}
		res := col[c][r1]
		if r0 > 0 {
			res -= col[c][r0-1]
		}
		return res
	}

	aliceMove := func(a, b, c, d int) int {
		row1 := rowSum(a, b, d)
		row2 := rowSum(c, b, d)
		col1 := colSum(b, a, c)
		col2 := colSum(d, a, c)
		if row1 <= row2 && row1 <= col1 && row1 <= col2 {
			return 0
		}
		if row2 <= row1 && row2 <= col1 && row2 <= col2 {
			return 1
		}

		if col1 <= row1 && col1 <= row2 && col1 <= col2 {
			return 2
		}

		return 3
	}

	getBob := func(a, b, c, d int) int64 {
		if a > c || b > d {
			return 0
		}
		return bob[a][b][c][d]
	}

	for w := 1; w <= n; w++ {
		for h := 1; h <= m; h++ {
			for x := 0; x < m; x++ {
				for y := 0; y < n; y++ {
					u, v := x+h-1, y+w-1
					if u >= m || v >= n {
						continue
					}
					a, b, c, d := x, y, u, v
					alice := aliceMove(a, b, c, d)
					if alice == 0 {
						a++
					} else if alice == 1 {
						c--
					} else if alice == 2 {
						b++
					} else {
						d--
					}
					var res int64
					if a <= u {
						res = max(res, rowSum(a, b, d)+getBob(a+1, b, c, d))
					}
					if c >= x {
						res = max(res, rowSum(c, b, d)+getBob(a, b, c-1, d))
					}
					if b <= v {
						res = max(res, colSum(b, a, c)+getBob(a, b+1, c, d))
					}
					if d >= y {
						res = max(res, colSum(d, a, c)+getBob(a, b, c, d-1))
					}
					bob[x][y][u][v] = res
				}
			}
		}
	}

	ans1 := bob[0][0][m-1][n-1]
	ans2 := total - ans1

	if ans1 > ans2 {
		return ans1
	} else if ans2 > ans1 {
		return ans2
	}
	return total
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
