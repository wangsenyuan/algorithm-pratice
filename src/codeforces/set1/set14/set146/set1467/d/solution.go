package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k, m := readThreeNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, k, Q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}
	fmt.Print(buf.String())
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

const MOD = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func solve(A []int, k int, Q [][]int) []int {
	n := len(A)
	// dp[i][j] = count of good paths after j moves ending at i
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}

	for j := 1; j <= k; j++ {
		for i := 0; i < n; i++ {
			if i-1 >= 0 {
				dp[i][j] = add(dp[i][j], dp[i-1][j-1])
			}
			if i+1 < n {
				dp[i][j] = add(dp[i][j], dp[i+1][j-1])
			}
		}
	}

	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			cnt[i] = add(cnt[i], mul(dp[i][j], dp[i][k-j]))
		}
	}

	ans := make([]int, len(Q))

	var sum int

	for i := 0; i < n; i++ {
		sum = add(sum, mul(A[i], cnt[i]))
	}

	for i, cur := range Q {
		j, x := cur[0], cur[1]
		j--
		sum = sub(sum, mul(A[i], cnt[i]))
		A[j] = x
		sum = add(sum, mul(A[i], cnt[i]))
		ans[i] = sum
	}

	return ans
}
