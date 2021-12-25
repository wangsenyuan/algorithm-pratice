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
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][0]))
		}
		buf.WriteByte('\n')
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][1]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

const INF = 1 << 30

func solve(A []int) [][]int {
	N := len(A)
	var n int

	for i := 0; i < len(A); i++ {
		if n > 0 && A[i] == A[n-1] {
			n--
		} else if n > 1 {
			if A[n-2] > A[n-1] && A[n-1] > A[i] {
				n--
			} else if A[n-2] < A[n-1] && A[n-1] < A[i] {
				n--
			}
		}
		A[n] = A[i]
		n++
	}

	var mp1 Pairs
	var mp2 Pairs

	add := func(mp Pairs, a, b int) Pairs {
		return append(mp, &Pair{a, b})
	}

	for i := 0; i < n; i++ {
		mp1 = add(mp1, A[i]+1, 0)
		mp1 = add(mp1, A[i]-1, 0)
		mp2 = add(mp2, A[i]+1, 0)
		mp2 = add(mp2, A[i]-1, 0)
	}

	sort.Sort(mp1)
	mp1 = unique(mp1)
	sort.Sort(mp2)
	mp2 = unique(mp2)

	update := func(mp Pairs, a, b int) {
		L := search(len(mp), func(i int) bool {
			return mp[i].first >= a
		})
		R := search(len(mp), func(i int) bool {
			return mp[i].first >= b
		})
		mp[L].second++
		if R+1 < len(mp) {
			mp[R+1].second--
		}
	}

	for i := 0; i+1 < n; i++ {
		if A[i] < A[i+1] {
			update(mp1, A[i]+1, A[i+1]-1)
		} else {
			update(mp2, A[i+1]+1, A[i]-1)
		}
	}

	for i := 1; i < len(mp1); i++ {
		mp1[i].second += mp1[i-1].second
		mp2[i].second += mp2[i-1].second
	}
	mm := make([]int, N+1)
	MM := make([]int, N+1)
	for i := 0; i <= N; i++ {
		mm[i] = INF
		MM[i] = -INF
	}

	for i := 0; i < len(mp1); i++ {
		f := mp1[i].second + mp2[i].second + 1
		mm[f] = min(mm[f], mp1[i].first)
		MM[f] = max(MM[f], mp1[i].first)
	}
	for i := N - 1; i >= 0; i-- {
		mm[i] = min(mm[i], mm[i+1])
		MM[i] = max(MM[i], MM[i+1])
	}
	res := make([][]int, N-1)
	for i := 2; i <= N; i++ {
		res[i-2] = []int{mm[i], MM[i]}
		if res[i-2][0] == INF {
			res[i-2][0] = -1
		}
		if res[i-2][1] == -INF {
			res[i-2][1] = -1
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct{ first, second int }

type Pairs []*Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func unique(ps Pairs) Pairs {
	var n int
	for i := 1; i <= len(ps); i++ {
		if i == len(ps) || ps[i].first > ps[i-1].first {
			ps[n] = ps[i-1]
			n++
		}
	}
	return ps[:n]
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
