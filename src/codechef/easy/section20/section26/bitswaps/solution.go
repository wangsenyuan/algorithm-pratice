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

		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		if res {
			buf.WriteString("YES")
		} else {
			buf.WriteString("NO")
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int) bool {
	// only true when A is sorted, or A is connected
	if sort.IntsAreSorted(A) {
		return true
	}

	n := len(A)

	for i := 0; i < n; i++ {
		if A[i] == 0 {
			if i > 0 && A[i-1] > 0 {
				// can't get 0 to front
				return false
			}
		}
	}

	for i := 0; i < n; i++ {
		if A[i] != 0 {
			A = A[i:]
			break
		}
	}

	n = len(A)

	set := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = i
		cnt[i]++
	}

	var find func(x int) int

	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	union := func(x, y int) {
		px := find(x)
		py := find(y)
		if px != py {
			if cnt[px] < cnt[py] {
				px, py = py, px
			}
			cnt[px] += cnt[py]
			set[py] = px
		}
	}

	for i := 0; i < 32; i++ {
		prev := -1
		for j := 0; j < n; j++ {
			num := A[j]
			if (num>>uint(i))&1 == 1 {
				if prev >= 0 {
					union(prev, j)
				}
				prev = j
			}
		}
	}

	groups := make(map[int][]Pair)

	for i := 0; i < n; i++ {
		p := find(i)
		if groups[p] == nil {
			groups[p] = make([]Pair, 0, cnt[p])
		}
		groups[p] = append(groups[p], Pair{A[i], i})
	}
	X := make([]int, n)
	B := make([]int, n)
	for _, v := range groups {
		for i := 0; i < len(v); i++ {
			X[i] = v[i].second
		}
		sort.Sort(Pairs(v))

		for i, p := range v {
			B[X[i]] = p.first
		}
	}

	return sort.IntsAreSorted(B)
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
