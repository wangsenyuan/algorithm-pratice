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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		R, C := readTwoNums(reader)
		G := make([]string, R)
		for i := 0; i < R; i++ {
			G[i], _ = reader.ReadString('\n')
			G[i] = G[i][:C]
		}

		fmt.Printf("%.9f\n", solve(R, C, G))
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(R, C int, G []string) float64 {
	color := make([]int, R*C)

	var dfs func(i, j, c int)

	dfs = func(i, j, c int) {
		if color[i*C+j] == c {
			return
		}
		color[i*C+j] = c

		for k := 0; k < 4; k++ {
			u, v := i+dd[k], j+dd[k+1]
			if u >= 0 && u < R && v >= 0 && v < C && G[u][v] == 'o' {
				dfs(u, v, c)
			}
		}
	}

	var cur int

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if G[i][j] == 'o' && color[i*C+j] == 0 {
				cur++
				dfs(i, j, cur)
			}
		}
	}

	X := make([]int, cur)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if color[i*C+j] > 0 {
				X[color[i*C+j]-1]++
			}
		}
	}

	var res float64 = 1

	for i := 1; i < cur; i++ {
		res += float64(X[i]) / float64(X[0]+X[i])
	}

	return res
}
