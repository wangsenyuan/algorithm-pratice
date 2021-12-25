package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		C, N, M := readThreeNums(reader)
		S := make([][]int, C)
		for i := 0; i < C; i++ {
			x := readNum(reader)
			S[i] = readNNums(reader, 2*x)
		}
		res := solve(C, N, M, S)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const MOD = 998244353

func solve(C, N, M int, S [][]int) int {
	var evtsCnt int
	for _, s := range S {
		evtsCnt += len(s) / 2
	}

	evts := make([]Event, 0, evtsCnt)

	for i, s := range S {
		for j := 0; j < len(s); j += 2 {
			cur := Event{s[j], s[j+1], i}
			evts = append(evts, cur)
		}
	}

	sort.Sort(Events(evts))

	uf := NewUF(2 * C)
	id := C
	free := N
	for i := 0; i < len(evts); {
		cur := evts[i]
		r := cur.right
		for i < len(evts) && evts[i].left <= r {
			uf.Union(evts[i].club, id)
			r = max(r, evts[i].right)
			i++
		}
		free -= r - cur.left + 1
		// i == len(evts) or evts[i].left > r
		id++
	}
	roots := make(map[int]bool)

	for i := 0; i < C; i++ {
		roots[uf.Find(i)] = true
	}

	return pow(M, len(roots)+free)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

type UF struct {
	arr  []int
	cnt  []int
	size int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}

	return &UF{arr, cnt, n}
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
	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.arr[pb] = pa
	uf.cnt[pa] += uf.cnt[pb]
	uf.size--
	return true
}

type Event struct {
	left, right int
	club        int
}

type Events []Event

func (evts Events) Len() int {
	return len(evts)
}

func (evts Events) Less(i, j int) bool {
	return evts[i].left < evts[j].left
}

func (evts Events) Swap(i, j int) {
	evts[i], evts[j] = evts[j], evts[i]
}
