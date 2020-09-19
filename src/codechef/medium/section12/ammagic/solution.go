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
		n, m := readTwoNums(reader)
		B := make([][]byte, n)
		for i := 0; i < n; i++ {
			B[i], _ = reader.ReadBytes('\n')
		}
		buf.WriteString(fmt.Sprintf("%d\n", solve(n, m, B, tc)))
		tc--
	}
	fmt.Print(buf.String())
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

var dd = []int{-1, 0, 1, 0, -1}

const MAX = 1000010

var vis [MAX]int

func solve(n int, m int, B [][]byte, tc int) int {

	var dfs func(x, y int, num int, level int)

	dfs = func(x, y int, num int, level int) {
		vis[num] = tc
		if level >= 6 {
			return
		}

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < n && v >= 0 && v < m {
				next := num*10 + int(B[u][v]-'0')
				dfs(u, v, next, level+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] != '0' {
				dfs(i, j, int(B[i][j]-'0'), 1)
			}
		}
	}

	for i := 1; i < MAX; i++ {
		if vis[i] != tc {
			return i
		}
	}
	return 0
}
