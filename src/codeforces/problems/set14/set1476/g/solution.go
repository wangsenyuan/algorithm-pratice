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
	n, m := readTwoNums(reader)
	A := readNNums(reader, n)

	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			Q[i] = readNNums2(s, 4)
		} else {
			Q[i] = readNNums2(s, 3)
		}
	}

	res := solve(A, Q)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
}

func readNNums2(buf []byte, n int) []int {
	res := make([]int, n)
	var pos int
	for i := 0; i < n; i++ {
		pos = readInt(buf, pos, &res[i]) + 1
	}
	return res
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

const BLK_SIZE = 2000
const N = 100013

func solve(A []int, Q [][]int) []int {
	q := make([]Query, 0, len(Q))
	upd := make([]Update, 0, len(Q))

	for i := 0; i < len(Q); i++ {
		if Q[i][0] == 1 {
			l, r, k := Q[i][1], Q[i][2], Q[i][3]
			q = append(q, Query{len(upd), l - 1, r - 1, k, len(q)})
		} else {
			p, x := Q[i][1], Q[i][2]
			p--
			upd = append(upd, Update{p, A[p], x})
			A[p] = x
		}
	}

	sort.Slice(q, func(i, j int) bool {
		return q[i].Less(q[j])
	})

	for i := len(upd) - 1; i >= 0; i-- {
		A[upd[i][0]] = upd[i][1]
	}

	cnt := make([]int, N)
	ord := make([]int, N)
	bounds := make([]*Pair, N)
	for i := 0; i < N; i++ {
		bounds[i] = &Pair{N, 0}
	}
	bounds[0] = &Pair{0, N - 1}

	add := func(x int) {
		c := cnt[x]
		ord[bounds[c].first]++
		bounds[c+1].second = bounds[c].first
		if bounds[c+1].first == N {
			bounds[c+1].first = bounds[c].first
		}
		if bounds[c].first == bounds[c].second {
			bounds[c].first = N - 1
		}
		bounds[c].first++
		cnt[x]++
	}

	rem := func(x int) {
		c := cnt[x]
		ord[bounds[c].second]--

		if bounds[c-1].first == N {
			bounds[c-1].second = bounds[c].second
		}
		bounds[c-1].first = bounds[c].second
		if bounds[c].first == bounds[c].second {
			bounds[c].first = N
		}
		bounds[c].second--
		cnt[x]--
	}

	var L, T int
	R := -1

	apply := func(i, fl int) {
		p := upd[i][0]
		x := upd[i][fl+1]
		if L <= p && p <= R {
			rem(A[p])
			add(x)
		}
		A[p] = x
	}

	res := make([]int, len(q))

	for _, cur := range q {
		t, l, r, k := cur.t, cur.l, cur.r, cur.k
		for T < t {
			apply(T, 1)
			T++
		}
		for T > t {
			T--
			apply(T, 0)
		}
		for R < r {
			R++
			add(A[R])
		}
		for L > l {
			L--
			add(A[L])
		}
		for R > r {
			rem(A[R])
			R--
		}
		for L < l {
			rem(A[L])
			L++
		}

		ans := N
		for i, j, sum := 0, 0, 0; i < N && ord[i] > 0; i = bounds[ord[i]].second + 1 {
			for j < N && ord[j] > 0 && sum < k {
				sum += bounds[ord[j]].second - bounds[ord[j]].first + 1
				j = bounds[ord[j]].second + 1
			}
			if sum >= k {
				ans = min(ans, ord[i]-ord[j-1])
			}
			sum -= bounds[ord[i]].second - bounds[ord[i]].first + 1
		}
		if ans == N {
			res[cur.i] = -1
		} else {
			res[cur.i] = ans
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

type Query struct {
	t, l, r, k, i int
}

func (a Query) Less(b Query) bool {
	if a.t/BLK_SIZE != b.t/BLK_SIZE {
		return a.t < b.t
	}

	if a.l/BLK_SIZE != b.l/BLK_SIZE {
		return a.l < b.l
	}

	if (a.l/BLK_SIZE)&1 == 1 {
		return a.r < b.r
	}
	return a.r > b.r
}

type Pair struct {
	first, second int
}

type Update [3]int
