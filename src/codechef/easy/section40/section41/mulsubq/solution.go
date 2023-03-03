package main

import (
	"sort"
)

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
		if buf.Len() > 10000 {
			fmt.Print(buf.String())
			buf.Reset()
		}
	}

	if buf.Len() > 0 {
		fmt.Print(buf.String())
	}
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

func solve(A []int, Q [][]int) []int64 {
	n := len(A)
	arr := make([]int, 20)
	check := func(cnt map[int]int) bool {
		if len(cnt) > 18 {
			return false
		}
		var p int
		for k := range cnt {
			arr[p] = k
			p++
		}
		sort.Ints(arr[:p])
		for i := 1; i < p; i++ {
			if arr[i]%arr[i-1] != 0 {
				return false
			}
		}
		return true
	}

	cnt := make(map[int]int)

	begin := make([]int, n)
	end := make([]int, n)
	var j int
	for i := 0; i < n; i++ {
		cnt[A[i]]++
		for !check(cnt) {
			if cnt[A[j]] == 1 {
				delete(cnt, A[j])
			} else {
				cnt[A[j]]--
			}
			end[j] = i - 1
			j++
		}
		begin[i] = j
	}
	for j < n {
		end[j] = n - 1
		j++
	}

	pref := make([]int64, n+1)
	pref[0] = 1
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + int64(i-begin[i]+1)
	}

	res := make([]int64, len(Q))
	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		r--
		if l == 0 {
			res[i] = pref[r]
			continue
		}
		from := min(end[l-1], r)
		res[i] = pref[r] - pref[from] + int64(from-l+1)*int64(from-l+2)/2
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(A []int, Q [][]int) []int64 {
	n := len(A)
	arr := make([]int, 20)
	check := func(cnt map[int]int) bool {
		if len(cnt) > 18 {
			return false
		}
		var p int
		for k := range cnt {
			arr[p] = k
			p++
		}
		sort.Ints(arr[:p])
		for i := 1; i < p; i++ {
			if arr[i]%arr[i-1] != 0 {
				return false
			}
		}
		return true
	}

	ps := make([]Pair, len(Q))

	for i := 0; i < len(Q); i++ {
		ps[i] = Pair{Q[i][0], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	cnt := make(map[int]int)

	t := NewTree(n)

	res := make([]int64, len(Q))

	for l, r, i := n-1, n-1, len(Q)-1; l >= 0 && i >= 0; l-- {
		cnt[A[l]]++
		for !check(cnt) {
			if cnt[A[r]] == 1 {
				delete(cnt, A[r])
			} else {
				cnt[A[r]]--
			}
			r--
		}
		t.Add(l, r, 1)
		for i >= 0 && ps[i].first-1 == l {
			j := ps[i].second
			res[j] = t.Get(Q[j][0]-1, Q[j][1]-1)
			i--
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Tree struct {
	arr  []int64
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	t := new(Tree)
	t.sz = n
	t.arr = make([]int64, 4*n)
	t.lazy = make([]int, 4*n)
	return t
}

func (t *Tree) push(i int, l int, r int) {
	if l == r {
		return
	}
	if t.lazy[i] != 0 {
		mid := (l + r) / 2
		t.add(2*i, l, mid, l, r, t.lazy[i])
		t.add(2*i+1, mid+1, r, l, r, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.arr[i] = t.arr[i*2] + t.arr[2*i+1]
}

func (t *Tree) add(i int, l int, r int, L int, R int, v int) {
	if R < l || r < L {
		return
	}
	t.push(i, l, r)
	if L <= l && r <= R {
		t.arr[i] += int64(v) * int64(r-l+1)
		t.lazy[i] += v
		return
	}
	mid := (l + r) / 2
	t.add(2*i, l, mid, L, R, v)
	t.add(2*i+1, mid+1, r, L, R, v)
	t.pull(i)
}

func (t *Tree) Add(L int, R int, v int) {
	t.add(1, 0, t.sz-1, L, R, v)
}

func (t *Tree) Get(L int, R int) int64 {
	var loop func(i int, l int, r int) int64
	loop = func(i int, l int, r int) int64 {
		if R < l || r < L {
			return 0
		}
		t.push(i, l, r)
		if L <= l && r <= R {
			return t.arr[i]
		}
		mid := (l + r) / 2
		a := loop(2*i, l, mid)
		b := loop(2*i+1, mid+1, r)
		return a + b
	}

	return loop(1, 0, t.sz-1)
}
