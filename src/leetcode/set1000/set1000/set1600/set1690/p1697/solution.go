package p1697

import "sort"

const H = 20

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	use := make([]bool, len(edgeList))
	uf := NewUFSet(n)
	degree := make([]int, n)
	for i := 0; i < len(edgeList); i++ {
		a, b := edgeList[i][0], edgeList[i][1]
		if uf.Union(a, b) {
			use[i] = true
			degree[a]++
			degree[b]++
		}
	}

	G := make([][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([]int, 0, degree[i])
	}

	for i := 0; i < len(edgeList); i++ {
		if !use[i] {
			continue
		}
		a, b := edgeList[i][0], edgeList[i][1]
		G[a] = append(G[a], i)
		G[b] = append(G[b], i)
	}

	P := make([][]int, n)
	X := make([][]int, n)
	D := make([]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
		X[i] = make([]int, H)
	}

	edgeLength := make([]int, n)
	edgeLength[0] = 0

	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		X[u][0] = edgeLength[u]

		for i := 1; i < H; i++ {
			X[u][i] = max(X[u][i-1], X[P[u][i-1]][i-1])
		}

		for _, j := range G[u] {
			edge := edgeList[j]
			v := edge[0] ^ edge[1] ^ u
			if v == p {
				continue
			}
			w := edge[2]
			edgeLength[v] = w
			D[v] = D[u] + 1
			dfs(u, v)
		}
	}

	//the tree may not connected
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		pi := uf.Find(i)
		if !vis[pi] {
			vis[pi] = true
			dfs(i, i)
		}
	}

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	getMaxEdgeLengthTill := func(u, p int) int {
		if u == p {
			return 0
		}
		var a int
		for i := H - 1; i >= 0; i-- {
			if D[u]-D[p] >= 1<<i {
				a = max(a, X[u][i])
				u = P[u][i]
			}
		}
		return a
	}

	getMaxEdgeLenthBetween := func(u, v int) int {
		p := lca(u, v)
		a := getMaxEdgeLengthTill(u, p)
		b := getMaxEdgeLengthTill(v, p)
		return max(a, b)
	}

	ans := make([]bool, len(queries))

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		u, v := cur[0], cur[1]

		if uf.Find(u) != uf.Find(v) {
			// u can't reach v
			continue
		}

		lim := cur[2]
		tmp := getMaxEdgeLenthBetween(u, v)
		ans[i] = tmp < lim
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type UFSet struct {
	set []int
	cnt []int
}

func NewUFSet(n int) *UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		cnt[i]++
	}
	return &UFSet{set, cnt}
}

func (uf *UFSet) Find(a int) int {
	if uf.set[a] != a {
		uf.set[a] = uf.Find(uf.set[a])
	}
	return uf.set[a]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.set[pb] = pa
	uf.cnt[pa] += uf.cnt[pb]
	return true
}
