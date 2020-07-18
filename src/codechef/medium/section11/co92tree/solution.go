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
		fmt.Println(solve(n, E))
	}
}

func solve(n int, E [][]int) int64 {
	degree := make([]int, n+1)

	for _, e := range E {
		u, v := e[0], e[1]
		degree[u]++
		degree[v]++
	}

	conn := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		conn[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
	}

	parent := make([]int, n+1)

	dfs(conn, 1, 1, parent)

	var ans int64 = int64(n)

	uf := NewUF(n + 1)

	for u := 1; u <= n; u++ {
		var cnt int = 1

		for v := u; v <= n; v = (v + 1) | u {
			w := parent[v]
			if w&u == u {
				// can reach parent
				if uf.Union(v, w) {
					cnt = max(cnt, uf.Count(uf.Find(w)))
				}
			}
		}

		for v := u; v <= n; v = (v + 1) | u {
			uf.Reset(v)
		}

		tmp := int64(u) * int64(cnt)

		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func dfs(conn [][]int, p, u int, parent []int) {
	parent[u] = p

	for _, v := range conn[u] {
		if p == v {
			continue
		}
		dfs(conn, u, v, parent)
	}
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(size int) UF {
	arr := make([]int, size)
	cnt := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UF{arr, cnt}
}

func (uf *UF) Reset(i int) {
	uf.arr[i] = i
	uf.cnt[i] = 1
}

func (uf UF) Count(x int) int {
	return uf.cnt[x]
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] >= uf.cnt[pb] {
		uf.arr[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	} else {
		uf.arr[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	}
	return true
}
