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
		A := readString(reader)
		B := readString(reader)
		res := solve(n, A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve1(n int, A string, B string) int {
	a := makeArray(A)
	a = append(a, 3)
	b := makeArray(B)
	b = append(b, 0)
	a = append(a, b...)
	p, lcp := suffixArray(a)

	var best int

	for i := 0; i+1 < len(a); i++ {
		l := p[i]
		r := p[i+1]
		if l < n && r > n || l > n && r < n {
			best = max(best, lcp[i])
		}
	}

	return best
}

func makeArray(s string) []int {
	n := len(s)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1 + int(s[i]-'0')
	}
	return res
}

const H = 16

func suffixArray(arr []int) ([]int, []int) {
	n := len(arr)

	p := make([]int, n)
	c := make([]int, n)
	cnt := make([]int, max(n, 16))

	for i := 0; i < n; i++ {
		cnt[arr[i]]++
	}

	for i := 1; i < 4; i++ {
		cnt[i] += cnt[i-1]
	}

	for i := 0; i < n; i++ {
		cnt[arr[i]]--
		p[cnt[arr[i]]] = i
	}

	c[p[0]] = 0
	classes := 1
	for i := 1; i < n; i++ {
		if arr[p[i]] != arr[p[i-1]] {
			classes++
		}
		c[p[i]] = classes - 1
	}

	pn := make([]int, n)
	cn := make([]int, n)

	for h := 0; (1 << uint(h)) < n; h++ {
		for i := 0; i < n; i++ {
			pn[i] = p[i] - (1 << uint(h))
			if pn[i] < 0 {
				pn[i] += n
			}
		}

		for i := 0; i < classes; i++ {
			cnt[i] = 0
		}
		for i := 0; i < n; i++ {
			cnt[c[pn[i]]]++
		}

		for i := 1; i < classes; i++ {
			cnt[i] += cnt[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			cnt[c[pn[i]]]--
			p[cnt[c[pn[i]]]] = pn[i]
		}

		cn[p[0]] = 0
		classes = 1
		for i := 1; i < n; i++ {
			cur := Pair{c[p[i]], c[(p[i]+(1<<uint(h)))%n]}
			prev := Pair{c[p[i-1]], c[(p[i-1]+(1<<uint(h)))%n]}
			if cur != prev {
				classes++
			}
			cn[p[i]] = classes - 1
		}
		copy(c, cn)
	}

	rank := make([]int, n)

	for i := 0; i < n; i++ {
		rank[p[i]] = i
	}

	var k int
	lcp := make([]int, n-1)
	for i := 0; i < n; i++ {
		if rank[i] == n-1 {
			k = 0
			continue
		}
		j := p[rank[i]+1]
		for i+k < n && j+k < n && arr[i+k] == arr[j+k] {
			k++
		}
		lcp[rank[i]] = k
		if k > 0 {
			k--
		}
	}

	return p, lcp
}

func iota(arr []int, start int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + start
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type State struct {
	len  int
	link int
	next map[byte]int
}

func solve(n int, A string, B string) int {
	st := make([]*State, 2*n+10)
	for i := 0; i < len(st); i++ {
		st[i] = new(State)
		st[i].next = make(map[byte]int)
	}
	st[0].link = -1

	var sz, last int
	sz++

	getNode := func() int {
		res := sz
		sz++
		return res
	}

	extend := func(c byte) {
		cur := getNode()
		st[cur].len = st[last].len + 1

		p := last

		for p != -1 && !contains(st[p].next, c) {
			st[p].next[c] = cur
			p = st[p].link
		}

		if p < 0 {
			st[cur].link = 0
		} else {
			q := st[p].next[c]

			if st[p].len+1 == st[q].len {
				st[cur].link = q
			} else {
				cl := getNode()
				st[cl].len = st[p].len + 1
				st[cl].next = st[q].next
				st[cl].link = st[q].link

				for p != -1 && st[p].next[c] == q {
					st[p].next[c] = cl
					p = st[p].link
				}
				st[cur].link = cl
				st[q].link = cl
			}
		}
		last = cur
	}

	for i := 0; i < n; i++ {
		extend(A[i])
	}

	var v, l, best int

	for i := 0; i < n; i++ {
		for v > 0 && !contains(st[v].next, B[i]) {
			v = st[v].link
			l = st[v].len
		}

		if contains(st[v].next, B[i]) {
			v = st[v].next[B[i]]
			l++
		}
		if l > best {
			best = l
		}
	}

	return best
}

func contains(m map[byte]int, x byte) bool {
	if _, ok := m[x]; ok {
		return ok
	}
	return false
}
