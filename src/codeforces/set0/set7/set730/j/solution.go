package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	a, b := solve(A, B)
	fmt.Println(fmt.Sprintf("%d %d\n", a, b))
}

const inf = 1 << 20

func solve(A []int, B []int) (int, int) {
	n := len(A)

	P := make([]Pair, n)
	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], B[i]}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].second > P[j].second
	})

	var sa int

	for i := 0; i < n; i++ {
		sa += P[i].first
	}

	var sb int
	var m int
	for i := 0; i < n; i++ {
		sb += P[i].second
		if sb >= sa && m == 0 {
			m = i + 1
		}
	}
	dp := make([][]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = make([]int, sb+1)
		for k := 0; k <= sb; k++ {
			dp[j][k] = -inf
		}
	}
	dp[0][0] = 0

	var sum int
	for i := n - 1; i >= 0; i-- {
		a := P[i].first
		b := P[i].second
		sum += b
		for j := m; j > 0; j-- {
			for k := sum; k >= b; k-- {
				dp[j][k] = max(dp[j][k], dp[j-1][k-b]+a)
			}
		}
	}

	var best int
	for _, x := range dp[m][sa:] {
		best = max(best, x)
	}
	return m, sa - best
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
