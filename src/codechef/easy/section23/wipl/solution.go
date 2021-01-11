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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, K := readTwoNums(reader)
		H := readNNums(reader, n)
		res := solve(K, H)
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

const INF = 1 << 30

func solve(K int, H []int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	n := len(H)
	suf := make([]int, n+1)
	sort.Ints(H)
	var p int
	for i := n - 1; i >= 0; i-- {
		q := p ^ 1
		for j := 0; j <= K; j++ {
			dp[q][j] = INF
		}
		for j := K; j >= 0; j-- {
			if j <= H[i] {
				dp[q][j] = H[i]
				continue
			}
			// j >= H[i]
			if dp[p][j-H[i]] == INF {
				dp[q][j] = INF
			} else {
				dp[q][j] = min(dp[p][j], dp[p][j-H[i]]+H[i])
			}
		}
		p = q

		suf[i] = suf[i+1] + H[i]
		if suf[i]-dp[p][K] >= K {
			return n - i
		}
	}

	return -1
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
