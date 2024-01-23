package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, s := readTwoNums(reader)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 3)
		}
		res := solve(n, s, edges)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const mod = 998244353

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, s int, edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// 可以倒过来进行处理
	uf := NewUFSet(n)

	res := 1

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		// 以u, v为最大边
		u = uf.Find(u)
		v = uf.Find(v)
		su := uf.cnt[u]
		sv := uf.cnt[v]
		// 如果存在一条边在u的子树和其他的节点中间为w + 1, 是不可能被选中的
		uf.Union(u, v)
		// 此时，如果存在一条边的长度为w + 1, 去连接u,v子树，是不能被选中的
		// 不包括u, v 中间的边
		tmp := su*sv - 1
		// 不存在边，或者存在一条w + 1, ... s 的边
		tmp = pow(s-w+1, tmp)
		//tmp = add(tmp, 1)
		res = mul(res, tmp)
	}

	return res
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

type UFSet struct {
	arr  []int
	cnt  []int
	size int
}

func NewUFSet(n int) UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UFSet{arr, cnt, n}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
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
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	uf.size--
	return true
}
