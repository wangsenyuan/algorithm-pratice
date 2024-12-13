package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}

	res := solve(grid)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i, x := range res {
		buf.WriteString(x)
		if i+1 < len(res) {
			buf.WriteByte(' ')
		} else {
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

type pair struct {
	first  int
	second int
}

func solve(grid []string) []string {

	n := len(grid)
	m := len(grid[0])

	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != '.' {
				cnt++
			}
		}
	}

	vis := make(map[pair]int)

	push := func(r int, c int) {
		if grid[r][c] != '.' {
			vis[pair{r, c}]++
		}
	}

	pop := func(r int, c int) {
		if grid[r][c] != '.' {
			key := pair{r, c}
			vis[key]--
			if vis[key] == 0 {
				delete(vis, key)
			}
		}
	}

	var dfs func(r int, c int, d int) bool

	dfs = func(r int, c int, d int) bool {
		if r < 0 || r == n || c < 0 || c == m {
			return len(vis) == cnt
		}
		push(r, c)
		or, oc := r, c
		// d = 0, from north,
		// d = 1, from west
		// d = 2, from south
		// d = 3, from east
		if grid[r][c] == '.' {
			if d == 0 {
				r++
			} else if d == 1 {
				c++
			} else if d == 2 {
				r--
			} else {
				c--
			}
		} else if grid[r][c] == '\\' {
			if d == 0 {
				c++
				d = 1
			} else if d == 1 {
				r++
				d = 0
			} else if d == 2 {
				c--
				d = 3
			} else {
				r--
				d = 2
			}
		} else {
			// grid[r][c] == '/'
			if d == 0 {
				c--
				d = 3
			} else if d == 1 {
				r--
				d = 2
			} else if d == 2 {
				c++
				d = 1
			} else {
				r++
				d = 0
			}
		}

		if dfs(r, c, d) {
			return true
		}

		pop(or, oc)
		return false
	}

	reset := func() {
		clear(vis)
	}

	var ans []string
	for i := 0; i < n; i++ {
		reset()
		// from west
		if dfs(i, 0, 1) {
			ans = append(ans, fmt.Sprintf("W%d", i+1))
		}
		reset()
		if dfs(i, m-1, 3) {
			ans = append(ans, fmt.Sprintf("E%d", i+1))
		}
	}

	for i := 0; i < m; i++ {
		reset()
		if dfs(0, i, 0) {
			ans = append(ans, fmt.Sprintf("N%d", i+1))
		}
		reset()
		if dfs(n-1, i, 2) {
			ans = append(ans, fmt.Sprintf("S%d", i+1))
		}
	}

	return ans
}
