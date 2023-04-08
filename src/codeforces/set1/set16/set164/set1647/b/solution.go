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
		n, m := readTwoNums(reader)
		A := make([]string, n)
		for i := 0; i < n; i++ {
			A[i] = readString(reader)[:m]
		}
		res := solve(A)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(A []string) bool {
	// 只有一整行（列），且没有相邻
	n := len(A)
	m := len(A[0])

	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		for j := 0; j < m; j++ {
			vis[i][j] = -1
		}
	}

	var dfs func(i int, j, c int)

	dfs = func(i int, j, c int) {
		vis[i][j] = c
		for k := 0; k < 4; k++ {
			u, v := i+dd[k], j+dd[k+1]
			if u >= 0 && u < n && v >= 0 && v < m && A[u][v] == '1' && vis[u][v] < 0 {
				dfs(u, v, c)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == '1' && vis[i][j] < 0 {
				dfs(i, j, i*m+j)
			}
		}
	}

	bound := make(map[int][]int)

	update := func(g int, i int, j int) {
		if _, ok := bound[g]; !ok {
			bound[g] = []int{i, j, i, j}
			return
		}
		bound[g][0] = min(bound[g][0], i)
		bound[g][1] = min(bound[g][1], j)
		bound[g][2] = max(bound[g][2], i)
		bound[g][3] = max(bound[g][3], j)
	}

	sz := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == '1' {
				g := vis[i][j]
				update(g, i, j)
				sz[g]++
			}
		}
	}

	for g, b := range bound {
		x := sz[g]
		a := (b[2] - b[0] + 1) * (b[3] - b[1] + 1)
		if a != x {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
