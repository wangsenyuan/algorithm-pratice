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
		s, _ := reader.ReadBytes('\n')
		var n uint64
		pos := readUint64(s, 0, &n) + 1
		var m int
		readInt(s, pos, &m)

		res := solve(int64(n), m)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const MOD = 1000000007

// ap[i][x] 顶点至少为x时, 树的大小为 (1 << i) - 1, 且 level 1 < level 2 > level 3 ...
// bp[i][x] 顶点至多为x时，树的大小为 (1 << i) - 1, 且 level 1 > level 2 < level 3 ...
// dp[i][x] = 为顶点为x, 树的大小为 (1 << i) - 1, 且 level 1 < level 2 > level 3
// dp[i][x] = bp[i-1][x-1] ** 2
// fp[i][x] = 顶点为x，且树的大小为 (1 << i) - 1, 且 level 1 > level 2 < level 3 ...
// fp[i][x] = ap[i-1][x+1] ** 2

func solve(n int64, M int) int {
	// n == 1 << h - 1
	n++
	var h int
	for 1<<uint64(h+1) <= n {
		h++
	}

	H := h + 1

	ap := make([]int64, M+2)
	bp := make([]int64, M+2)
	dp := make([]int64, M+2)
	fp := make([]int64, M+2)

	for x := 1; x <= M; x++ {
		bp[x] = 1 + bp[x-1]
	}
	for x := M; x >= 1; x-- {
		ap[x] = 1 + ap[x+1]
	}

	for i := 2; i < H; i++ {
		for x := 1; x <= M; x++ {
			dp[x] = bp[x-1] * bp[x-1] % MOD
		}
		for x := M; x >= 1; x-- {
			fp[x] = ap[x+1] * ap[x+1] % MOD
		}
		for x := M; x >= 1; x-- {
			ap[x] = (dp[x] + ap[x+1]) % MOD
		}
		for x := 1; x <= M; x++ {
			bp[x] = (fp[x] + bp[x-1]) % MOD
		}
	}

	return int(ap[1])
}
