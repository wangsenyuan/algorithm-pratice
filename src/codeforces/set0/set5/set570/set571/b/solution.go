package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)

	fmt.Println(res)
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

const inf = 1 << 50

func solve(a []int, k int) int {
	sort.Ints(a)
	n := len(a)
	part := n / k
	large := n % k
	dp := make([][]int, large+1)
	for i := 0; i <= large; i++ {
		dp[i] = make([]int, k-large+1)
		for j := 0; j <= k-large; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := 0; i <= large; i++ {
		for j := 0; j <= k-large; j++ {
			nex := i*(part+1) + j*part
			if i < large {
				add := a[nex+part] - a[nex]
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+add)
			}
			if j < k-large {
				add := a[nex+part-1] - a[nex]
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+add)
			}
		}
	}

	return dp[large][k-large]
}
