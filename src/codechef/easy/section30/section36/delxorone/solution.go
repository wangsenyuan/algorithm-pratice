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

func solve(A []int) int {
	//A[i] ^ A[i+1] <= 1
	// A[i] = A[i+1] or A[i] ^ A[i+1] = 1
	n := len(A)
	// dp[i] = max nums kept in first i items, with condition, and leave A[i]
	cnt := make([]int, n+1)

	for i := 0; i < n; i++ {
		x := A[i]
		if x&1 == 1 {
			x ^= 1
		}
		cnt[x]++
	}

	var res int
	for i := 0; i <= n; i++ {
		res = max(res, cnt[i])
	}

	return n - res
}

func solve1(A []int) int {
	//A[i] ^ A[i+1] <= 1
	// A[i] = A[i+1] or A[i] ^ A[i+1] = 1
	n := len(A)
	// dp[i] = max nums kept in first i items, with condition, and leave A[i]
	dp := make([]int, n+1)
	pos := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = 1
		pos[i] = -1
	}

	for i := 0; i < n; i++ {
		if pos[A[i]] >= 0 {
			dp[i] = dp[pos[A[i]]] + 1
		}
		x := (A[i] ^ 1)
		if x <= n && pos[x] >= 0 {
			dp[i] = max(dp[i], dp[pos[x]]+1)
		}
		pos[A[i]] = i
	}

	var res int
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}

	return n - res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
