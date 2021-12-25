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
		n := readNum(reader)
		S := make([]string, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadString('\n')
			S[i] = s[:len(s)-1]
		}
		res := solve1(n, S)
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

func solve(n int, S []string) int64 {
	hash := make([][]Key, n)
	var nx int
	for i := 0; i < n; i++ {
		nx = max(nx, len(S[i]))
		hash[i] = getHash(S[i])
	}

	B := make([]Key, nx+1)
	B[0] = Key{1, 1}
	for i := 0; i < nx; i++ {
		B[i+1] = B[i].Mul(B1, B2)
	}
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := getLngComLen(hash[i], hash[j], B)
			edges = append(edges, Edge{d, i, j})
		}
	}

	sort.Sort(Edges(edges))

	var res int64

	set := NewBit(n)

	for _, edge := range edges {
		if set.Union(edge.u, edge.v) {
			res += int64(edge.dist)
		}
	}

	return res
}

type Bit struct {
	arr []int
	cnt []int
}

func NewBit(n int) *Bit {
	b := new(Bit)
	b.arr = make([]int, n)
	b.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		b.arr[i] = i
		b.cnt[i] = 1
	}
	return b
}

func (b *Bit) Find(x int) int {
	if b.arr[x] != x {
		b.arr[x] = b.Find(b.arr[x])
	}
	return b.arr[x]
}

func (b *Bit) Union(x, y int) bool {
	pa := b.Find(x)
	pb := b.Find(y)
	if pa == pb {
		return false
	}
	if b.cnt[pa] < b.cnt[pb] {
		pa, pb = pb, pa
	}
	b.arr[pb] = pa
	b.cnt[pa] += b.cnt[pb]
	return true
}

type Edge struct {
	dist int
	u, v int
}

type Edges []Edge

func (this Edges) Len() int {
	return len(this)
}

func (this Edges) Less(i, j int) bool {
	return this[i].dist > this[j].dist
}

func (this Edges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func getLngComLen(h1, h2 []Key, B []Key) int {
	right := min(len(h1), len(h2)) + 1
	var left int

	for left < right {
		mid := (left + right) / 2
		if !check(h1, h2, B, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}

func check(h1, h2, B []Key, l int) bool {
	if l == 0 {
		return true
	}

	if l >= len(h1) || l >= len(h2) {
		return false
	}
	mem := make(map[Key]bool)
	// whether some l length sub string exists in both 1 & 2
	for i := 0; i+l < len(h1); i++ {
		a := h1[i]
		b := h1[i+l]
		// a * B[l] + x = b
		// x = b - a * B[l]
		x := b.Sub(a.Mul2(B[l]))
		mem[x] = true
	}

	for i := 0; i+l < len(h2); i++ {
		a := h2[i]
		b := h2[i+l]
		x := b.Sub(a.Mul2(B[l]))
		if mem[x] {
			return true
		}
	}
	return false
}

func getHash(s string) []Key {
	n := len(s)
	k := make([]Key, n+1)
	k[0] = Key{0, 0}

	for i := 0; i < len(s); i++ {
		k[i+1] = k[i].Add(int(s[i]-'a') + 1)
	}
	return k
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

const M = 1000000007

const B1 = 29
const B2 = 31

type Key struct {
	first  int64
	second int64
}

func (this Key) Add(v int) Key {
	first := (this.first*B1%M + int64(v)) % M
	second := (this.second*B2%M + int64(v)) % M
	return Key{first, second}
}

func (this Key) Mul(a, b int64) Key {
	first := this.first * a % M
	second := this.second * b % M
	return Key{first, second}
}

func (this Key) Mul2(that Key) Key {
	first := this.first * that.first % M
	second := this.second * that.second % M
	return Key{first, second}
}

func (this Key) Sub(that Key) Key {
	first := (this.first + M - that.first) % M
	second := (this.second + M - that.second) % M
	return Key{first, second}
}

const alphabet = 28
const first_ch = '`'
const last_ch = '{'

func solve1(n int, S []string) int64 {

	sortCyclic := func(s string) []int {
		n := len(s)
		//p[i] = i-th sub string starting index
		p := make([]int, n)
		//c[i] = i-th position rank (classs)
		c := make([]int, n)
		cnt := make([]int, max(alphabet, n)+1)

		for i := 0; i < n; i++ {
			cnt[int(s[i]-first_ch)]++
		}

		for i := 1; i < alphabet; i++ {
			cnt[i] += cnt[i-1]
		}

		for i := 0; i < n; i++ {
			cnt[int(s[i]-first_ch)]--
			x := cnt[int(s[i]-first_ch)]
			p[x] = i
		}

		c[p[0]] = 0

		rank := 1
		for i := 1; i < n; i++ {
			if s[p[i]] != s[p[i-1]] {
				rank++
			}
			c[p[i]] = rank - 1
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
			for i := 0; i < len(cnt); i++ {
				cnt[i] = 0
			}

			for i := 0; i < n; i++ {
				cnt[c[pn[i]]]++
			}

			for i := 1; i < len(cnt); i++ {
				cnt[i] += cnt[i-1]
			}
			for i := n - 1; i >= 0; i-- {
				cnt[c[pn[i]]]--
				p[cnt[c[pn[i]]]] = pn[i]
			}

			cn[pn[0]] = 0
			rank := 1
			for i := 1; i < n; i++ {
				if c[p[i]] != c[p[i-1]] || c[(p[i]+(1<<uint(h)))%n] != c[(p[i-1]+(1<<uint(h)))%n] {
					rank++
				}
				cn[p[i]] = rank - 1
			}

			c, cn = cn, c
		}
		return p
	}

	suffixArray := func(s string) []int {
		buf := make([]byte, len(s)+1)
		copy(buf, s)
		buf[len(s)] = first_ch
		ret := sortCyclic(string(buf))
		return ret[1:]
	}

	lcp := func(s string, p []int) []int {
		n := len(s)
		rank := make([]int, n)
		for i := 0; i < n; i++ {
			rank[p[i]] = i
		}

		var k int
		res := make([]int, n-1)

		for i := 0; i < n; i++ {
			if rank[i] == n-1 {
				k = 0
				continue
			}
			j := p[rank[i]+1]
			for i+k < n && j+k < n && s[i+k] == s[j+k] {
				k++
			}
			res[rank[i]] = k
			if k > 0 {
				k--
			}
		}
		return res
	}
	var which []int

	var buf bytes.Buffer

	pref := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			pref[i] += pref[i-1]
		}
		pref[i]++
		pref[i] += len(S[i])
		for j := 0; j < len(S[i]); j++ {
			buf.WriteByte(S[i][j])
			which = append(which, i)
		}
		if i+1 < n {
			buf.WriteByte(last_ch)
			which = append(which, -1)
		}
	}
	combined := buf.String()
	sa := suffixArray(combined)
	lp := lcp(combined, sa)

	edges := make([]Edge, 0, n*(n-1)/2)

	for i := 0; i < len(lp)-1; i++ {
		a := which[sa[i]]
		b := which[sa[i+1]]
		if a == b || a < 0 || b < 0 {
			continue
		}
		wt := min(lp[i], pref[a]-1-sa[i])
		edges = append(edges, Edge{wt, a, b})
	}
	sort.Sort(Edges(edges))

	var res int64

	set := NewBit(n)

	for _, edge := range edges {
		if set.Union(edge.u, edge.v) {
			res += int64(edge.dist)
		}
	}

	return res
}
