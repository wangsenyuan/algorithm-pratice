package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNInt64s(reader, n)

	solver := NewSolver(A)

	q := readNum(reader)
	var buf bytes.Buffer
	for q > 0 {
		q--
		k := readNum(reader)
		T := readNInt64s(reader, n)
		ans := solver.Query(k, T)

		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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
		if s[i] == '\n' || s[i] == '\r' {
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

type Solver struct {
	A []int64
}

func NewSolver(A []int64) Solver {
	return Solver{A}
}

const H = 55

func (solver Solver) Query(K int, T []int64) int64 {

	n := len(T)

	var X []int64

	for i := 0; i < n; i++ {
		if T[i] > 0 {
			T[i] *= solver.A[i]
			X = append(X, T[i])
		}

	}
	if K > len(X) {
		return 0
	}
	n = len(X)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, K+1)
	}

	check := func(sum int64) bool {
		// can we get sum by partition k
		for i := 0; i <= n; i++ {
			for j := 0; j <= K; j++ {
				dp[i][j] = false
			}
		}
		dp[0][0] = true
		for i := 1; i <= n; i++ {
			var tmp int64
			for j := i; j > 0; j-- {
				tmp += X[j-1]
				if sum&tmp == sum {
					for t := 1; t <= K; t++ {
						dp[i][t] = dp[i][t] || dp[j-1][t-1]
					}
				}
			}
		}
		return dp[n][K]
	}
	var l, r int64 = 0, 1 << uint64(H)
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func (solver Solver) Query1(K int, T []int64) int64 {
	// T[i] <= 10
	// and count(T, > 0) < 50
	// split it into K clusters
	// for each cluster, S = T[i] * A[i]
	// find max S &&

	n := len(T)

	var X []int64

	for i := 0; i < n; i++ {
		if T[i] > 0 {
			T[i] *= solver.A[i]
			X = append(X, T[i])
		}

	}
	if K > len(X) {
		return 0
	}
	// K <= not_zero_cnt
	// 对于i，有两个选择
	//      1， 开始一个新的cluster，
	//      2， 放入前面的cluster， S[]
	n = len(X)
	// 还需要一颗trie
	dp := make([][]map[int64]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]map[int64]bool, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = make(map[int64]bool)
		}
	}

	for i := 0; i < n; i++ {
		for k := 1; k <= K && k <= i+1; k++ {
			var sum int64
			for j := i; j >= 0; j-- {
				sum += X[j]
				if j == 0 {
					if k == 1 {
						dp[i][k][sum] = true
					}
				} else {
					for tmp := range dp[j-1][k-1] {
						dp[i][k][tmp&sum] = true
					}
				}
			}
		}
	}
	var best int64
	for k := range dp[n-1][K] {
		best = max(best, k)
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
