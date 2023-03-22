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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, k := readTwoNums(reader)
		mines := make([][]int, n)
		for i := 0; i < n; i++ {
			mines[i] = readNNums(reader, 3)
		}
		res := solve(mines, k)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1e9 + 10

func solve(mines [][]int, k int) int {
	n := len(mines)
	ms := make([]Mine, n)

	for i := 0; i < n; i++ {
		ms[i] = Mine{mines[i][0], mines[i][1], mines[i][2]}
	}

	xs := make(map[int][]Pair)
	ys := make(map[int][]Pair)

	for i, m := range ms {
		if len(xs[m.x]) == 0 {
			xs[m.x] = make([]Pair, 0, 1)
		}
		xs[m.x] = append(xs[m.x], Pair{m.y, i})
		if len(ys[m.y]) == 0 {
			ys[m.y] = make([]Pair, 0, 1)
		}
		ys[m.y] = append(ys[m.y], Pair{m.x, i})
	}

	uf := NewUF(n)

	for _, v := range xs {
		sort.Slice(v, func(i, j int) bool {
			return v[i].first < v[j].first
		})

		// 同一x
		for i := 1; i < len(v); i++ {
			if v[i].first-v[i-1].first <= k {
				uf.Union(v[i].second, v[i-1].second)
			}
		}
	}

	for _, v := range ys {
		sort.Slice(v, func(i, j int) bool {
			return v[i].first < v[j].first
		})
		for i := 1; i < len(v); i++ {
			if v[i].first-v[i-1].first <= k {
				uf.Union(v[i].second, v[i-1].second)
			}
		}
	}

	time := make(map[int]int)
	for i := 0; i < n; i++ {
		j := uf.Find(i)
		if v, ok := time[j]; !ok || v > ms[i].t {
			time[j] = ms[i].t
		}
	}
	// 检查分组，每个分组的起爆时间是其中的最小值
	var arr []int
	for _, v := range time {
		arr = append(arr, v)
	}

	sort.Ints(arr)
	l := len(arr)

	best := min(arr[l-1], l-1)

	for i := l - 2; i >= 0; i-- {
		best = min(best, max(arr[i], l-2-i))
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	uf := new(UF)
	uf.arr = make([]int, n)
	uf.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		uf.arr[i] = i
		uf.cnt[i] = 1
	}
	return uf
}

func (uf *UF) Find(u int) int {
	if uf.arr[u] != u {
		uf.arr[u] = uf.Find(uf.arr[u])
	}
	return uf.arr[u]
}

func (uf *UF) Union(a, b int) bool {
	a = uf.Find(a)
	b = uf.Find(b)
	if a == b {
		return false
	}
	if uf.cnt[a] < uf.cnt[b] {
		a, b = b, a
	}
	uf.arr[b] = a
	uf.cnt[a] += uf.cnt[b]
	return true
}

type Pair struct {
	first  int
	second int
}

type Mine struct {
	x int
	y int
	t int
}
