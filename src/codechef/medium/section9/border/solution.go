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
		_, q := readTwoNums(reader)

		s, _ := reader.ReadString('\n')

		solver := NewSolver(s[:len(s)-1])

		for q > 0 {
			q--
			i, k := readTwoNums(reader)

			fmt.Println(solver.Query(i, k))
		}
	}
}

type Solver struct {
	cnt []int
	dp  [][]int
}

const H = 20

func NewSolver(s string) Solver {
	n := len(s)
	F := make([]int, n)

	for i := 1; i < n; i++ {
		F[i] = F[i-1]

		for s[i] != s[F[i]] {
			F[i]--
			if F[i] < 0 {
				break
			}
			F[i] = F[F[i]]
		}
		F[i]++
	}

	cnt := make([]int, n)
	dp := make([][]int, H)

	for i := 0; i < H; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		if F[i] == 0 {
			cnt[i] = 1
			dp[0][i] = i
		} else {
			cnt[i] = cnt[F[i]-1] + 1
			dp[0][i] = F[i] - 1
		}
	}

	for i := 1; i < H; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	return Solver{cnt, dp}
}

func (solver Solver) Query(i int, k int) int {
	i--
	if solver.cnt[i] < k {
		return -1
	}
	k = solver.cnt[i] - k
	var st int

	for k > 0 {
		if k&1 == 1 {
			i = solver.dp[st][i]
		}
		st++
		k >>= 1
	}
	return i + 1
}
