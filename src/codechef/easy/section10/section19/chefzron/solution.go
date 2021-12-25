package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
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

func solve(n int, A []int) int {
	pos := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if A[i] == 1 {
			pos = append(pos, i)
		}
	}
	if len(pos) == 0 {
		return 0
	}
	if len(pos) == 1 {
		return -1
	}
	back := make([]int, len(pos))
	dp := make([]int, len(pos)+1)
	var best = n
	for i := 0; i < 3; i++ {
		best = min(best, solveDirectly(pos, dp))
		copy(back, pos[1:])
		back[len(pos)-1] = pos[0] + n
		copy(pos, back)
	}
	return best
}

func solveDirectly(arr []int, dp []int) int {
	dp[0] = 0
	dp[1] = INF

	for i := 2; i <= len(arr); i++ {
		dp[i] = arr[i-1] - arr[i-2] + dp[i-2]
		if i >= 3 {
			dp[i] = min(dp[i], arr[i-1]-arr[i-3]+dp[i-3])
		}
	}

	return dp[len(arr)]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
