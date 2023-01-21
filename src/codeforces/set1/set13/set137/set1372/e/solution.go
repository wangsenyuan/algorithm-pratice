package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	segs := make([][][]int, n)

	for i := 0; i < n; i++ {
		x := readNum(reader)
		segs[i] = make([][]int, x)
		for j := 0; j < x; j++ {
			segs[i][j] = readNNums(reader, 2)
		}
	}
	fmt.Println(solve(n, m, segs))
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

func solve(n, m int, segs [][][]int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	left := make([][]int, n)
	right := make([][]int, n)
	for i := 0; i < n; i++ {
		left[i] = make([]int, m)
		right[i] = make([]int, m)
		for j := 0; j < m; j++ {
			left[i][j] = m
			right[i][j] = m
		}
		ss := segs[i]
		for _, s := range ss {
			for j := s[0] - 1; j <= s[1]-1; j++ {
				left[i][j] = s[0] - 1
				right[i][j] = s[1] - 1
			}
		}
	}

	get := func(l, r int) int {
		if l > r {
			return 0
		}
		return dp[l][r]
	}

	for l := m - 1; l >= 0; l-- {
		for r := l; r < m; r++ {
			for x := l; x <= r; x++ {
				var cnt int
				for i := 0; i < n; i++ {
					// x所在的区间处于 范围 [l...r]中间
					if left[i][x] >= l && right[i][x] <= r {
						cnt++
					}
				}
				dp[l][r] = max(dp[l][r], cnt*cnt+get(l, x-1)+get(x+1, r))
			}
		}
	}
	return dp[0][m-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
