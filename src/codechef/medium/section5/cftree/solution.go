package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N, M := readTwoNums(scanner)
	parent := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		parent[i] = readNum(scanner)
	}

	queries := make([][]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		if bs[0] == 'U' {
			var x, k int
			pos := readInt(bs, 2, &x)
			readInt(bs, pos+1, &k)
			queries[i] = []int{0, x, k}
		} else {
			var x int
			readInt(bs, 2, &x)
			queries[i] = []int{1, x}
		}
	}
	res := solve(N, M, parent, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

const MOD = 1e9 + 7
const MAX_N = 1e5 + 10

var F []int64

func init() {
	F = make([]int64, 2*MAX_N)
	F[0] = 0
	F[1] = 1

	for i := 2; i < 2*MAX_N; i++ {
		F[i] = F[i-1] + F[i-2]
		if F[i] >= MOD {
			F[i] -= MOD
		}
	}
}

func solve(N, M int, parent []int, queries [][]int) []int {
	sqrtn := int(math.Sqrt(float64(N)))

	children := make([][]int, N)
	for i := 0; i < N; i++ {
		children[i] = make([]int, 0, 3)
	}

	for i := 0; i < N-1; i++ {
		p := parent[i] - 1
		children[p] = append(children[p], i+1)
	}

	start := make([]int, N)
	end := make([]int, N)
	depth := make([]int, N)
	val := make([]int64, N)
	var dfs func(u int, d int, time *int)

	dfs = func(u int, d int, time *int) {
		depth[u] = d
		start[u] = *time
		*time++
		for _, v := range children[u] {
			dfs(v, d+1, time)
		}
		end[u] = *time
	}

	dfs(0, 0, new(int))

	updates := make([]Update, sqrtn)
	var ui int

	nodeUpdates := make([][]int, N)
	for i := 0; i < N; i++ {
		nodeUpdates[i] = make([]int, 0, 3)
	}

	var rebuild func(v int, current, prev int64)

	rebuild = func(v int, current, prev int64) {
		var upd int64

		for i := 0; i < len(nodeUpdates[v]); i++ {
			modAdd(&current, F[nodeUpdates[v][i]])
			modAdd(&upd, F[nodeUpdates[v][i]-1])
		}
		nodeUpdates[v] = nodeUpdates[v][:0]
		modAdd(&val[v], current)

		var newCurrent int64

		modAdd(&newCurrent, current)
		modAdd(&newCurrent, upd)
		modAdd(&newCurrent, prev)

		for _, c := range children[v] {
			rebuild(c, newCurrent, current)
		}
	}

	doUpdate := func(x, k int) {
		// record the update
		updates[ui] = Update{x, k}
		nodeUpdates[x] = append(nodeUpdates[x], k)
		ui++
		if ui < sqrtn {
			return
		}
		//rebuild the tree
		rebuild(0, 0, 0)
		ui = 0
	}

	doQuery := func(x int) int {
		res := val[x]

		for i := 0; i < ui; i++ {
			upd := updates[i]
			j, k := upd.x, upd.k
			if start[j] <= start[x] && end[x] <= end[j] {
				// x is children of j
				modAdd(&res, F[k+depth[x]-depth[j]])
			}
		}

		return int(res)
	}

	res := make([]int, M)
	var ri int

	for i := 0; i < M; i++ {
		query := queries[i]
		if query[0] == 0 {
			// update
			doUpdate(query[1]-1, query[2])
		} else {
			// query
			res[ri] = doQuery(query[1] - 1)
			ri++
		}
	}

	return res[:ri]
}

type Update struct {
	x, k int
}

func modAdd(res *int64, val int64) {
	*res += val
	if *res >= MOD {
		*res -= MOD
	}
}
