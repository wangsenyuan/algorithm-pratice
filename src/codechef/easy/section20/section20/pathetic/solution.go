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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", res[i])
		}
		fmt.Println()
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

const N = 101
const M = 2000000000000000001

var P []uint64

func init() {
	primes := make([]uint64, 0, 100)
	set := make([]bool, N)

	for i := 2; i < N; i++ {
		if !set[i] {
			primes = append(primes, uint64(i))
			for j := i * i; j < N; j += i {
				set[j] = true
			}
		}
	}

	back := make([]uint64, len(primes))

	for {
		i := len(primes) - 1
		for i > 0 && primes[0] >= M/primes[i] {
			i--
		}
		if i == 0 {
			break
		}
		x := primes[0]
		y := primes[i]
		x *= y
		for j := 1; j < i; j++ {
			back[j-1] = primes[j]
		}
		for j := i + 1; j < len(primes); j++ {
			back[j-2] = primes[j]
		}
		back[len(primes)-2] = x
		copy(primes, back)
		primes = primes[:len(primes)-1]
	}
	P = primes
	// len(P) == 2
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
	}
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}

func solve(n int, E [][]int) []uint64 {
	adj := getGraph(n, E)

	depth := make([]int, n)

	var dfs func(p, u, d int)

	dfs = func(p, u, d int) {
		depth[u] = d
		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs(u, v, d+1)
		}
	}
	dfs(-1, 0, 0)

	res := make([]uint64, n)
	for i := 0; i < n; i++ {
		res[i] = P[depth[i]&1]
	}

	return res
}
