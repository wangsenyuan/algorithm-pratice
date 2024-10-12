package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(reader, m)
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const inf = 1 << 60

func solve(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	gp := make([][]int, m)
	for j := 0; j < m; j++ {
		gp[j] = make([]int, 1<<n)
		for state := 0; state < 1<<n; state++ {
			gp[j][state] = -inf
			for r := 0; r < n; r++ {
				var sum int
				for i := 0; i < n; i++ {
					if (state>>i)&1 == 1 {
						sum += grid[(r+i)%n][j]
					}
				}
				gp[j][state] = max(gp[j][state], sum)
			}
		}
	}

	dp := make([]int, 1<<n)
	fp := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = -inf
	}
	dp[0] = 0

	for j := 0; j < m; j++ {
		copy(fp, dp)

		for state := 0; state < 1<<n; state++ {
			tmp := (1<<n - 1) ^ state
			for next := tmp; next > 0; next = (next - 1) & tmp {
				fp[next|state] = max(fp[next|state], dp[state]+gp[j][next])
			}
		}
		copy(dp, fp)
	}

	return dp[(1<<n)-1]
}
