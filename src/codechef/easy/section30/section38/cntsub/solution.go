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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' || s[i] == '\r' {
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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(P []int) int {
	n := len(P)
	dp := make([]int, n+1)

	var res int

	// mask := (1 << 30) - 1

	for i := n - 1; i >= 0; i-- {
		for x := P[i]; x > 0; x = (x - 1) & P[i] {
			dp[P[i]] = modAdd(dp[P[i]], dp[x])
		}
		dp[P[i]] = modAdd(dp[P[i]], 1)
		res = modAdd(res, dp[P[i]])
	}

	return res
}

func solve1(P []int) int {
	n := len(P)

	var h int
	for (1 << h) < n {
		h++
	}

	N := 1 << h
	dp := make([]int, N+1)

	var res int

	// mask := (1 << 30) - 1

	for i := 0; i < n; i++ {
		for x := P[i]; x <= N; x = (x + 1) | P[i] {
			dp[P[i]] = modAdd(dp[P[i]], dp[x])
		}
		dp[P[i]] = modAdd(dp[P[i]], 1)
		res = modAdd(res, dp[P[i]])
	}

	return res
}
