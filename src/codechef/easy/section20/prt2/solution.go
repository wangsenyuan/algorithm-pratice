package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		fmt.Println(solve(n, k, A, E))
	}
}

func solve(n, k int, A []int, E [][]int) int64 {
	adj := getGraph(n, E)

	var z = n
	h1 := make([]int, n)
	h2 := make([]int, n)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		h1[u] = n
		h2[u] = n
		if len(adj[u]) == 1 {
			// a leaf
			h1[u] = 0
		}
		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs(u, v)
			if h1[v]+1 < h1[u] {
				h2[u] = h1[u]
				h1[u] = h1[v] + 1
			} else if h1[v]+1 < h2[u] {
				h2[u] = h1[v] + 1
			}
		}
		z = min(z, h1[u]+h2[u])
	}

	dfs(-1, 0)
	// add the center node (having h1[r] + h2[r] == z) back
	z++

	sort.Ints(A)
	K := int64(k)
	Z := int64(z)

	var ans int64
	for i := n - 1; i >= n-z; i-- {
		cnt := (K / (2 * Z)) * 2
		if k%(2*z) == 2 {
			if i >= n-2 {
				cnt++
			}
		} else {
			if i >= n-(k%(2*z)/2) {
				cnt += 2
			} else if i == n-(k%(2*z)/2)-1 && (k%(2*z))%2 == 1 {
				cnt++
			}
		}
		ans += int64(cnt) * int64(A[i])
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
