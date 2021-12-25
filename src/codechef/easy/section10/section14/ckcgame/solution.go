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
		line := readNNums(reader, 5)
		n := line[0]
		A := line[1:3]
		B := line[3:]

		V := make([][]int, n)
		for i := 0; i < n; i++ {
			V[i] = readNNums(reader, n)
		}

		res := solve(n, V, A, B)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, V [][]int, A, B []int) int64 {

	dp := make([][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int64, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	for i := 0; i < 2; i++ {
		A[i]--
		B[i]--
	}

	var dfs func(x, y, steps int) int64

	dfs = func(x, y, steps int) int64 {
		if x < 0 || x >= n || y < 0 || y >= n {
			return 0
		}
		if max(abs(x-B[0]), abs(y-B[1])) <= steps {
			return int64(V[x][y])
		}
		if dp[x][y][steps] >= 0 {
			return dp[x][y][steps]
		}
		dp[x][y][steps] = 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				tmp := dfs(x+dx, y+dy, steps+1)
				if tmp > dp[x][y][steps] {
					dp[x][y][steps] = tmp
				}
			}
		}
		dp[x][y][steps] = dp[x][y][steps] + int64(V[x][y])

		return dp[x][y][steps]
	}

	return dfs(A[0], A[1], 0) - int64(V[A[0]][A[1]])
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
