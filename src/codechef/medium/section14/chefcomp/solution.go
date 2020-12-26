package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		P := readNNums(reader, n)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, E, P, A, B)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(n int, E [][]int, P []int, A []int, B []int) []int {
	G := getGraph(n, E)

	nt := make([]bool, n)

	parent := make([]int, n)

	cache := make([]int, n)
	day := make([]int, n)

	for i := 0; i < n; i++ {
		cache[i] = i
		day[P[i]-1] = i + 1
	}
	uf := NewUF(n)

	for i := n - 1; i >= 0; i-- {
		u := P[i] - 1
		parent[u] = u
		nt[u] = true
		for _, v := range G[u] {
			if nt[v] {
				parent[cache[uf.Find(v)]] = u
				uf.Union(u, v)
				cache[uf.Find(u)] = u
			}
		}
	}

	V := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		if parent[i] == i {
			continue
		}
		degree[i]++
		degree[parent[i]]++
	}
	for i := 0; i < n; i++ {
		V[i] = make([]int, 0, degree[i])
	}
	var root int
	for i := 0; i < n; i++ {
		if parent[i] != i {
			V[i] = append(V[i], parent[i])
			V[parent[i]] = append(V[parent[i]], i)
		} else {
			root = i
		}
	}

	sum := make([]int, n+1)
	days := make([]int, n+1)
	var dfs func(p, u, i int)

	ans := make([]int, n)

	dfs = func(p, u, i int) {
		sum[i] = sum[i-1] + A[u]
		days[i] = day[u]
		ans[u] = -1

		if sum[i] >= B[u] {
			k := search(i, func(k int) bool {
				return sum[k] >= B[u]
			})
			ans[u] = days[k]
		}

		for _, v := range V[u] {
			if p == v {
				continue
			}
			dfs(u, v, i+1)
		}
	}

	dfs(-1, root, 1)

	return ans
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &UF{arr, cnt}
}

func (this *UF) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *UF) Union(a, b int) bool {
	pa := this.Find(a)
	pb := this.Find(b)
	if pa == pb {
		return false
	}
	if this.cnt[pa] < this.cnt[pb] {
		pa, pb = pb, pa
	}
	this.cnt[pa] += this.cnt[pb]
	this.arr[pb] = pa
	return true
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
