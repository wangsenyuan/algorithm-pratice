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
		n, m, k := readThreeNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, E, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
func solve(n int, E [][]int, k int) int64 {
	// 如果都小于等于k，那么可以 k - max(w)
	// 如果存在大于k的部分，那么就需要把这些大的部分给（如果选如的话）就需要降到k
	sort.Slice(E, func(i, j int) bool {
		return E[i][2] < E[j][2]
	})

	// 先把所有小于等于k的边给连起来

	set := NewDSU(n)
	id := 0
	var max_edge int
	for id < len(E) && E[id][2] <= k {
		e := E[id]
		u, v := e[0]-1, e[1]-1
		set.Union(u, v)
		max_edge = E[id][2]
		id++
	}
	// 如果使用了较大速度的边，就必须降速
	var cost1 int64 = 1 << 60
	if id > 0 {
		cost1 = int64(k - max_edge)

		if set.size > 1 {
			// it will use some higher edge
			cost1 = 0

			for i := id; i < len(E); i++ {
				e := E[i]
				u, v, w := e[0]-1, e[1]-1, e[2]
				if set.Union(u, v) {
					cost1 += int64(w - k)
				}
			}
		}
	}
	var cost2 int64 = 1 << 60

	if id < len(E) {
		set2 := NewDSU(n)
		set2.Union(E[id][0]-1, E[id][1]-1)
		cost2 = int64(E[id][2] - k)
		for i := 0; i < len(E); i++ {
			e := E[i]
			u, v, w := e[0]-1, e[1]-1, e[2]
			if set2.Union(u, v) && w > k {
				cost2 += int64(w - k)
			}
		}
	}

	return min(cost1, cost2)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

func (this *DSU) Copy() *DSU {
	that := NewDSU(len(this.arr))
	copy(that.arr, this.arr)
	copy(that.cnt, this.cnt)
	that.size = this.size
	return that
}
