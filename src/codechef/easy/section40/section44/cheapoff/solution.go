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
		n, m, k := readThreeNums(reader)
		A := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(A, k, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const H = 31

func solve(A []int, k int, E [][]int) int {
	if k == 1 {
		return 0
	}
	n := len(A)

	dsu := NewDSU(n)

	var ok bool

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		dsu.Union(u, v)
		if dsu.cnt[dsu.Find(u)] >= k {
			ok = true
		}
	}
	if !ok {
		return -1
	}

	// 如果到某一位d时，所有的A[i]都一致，（能形成至少k个room组成的分组)
	// 但是再第i-1位时，无法找到这样的条件，那是否一定要将该位设置位1吗？
	// 其实时不一定的，因为可能可以在更低的位上满足条件

	// 从低位到高位，按位相同的那些边被选中

	var dfs func(d int, arr []int) int

	dfs = func(d int, arr []int) int {
		// arr = index of edges that have
		if d < 0 {
			return 0
		}
		dsu.Reset()

		a := make([][]int, 2)
		for i := 0; i < 2; i++ {
			a[i] = make([]int, 0, 1)
		}
		yes := make([]bool, n)

		for _, i := range arr {
			e := E[i]
			u, v := e[0]-1, e[1]-1
			if (A[u]>>d)&1 == (A[v]>>d)&1 {
				c := (A[u] >> d) & 1
				a[c] = append(a[c], i)
				dsu.Union(u, v)
				if dsu.cnt[dsu.Find(u)] >= k {
					yes[c] = true
				}
			}
		}

		ans := 1 << 30
		if yes[0] {
			ans = min(ans, dfs(d-1, a[0]))
		}
		if yes[1] {
			ans = min(ans, dfs(d-1, a[1]))
		}
		if !yes[0] && !yes[1] {
			ans = (1 << d) | dfs(d-1, arr)
		}
		return ans
	}

	id := make([]int, len(E))

	for i := 0; i < len(E); i++ {
		id[i] = i
	}

	return dfs(H-1, id)
}

func min(a, b int) int {
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

func (set *DSU) Reset() {
	for i := 0; i < len(set.arr); i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = len(set.arr)
}
