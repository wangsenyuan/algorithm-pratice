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
		n, q := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(n, q, E, queries)

		for _, ans := range res {
			fmt.Println(ans)
		}
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

const H = 20

func solve(N int, Q int, E [][]int, QR [][]int) []int64 {
	adj := createAdjMatrix(N, E)

	PP := make([][]int, N)
	for i := 0; i < N; i++ {
		PP[i] = make([]int, H)
		for j := 0; j < H; j++ {
			PP[i][j] = -1
		}
	}

	ans3 := make([]int64, N)
	dep := make([]int, N)
	subtree := make([]int64, N)
	ans2 := make([]int64, N)
	ans1 := make([]int64, N)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		PP[u][0] = p
		ans3[u] = 0
		if p >= 0 {
			dep[u] = dep[p] + 1
		}
		subtree[u] = 1

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(u, v)
			subtree[u] += subtree[v]
		}

		ans2[u] = subtree[u] * (subtree[u] + 1)

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			ans2[u] -= subtree[v] * (subtree[v] + 1)
		}

		ans2[u] /= 2

		ans1[u] = ans2[u] + subtree[u]*(int64(N)-subtree[u])
	}

	dfs(-1, 0)

	for h := 1; h < H; h++ {
		for i := 0; i < N; i++ {
			p := PP[i][h-1]
			if p >= 0 {
				PP[i][h] = PP[p][h-1]
			}
		}
	}

	preans2 := make([]int64, N)
	preans3 := make([]int64, N)

	var ddfs func(p, u int)

	ddfs = func(p, u int) {
		preans2[u] = ans2[u]
		preans3[u] = ans3[u]
		if p >= 0 {
			preans2[u] += preans2[p]
			preans3[u] += preans3[p]
		}

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			ans3[v] = (subtree[u] - subtree[v]) * subtree[v]
			ddfs(u, v)
		}
	}

	ddfs(-1, 0)

	lca := func(u, v int) int {
		if dep[u] > dep[v] {
			u, v = v, u
		}
		for h := H - 1; h >= 0; h-- {
			if dep[v]-(1<<uint(h)) >= dep[u] {
				v = PP[v][h]
			}
		}
		if u == v {
			return u
		}

		for h := H - 1; h >= 0; h-- {
			if PP[u][h] != PP[v][h] {
				u = PP[u][h]
				v = PP[v][h]
			}
		}
		return PP[u][0]
	}

	getDeepParent := func(u int, deep int) int {
		for h := H - 1; h >= 0; h-- {
			if dep[u]-(1<<uint(h)) >= deep {
				u = PP[u][h]
			}
		}
		return u
	}

	find := func(u, v int) int64 {
		foo := getDeepParent(v, dep[u]+1)
		wow := preans2[v] - preans2[u]
		wow -= preans3[v] - preans3[foo]
		wow += ans1[u]
		wow -= (subtree[foo]) * (int64(N) - subtree[foo])

		return wow
	}

	ans := make([]int64, Q)

	for i := 0; i < Q; i++ {
		u, v := QR[i][0], QR[i][1]
		u--
		v--
		if dep[u] > dep[v] {
			u, v = v, u
		}

		if u == v {
			ans[i] = ans1[u]
			continue
		}

		gg := lca(u, v)

		if gg == u {
			ans[i] = find(u, v)
			continue
		}

		foo1 := getDeepParent(v, dep[gg]+1)
		foo2 := getDeepParent(u, dep[gg]+1)
		ans[i] = find(gg, v) + find(gg, u) + subtree[foo1]*subtree[foo2] - ans1[gg]
	}

	return ans
}

func createAdjMatrix(N int, E [][]int) [][]int {
	degree := make([]int, N)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		degree[u]++
		degree[v]++
	}

	adj := make([][]int, N)

	for i := 0; i < N; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}
