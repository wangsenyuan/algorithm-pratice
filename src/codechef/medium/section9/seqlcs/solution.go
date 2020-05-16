package main

import (
	"bufio"
	"fmt"
	"os"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		line := readNNums(reader, 3)
		N, K, L := line[0], line[1], line[2]
		A := readNNums(reader, N)

		fmt.Println(solve(N, K, L, A))
	}
}

const MOD = 1000000007

func solve(N int, K int, L int, A []int) int {

	lcs := make([][]int, 2)
	for i := 0; i < 2; i++ {
		lcs[i] = make([]int, 20)
	}

	next := func(v int, mask int) int {
		var ret int
		var cur int
		for i := 0; i < N; i++ {
			if mask&(1<<uint(i)) > 0 {
				cur++
			}
			lcs[0][i+1] = cur
			lcs[1][i+1] = 0
		}

		for i := 1; i <= N; i++ {
			// normal lcs calculate
			if A[i-1] == v {
				lcs[1][i] = lcs[0][i-1] + 1
			} else {
				lcs[1][i] = max(lcs[0][i], lcs[1][i-1])
			}
			if lcs[1][i] > lcs[1][i-1] {
				ret |= 1 << uint(i-1)
			}
		}
		return ret
	}

	nextMask := make([][]int, K+1)

	for i := 1; i <= K; i++ {
		nextMask[i] = make([]int, 1<<uint(N))
		for j := 0; j < 1<<uint(N); j++ {
			nextMask[i][j] = next(i, j)
		}
	}

	var dp func(n int, mask int) int

	mem := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		mem[i] = make([]int, 1<<uint(N))
		for j := 0; j < 1<<uint(N); j++ {
			mem[i][j] = -1
		}
	}

	dp = func(n int, mask int) int {
		if n > N {
			var cur int
			for i := 0; i < N; i++ {
				if mask&(1<<uint(i)) > 0 {
					cur++
				}
			}

			if cur == L {
				return 1
			}
			return 0
		}

		if mem[n][mask] >= 0 {
			return mem[n][mask]
		}
		mem[n][mask] = 0

		for i := 1; i <= K; i++ {
			mem[n][mask] += dp(n+1, nextMask[i][mask])
			if mem[n][mask] >= MOD {
				mem[n][mask] -= MOD
			}
		}
		return mem[n][mask]
	}

	return dp(1, 0)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
