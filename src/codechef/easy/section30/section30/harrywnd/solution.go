package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	N, k := readTwoNums(reader)
	wands := readNNums(reader, k)
	min_res, max_res := solve(wands, N)

	fmt.Printf("%d %d\n", min_res, max_res)
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

const INF = 1 << 29

func solve(wands []int, N int) (int, int) {
	// N <= 10000
	k := len(wands)
	dp := make([]int, N+1)
	fp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = INF
		fp[i] = -INF
	}

	dp[0] = 0
	fp[0] = 0

	for l := 1; l <= N; l++ {
		for i := 0; i < k; i++ {
			if l >= wands[i] {
				dp[l] = min(dp[l], dp[l-wands[i]]+1)
				fp[l] = max(fp[l], fp[l-wands[i]]+1)
			}
		}
	}

	min_cost := dp[N]
	max_cost := fp[N]
	if min_cost >= INF {
		return -1, -1
	}
	return min_cost, max_cost
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
