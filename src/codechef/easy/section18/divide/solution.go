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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R *= A
			R %= MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}

	return int(R)
}

func calSimilarity(a, b int) int {
	ab := pow(a, b)
	ba := pow(b, a)
	return min(ab, ba)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve(n int, A []int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := calSimilarity(A[i], A[j])
			dp[i][j] = x
			dp[j][i] = x
		}
	}

	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 10)
	}
	color := make([]int, n)

	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if color[u] == c {
			return true
		}
		if color[u] != 0 {
			return false
		}
		color[u] = c

		for _, v := range conns[u] {
			if !dfs(v, -c) {
				return false
			}
		}

		return true
	}

	check := func(X int) bool {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if dp[i][j] > X {
					conns[i] = append(conns[i], j)
					conns[j] = append(conns[j], i)
				}
			}
		}
		can := true
		for i := 0; i < n; i++ {
			if color[i] == 0 {
				if !dfs(i, 1) {
					can = false
					break
				}
			}
		}

		for i := 0; i < n; i++ {
			color[i] = 0
			conns[i] = conns[i][:0]
		}
		return can
	}

	left, right := 0, MOD

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
