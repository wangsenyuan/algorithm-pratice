package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}

		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, row := range res {
			for i := 0; i < len(row); i++ {
				buf.WriteString(fmt.Sprintf("%d ", row[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

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

func solve(grid []string) [][]int {
	n := len(grid)
	m := len(grid[0])
	maze := make([][]byte, n)

	for i := 0; i < n; i++ {
		maze[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				maze[i][j] = 0
			} else {
				maze[i][j] = 1
			}
		}
	}

	var res [][]int

	process := func(i int, j int) {
		// 必须保证左上角是0
		if maze[i][j] == 0 {
			return
		}
		var cur []int
		if maze[i][j] == 1 {
			cur = append(cur, i, j)
			maze[i][j] = 0
		}
		if maze[i][j+1] == 1 {
			cur = append(cur, i, j+1)
			maze[i][j+1] = 0
		}
		if maze[i+1][j] == 1 {
			cur = append(cur, i+1, j)
			maze[i+1][j] = 0
		}

		if len(cur) < 6 {
			cur = append(cur, i+1, j+1)
			maze[i+1][j+1] ^= 1
		}
		if len(cur) < 6 {
			cur = append(cur, i+1, j)
			maze[i+1][j] ^= 1
		}
		if len(cur) < 6 {
			cur = append(cur, i, j+1)
			maze[i][j+1] ^= 1
		}
		res = append(res, cur)
	}

	for i := 0; i+2 < n; i++ {
		for j := 0; j+2 < m; j++ {
			process(i, j)
		}
		// 保证上面一行是0
		if maze[i][m-2] == 0 && maze[i][m-1] == 0 {
			continue
		}
		var arr []int
		if maze[i][m-2] == 1 {
			arr = append(arr, i, m-2)
			maze[i][m-2] = 0
		}
		if maze[i][m-1] == 1 {
			arr = append(arr, i, m-1)
			maze[i][m-1] = 0
		}
		if len(arr) < 6 {
			arr = append(arr, i+1, m-2)
			maze[i+1][m-2] ^= 1
		}
		if len(arr) < 6 {
			arr = append(arr, i+1, m-1)
			maze[i+1][m-1] ^= 1
		}
		res = append(res, arr)
	}

	// 保证最后两行都是0
	for j := 0; j+2 < m; j++ {
		if maze[n-2][j] == 0 && maze[n-1][j] == 0 {
			continue
		}
		var arr []int
		if maze[n-2][j] == 1 {
			arr = append(arr, n-2, j)
			maze[n-2][j] ^= 1
		}
		if maze[n-1][j] == 1 {
			arr = append(arr, n-1, j)
			maze[n-1][j] ^= 1
		}
		if len(arr) < 6 {
			arr = append(arr, n-2, j+1)
			maze[n-2][j+1] ^= 1
		}
		if len(arr) < 6 {
			arr = append(arr, n-1, j+1)
			maze[n-1][j+1] ^= 1
		}
		res = append(res, arr)
	}
	// 最后2*2的区域
	sum := maze[n-2][m-2] + maze[n-2][m-1] + maze[n-1][m-2] + maze[n-1][m-1]

	if sum == 4 {
		res = append(res, []int{n - 2, m - 2, n - 2, m - 1, n - 1, m - 2})
		maze[n-2][m-2] ^= 1
		maze[n-2][m-1] ^= 1
		maze[n-1][m-2] ^= 1
		sum = 1
	}
	if sum == 1 {
		// flice the one, and another two 0 to get 2
		var one [][]int
		var zero [][]int
		for i := n - 2; i < n; i++ {
			for j := m - 2; j < m; j++ {
				if maze[i][j] == 0 {
					zero = append(zero, []int{i, j})
				} else {
					one = append(one, []int{i, j})
				}
			}
		}
		var arr []int
		arr = append(arr, zero[0]...)
		maze[zero[0][0]][zero[0][1]] ^= 1
		arr = append(arr, zero[1]...)
		maze[zero[1][0]][zero[1][1]] ^= 1
		arr = append(arr, one[0]...)
		maze[one[0][0]][one[0][1]] ^= 1
		res = append(res, arr)
		sum = 2
	}
	if sum == 2 {
		var one [][]int
		var zero [][]int
		for i := n - 2; i < n; i++ {
			for j := m - 2; j < m; j++ {
				if maze[i][j] == 0 {
					zero = append(zero, []int{i, j})
				} else {
					one = append(one, []int{i, j})
				}
			}
		}
		// 一个1个2个0
		var arr []int
		arr = append(arr, zero[0]...)
		maze[zero[0][0]][zero[0][1]] ^= 1
		arr = append(arr, zero[1]...)
		maze[zero[1][0]][zero[1][1]] ^= 1
		arr = append(arr, one[0]...)
		maze[one[0][0]][one[0][1]] ^= 1
		res = append(res, arr)
		sum = 3
	}
	if sum == 3 {
		// just found those 3 ones, and reverse them
		var arr []int
		for i := n - 2; i < n; i++ {
			for j := m - 2; j < m; j++ {
				if maze[i][j] == 1 {
					arr = append(arr, i, j)
					maze[i][j] = 0
				}
			}
		}
		res = append(res, arr)
	}

	for _, cur := range res {
		for i := 0; i < len(cur); i++ {
			cur[i]++
		}
	}

	return res
}
