package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		res := solve(grid)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(grid []string) bool {
	n := len(grid)
	m := len(grid[0])

	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}

	var dfs func(r int, c int, cnt int) (bool, [][]int)

	dfs = func(r int, c int, cnt int) (bool, [][]int) {
		if cnt == 3 {
			return false, nil
		}
		vis[r][c] = true
		var path [][]int
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 || r+i < 0 || r+i == n || c+j < 0 || c+j == m || vis[r+i][c+j] || grid[r+i][c+j] == '.' {
					continue
				}
				ok, tmp := dfs(r+i, c+j, cnt+1)
				if !ok {
					return false, nil
				}
				path = append(path, tmp...)
			}
		}
		path = append(path, []int{r, c})
		return true, path
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				if !vis[i][j] {
					ok, path := dfs(i, j, 0)
					if !ok || len(path) != 3 {
						return false
					}
					xs := get(path, 0)
					ys := get(path, 1)
					if len(xs) != 2 || xs[1] != xs[0]+1 {
						return false
					}
					if len(ys) != 2 || ys[1] != ys[0]+1 {
						return false
					}
				}
			}
		}
	}

	return true
}

func get(arr [][]int, id int) []int {
	res := make([]int, 0, 2)
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[i][id])
	}
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
}
