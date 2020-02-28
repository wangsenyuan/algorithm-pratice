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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readNInts(scanner *bufio.Scanner, sizeHint int) []int {
	res := make([]int, 0, sizeHint)
	x := 0
	scanner.Scan()
	for i := 0; i < sizeHint; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}

		if x == len(scanner.Bytes()) {
			break
		}

		var num int

		x = readInt(scanner.Bytes(), x, &num)

		res = append(res, num)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, H := readTwoNums(scanner)
		T := readNNums(scanner, N)
		children := make([][]int, N)

		for i := 0; i < N; i++ {
			children[i] = readNInts(scanner, 11)
		}
		fmt.Println(solve(N, H, T, children))
	}
}

const INF = 1 << 30

func solve(n int, H int, T []int, children [][]int) int {

	// dfs u returns the end time of sub tree rooted at u
	//dp[i][x] = sub tree (rooted at i) starts from hour x, the ending time

	dp := make([][]*Pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]*Pair, H+1)
	}

	M := 1 << 10

	F := make([][]*Pair, M)

	for i := 0; i < M; i++ {
		F[i] = make([]*Pair, H+1)
	}

	var dfs func(u int)

	dfs = func(u int) {

		m := children[u][0]

		M := 1 << uint(m)

		for i := 0; i < m; i++ {
			dfs(children[u][i+1] - 1)
		}

		for i := 0; i < (1 << uint(m)); i++ {
			for j := 0; j <= H; j++ {
				F[i][j] = NewPair(1000, H)
			}
		}

		t := T[u]

		for p := 0; p+t <= H; p++ {
			F[0][p] = NewPair(0, p+t)
		}

		for p := H + 1 - t; p <= H; p++ {
			F[0][p] = NewPair(1, t)
		}

		for state := 0; state < M; state++ {
			for j := 0; j < m; j++ {
				if state&(1<<uint(j)) == 0 {
					v := children[u][j+1] - 1
					for p := 0; p <= H; p++ {
						a := F[state][p]

						c := dp[v][a.second]

						F[state|1<<uint(j)][p] = minPair(F[state|1<<uint(j)][p], NewPair(a.first+c.first, c.second))
					}
				}
			}
		}

		for p := 0; p <= H; p++ {
			dp[u][p] = F[M-1][p]
		}
	}

	dfs(0)

	return dp[0][H].first
}

func minPair(a, b *Pair) *Pair {
	if a.first > b.first {
		return b
	}
	return a
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

func NewPair(first, second int) *Pair {
	pair := new(Pair)
	pair.first = first
	pair.second = second
	return pair
}
