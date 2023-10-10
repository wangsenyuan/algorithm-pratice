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
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 50

func solve(n int, E [][]int) int {
	// the network is connected
	// pick an edge i, connect ui, vi
	// 存在一种情况，把一条线，不断的连接到1，然后再不断的连接到n
	// dp[a][b]是节点a，b间的最短距离
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = inf
		}
		f[i][i] = 0
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		f[u][v] = 1
		f[v][u] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}

	var ans = inf

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		ans = min(ans, w*(1+f[0][u]+f[v][n-1]))
		ans = min(ans, w*(1+f[0][v]+f[u][n-1]))

		for i := 0; i < n; i++ {
			ans = min(ans, w*(f[0][i]+f[i][n-1]+f[v][i]+2))
			ans = min(ans, w*(f[0][i]+f[i][n-1]+f[u][i]+2))
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
