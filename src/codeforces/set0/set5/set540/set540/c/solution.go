package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}
	src := readNNums(reader, 2)
	dst := readNNums(reader, 2)
	res := solve(n, m, grid, src, dst)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, m int, grid []string, src []int, dst []int) bool {

	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
	}

	for i := 0; i < 2; i++ {
		src[i]--
		dst[i]--
	}

	// 那些已经cracked的，vis[i][j] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'X' && (src[0] != i || src[1] != j) {
				vis[i][j] = 1
			}
		}
	}

	var dfs func(u int, v int) bool

	dfs = func(u int, v int) bool {
		if u == dst[0] && v == dst[1] {
			if vis[u][v] == 1 {
				return true
			}
		}

		if vis[u][v] == 1 {
			return false
		}

		vis[u][v]++

		for i := 0; i < 4; i++ {
			r, c := u+dd[i], v+dd[i+1]
			if r < 0 || r == n || c < 0 || c == m {
				continue
			}
			if dfs(r, c) {
				return true
			}
		}

		return false
	}

	return dfs(src[0], src[1])
}
