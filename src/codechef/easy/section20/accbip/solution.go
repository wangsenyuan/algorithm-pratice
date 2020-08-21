package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, C, K := readThreeNums(reader)

		L := make([][]int, N)
		for i := 0; i < N; i++ {
			L[i] = readNNums(reader, 3)
		}
		V := readNNums(reader, C)

		fmt.Println(solve(N, C, K, L, V))
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

const INF = 1 << 60

func solve(N, C, K int, L [][]int, V []int) int64 {

	freq := make([]map[int]int, C)
	for i := 0; i < C; i++ {
		freq[i] = make(map[int]int)
	}
	for _, line := range L {
		a, c := line[0], line[2]
		freq[c-1][a]++
	}

	dp := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int64, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = INF
		}
	}

	ans := make([][]int64, C)
	sz := make([]int, C)

	vec := make([]int, N)
	for c := 0; c < C; c++ {
		var p int
		for _, cnt := range freq[c] {
			vec[p] = cnt
			p++
			sz[c] += cnt
		}
		sort.Ints(vec[:p])
		ans[c] = make([]int64, sz[c]+1)
		var y int

		for j := 0; j < p; j++ {
			for vec[j] > 0 {
				var cnt1, cnt2, cnt3 int64
				for jj := 0; jj < p; jj++ {
					z := int64(vec[jj])
					cnt1 += z
					cnt2 += z * z
					cnt3 += z * z * z
				}
				ans[c][y] = (cnt1*cnt1*cnt1 + 2*cnt3 - 3*cnt2*cnt1) / 6
				y++
				vec[j]--
			}
		}
	}

	var ii int

	dp[0][0] = 0

	for i := 0; i < C; i++ {
		for j := 0; j <= K; j++ {
			dp[1-ii][j] = INF
		}
		for j := 0; j <= K; j++ {
			for l := 0; l <= sz[i]; l++ {
				x := int64(j) + int64(l)*int64(V[i])
				if x > int64(K) {
					break
				}
				dp[1-ii][int(x)] = min(dp[1-ii][int(x)], dp[ii][j]+ans[i][l])
			}
		}
		ii = 1 - ii
	}

	var best int64 = INF

	for i := 0; i <= K; i++ {
		best = min(best, dp[ii][i])
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
