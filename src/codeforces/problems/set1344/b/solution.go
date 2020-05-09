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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)

	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i], _ = reader.ReadString('\n')
		grid[i] = grid[i][:m]
	}

	fmt.Println(solve(n, m, grid))
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m int, grid []string) int {
	col := make([]bool, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				// black
				if i > 0 && grid[i-1][j] == '.' {
					// white
					if col[j] {
						// has another black cell on top of white
						return -1
					}
				}
				col[j] = true
			}
		}
	}

	row := make([]bool, n)

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if grid[i][j] == '#' {
				if j > 0 && grid[i][j-1] == '.' {
					if row[i] {
						return -1
					}
				}
				row[i] = true
			}
		}
	}

	var c1, c2 int

	for j := 0; j < m; j++ {
		if col[j] {
			c1++
		}
	}

	for i := 0; i < n; i++ {
		if row[i] {
			c2++
		}
	}

	if c1 > 0 && c1 < m && c2 == n {
		return -1
	}

	if c2 > 0 && c2 < n && c1 == m {
		return -1
	}

	// safe

	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}

	var dfs func(x, y int)

	dfs = func(x, y int) {
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u < 0 || u >= n || v < 0 || v >= m {
				continue
			}
			if grid[u][v] == '#' && !vis[u][v] {
				vis[u][v] = true
				dfs(u, v)
			}
		}
	}

	var cnt int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' && !vis[i][j] {
				cnt++
				vis[i][j] = true
				dfs(i, j)
			}
		}
	}

	return cnt
}
