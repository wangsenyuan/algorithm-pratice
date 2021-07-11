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
		n, P := readTwoNums(reader)
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, 3)
		}
		res := solve(n, P, G)

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

const Y = 2001

func solve(n int, P int, G [][]int) int {
	vecs := make([][]int, Y)

	for y := 0; y < Y; y++ {
		vecs[y] = make([]int, 0, 2)
	}

	for _, g := range G {
		y, e := g[0], g[2]
		vecs[y] = append(vecs[y], e)
	}

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = P + 1
		}
	}

	dp[0][0] = 0
	dp[1][0] = 0

	fp := make([]int, n+1)

	for y := 0; y < Y; y++ {
		for i := 0; i <= n; i++ {
			fp[i] = dp[1][i] + 1
		}

		sort.Ints(vecs[y])

		for i := 1; i <= len(vecs[y]); i++ {
			for j := i; j <= n; j++ {
				dp[0][j] = min(dp[0][j], dp[1][j-i]+vecs[y][i-1])
				fp[j] = min(fp[j], dp[1][j-i]+2*vecs[y][i-1]+1)
			}
		}
		copy(dp[1], fp)
	}
	var ans int
	for ans < n && dp[0][ans+1] <= P {
		ans++
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
