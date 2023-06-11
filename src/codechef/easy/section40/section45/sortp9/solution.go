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

	tc := 1
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const H = 20

func solve(A []int) []int {
	// A[i] & A[j] = 0, then can swap them
	// 如果 A[i] & A[j] = 0, A[j] & A[k] = 0, A[i] & A[k] 却不一定为0
	// 如果 A[i] & A[k] != 0, 但是如果存在某个j,A[i] & A[j] = 0, A[j] & A[k] = 0, 也可以交换它们
	// 如果 a & b != 0, a & c = 0, b & d = 0, and c & d = 0
	// 能否交换  a, b?
	// 也就是说它们组成了多个独立集，在独立集里面可以升序排列
	// 然后如何进行合并？
	//n := len(A)
	N := 1 << H
	where := make([][]int, N)

	for i, a := range A {
		if len(where[a]) == 0 {
			where[a] = make([]int, 0, 1)
		}
		where[a] = append(where[a], i)
	}
	dsu := NewDSU(N)

	vis := make([]bool, N*2)

	var dfs func(root int, u int)

	dfs = func(root int, u int) {
		if vis[u] {
			return
		}
		vis[u] = true
		if u&N > 0 {
			u ^= N
			dsu.Union(u, root)
			dfs(root, u)
			return
		}
		for i := 0; i < H; i++ {
			if (u>>i)&1 == 0 {
				dfs(root, u|(1<<i))
			}
		}
		flip := N - 1 - u
		if len(where[flip]) > 0 {
			dfs(root, flip|N)
		}
	}

	for i := 0; i < N; i++ {
		if len(where[i]) > 0 && !vis[i|N] {
			dfs(i, i|N)
		}
	}

	bunch := make(map[int][]int)

	for i := 0; i < N; i++ {
		if len(where[i]) > 0 {
			j := dsu.Find(i)
			if len(bunch[j]) == 0 {
				bunch[j] = make([]int, 0, 1)
			}
			bunch[j] = append(bunch[j], i)
		}
	}

	res := make([]int, len(A))
	for _, v := range bunch {
		var id []int
		var val []int
		for _, j := range v {
			for _, k := range where[j] {
				id = append(id, k)
				val = append(val, j)
			}
		}
		sort.Ints(id)
		sort.Ints(val)

		for i := 0; i < len(id); i++ {
			res[id[i]] = val[i]
		}
	}

	return res
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
