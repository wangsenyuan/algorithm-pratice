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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func solve(A []int) int {
	// |xi - mex(x1..xi)| <= 1
	// mex(x1...xi) = xi-1 or xi + 1
	// mex(x1...xi) = xi-1 => x1...xi-1 都必须出现至少一次， 因为xi肯定出现了， 剩下的i-1 必须包括 0...xi-2
	// 所以必须是 0,1,2,... xi-2, xi
	// 如果 mex(x1...xi) = 0.1,2...xi => xi+1
	// dp[x] = 到目前为止mex = x的子序列的数量
	// ai = x + 2, dp[x] *= 2, 如果ai = x + 1, dp[x+1] += dp[x]
	n := len(A)
	dp := make([]int, n+2)
	fp := make([]int, n+2)
	dp[0] = 1

	for _, x := range A {
		dp[x+1] = modMul(dp[x+1], 2)
		dp[x+1] = modAdd(dp[x+1], dp[x])
		if x > 0 {
			fp[x-1] = modAdd(fp[x-1], fp[x-1])
			fp[x-1] = modAdd(fp[x-1], dp[x-1])
		}
		fp[x+1] = modAdd(fp[x+1], fp[x+1])
	}

	var res int

	for _, x := range dp {
		res = modAdd(res, x)
	}

	for _, x := range fp {
		res = modAdd(res, x)
	}

	return modAdd(res, MOD-1)
}
