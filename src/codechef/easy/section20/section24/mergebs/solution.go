package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readString(reader)
		B := readString(reader)
		res := solve(n, A, B)
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

func solve(n int, A string, B string) int {
	// dp[i][j] = merge A[i:..], B[j...]
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[n][n] = 0

	ca := count(n, A)
	cb := count(n, B)

	for i := n; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i < n {
				if A[i] == '0' {
					dp[i][j] = min(dp[i][j], dp[i+1][j])
				} else {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+ca[i+1]+cb[j])
				}
			}
			if j < n {
				if B[j] == '0' {
					dp[i][j] = min(dp[i][j], dp[i][j+1])
				} else {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+ca[i]+cb[j+1])
				}
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func count(n int, s string) []int {
	res := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		res[i] = 1 - int(s[i]-'0')
		res[i] += res[i+1]
	}
	return res
}
