package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		mat := make([]string, n)
		for i := 0; i < n; i++ {
			mat[i] = readString(reader)[:m]
		}
		res := solve(mat, k)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(grid []string, k int) bool {
	n := len(grid)
	m := len(grid[0])
	vis := make([]int, n)

	check := func(x int, y int) bool {
		var d int
		for x-(d+1) >= 0 && y-(d+1) >= 0 && y+(d+1) < m && grid[x-(d+1)][y-(d+1)] == '*' && grid[x-(d+1)][y+(d+1)] == '*' {
			d++
		}

		return d >= k
	}

	mark := func(x int, y int) {
		vis[x] |= 1 << y
		var d int
		for x-(d+1) >= 0 && y-(d+1) >= 0 && y+(d+1) < m && grid[x-(d+1)][y-(d+1)] == '*' && grid[x-(d+1)][y+(d+1)] == '*' {
			d++
			vis[x-d] |= 1 << (y - d)
			vis[x-d] |= 1 << (y + d)
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				// not checked yet
				if check(i, j) {
					mark(i, j)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' && (vis[i]>>j)&1 == 0 {
				return false
			}
		}
	}

	return true
}
