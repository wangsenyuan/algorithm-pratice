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
		n, x := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, x)
		buf.WriteString(fmt.Sprintf("%s\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string, k int) string {
	n := len(s)
	dp := make([]string, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = string(byte('z' + 1))
	}

	for i := 0; i < n; i++ {
		c := int(s[i] - 'a')
		nc := min(c, (c+1)%k, (c+k-1)%k)
		dp[i+1] = min(dp[i+1], dp[i]+string(byte('a'+nc)))
		if i > 0 {
			dp[i+1] = min(dp[i+1], dp[i-1]+string(byte('a'+nc))+s[i-1:i])
			dp[i+1] = min(dp[i+1], dp[i][:i-1]+s[i:i+1]+dp[i][len(dp[i])-1:])
		}
		if i > 1 {
			dp[i+1] = min(dp[i+1], dp[i-1][:i-2]+s[i:i+1]+dp[i-1][len(dp[i-1])-1:]+s[i-1:i])
		}
	}
	return dp[n]
}

type Comparable interface {
	int | string
}

func min[T Comparable](arr ...T) T {
	res := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}
	return res
}
