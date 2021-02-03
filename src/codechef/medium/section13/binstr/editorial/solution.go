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

	n, q := readTwoNums(reader)
	A := make([]string, n)
	for i := 0; i < n; i++ {
		A[i], _ = reader.ReadString('\n')
		A[i] = normalize(A[i])
	}
	L, R, X := make([]int, q), make([]int, q), make([]string, q)
	for i := 0; i < q; i++ {
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &L[i]) + 1
		pos = readInt(s, pos, &R[i]) + 1
		X[i] = normalize(string(s[pos:]))
	}
	res := solve(n, q, A, L, R, X)
	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {
	buf := []byte(s)
	var n int
	for n < len(buf) {
		if buf[n] != '0' && buf[n] != '1' {
			break
		}
		n++
	}
	return string(buf[:n])
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

const INF = 1000000010

var root [100005]int
var left [6000006]int
var right [6000006]int
var t [6000006]int
var tt [400004]int
var maxPref [100005]int
var pos [100005]int

func solve(N, Q int, A []string, L []int, R []int, X []string) []int {
	bb := make(BB, N)
	var maxLen int
	for i := 0; i < N; i++ {
		bb[i] = B{A[i], i + 1}
		if len(A[i]) > maxLen {
			maxLen = len(A[i])
		}
	}
	sort.Sort(bb)

	for i := 0; i < N; i++ {
		pos[bb[i].i] = i + 1
	}

	for i := 0; i < Q; i++ {
		if len(X[i]) > maxLen {
			maxLen = len(X[i])
		}
	}

	bts := make([][]int, maxLen)

	for i := 0; i < maxLen; i++ {
		bts[i] = make([]int, 0, 3)
	}

	addNumber := func(s string, x int) {
		buf := []byte(s)
		reverse(buf)
		for i := 0; i < len(buf); i++ {
			if buf[i] == '1' {
				bts[i] = append(bts[i], x)
			}
		}
	}

	for i := 0; i < N; i++ {
		addNumber(bb[i].s, i+1)
	}

	sz := 1
	for sz < N {
		sz += sz
	}
	// sz >= N

	for i := 1; i < sz; i++ {
		left[i] = i << 1
		right[i] = left[i] | 1
	}
	all := sz << 1
	for i := 1; i <= all; i++ {
		t[i] = INF
	}

	vcopy := func(cur int) int {
		all++
		left[all] = left[cur]
		right[all] = right[cur]
		t[all] = t[cur]
		return all
	}

	var update func(l, r, x, y, z int)

	update = func(l, r, x, y, z int) {
		if l == r {
			t[y] = z
			return
		}
		mid := (l + r) >> 1
		if x <= mid {
			left[y] = vcopy(left[y])
			update(l, mid, x, left[y], z)
		} else {
			right[y] = vcopy(right[y])
			update(mid+1, r, x, right[y], z)
		}
		t[y] = min(t[left[y]], t[right[y]])
	}

	root[N+1] = 1

	for i := N; i > 0; i-- {
		root[i] = vcopy(root[i+1])
		update(1, sz, pos[i], root[i], i)
	}

	for i := 1; i < N; i++ {
		if len(bb[i-1].s) != len(bb[i].s) {
			continue
		}
		var z int
		for z < len(bb[i-1].s) && bb[i-1].s[z] == bb[i].s[z] {
			z++
		}
		maxPref[i] = z
	}

	for i := 0; i < len(tt); i++ {
		tt[i] = INF
	}

	for i := 1; i <= N; i++ {
		tt[i+sz-1] = maxPref[i-1]
	}

	for i := sz - 1; i > 0; i-- {
		tt[i] = min(tt[i<<1], tt[(i<<1)|1])
	}

	var get func(l, r, ll, rr, x int) int
	get = func(l, r, ll, rr, x int) int {
		if l > r || ll > r || l > rr || ll > rr {
			return INF
		}
		if l >= ll && r <= rr {
			return t[x]
		}
		mid := (l + r) >> 1
		return min(get(l, mid, ll, rr, left[x]), get(mid+1, r, ll, rr, right[x]))
	}

	check := func(l, r, ll, rr int) bool {
		return get(1, sz, l, r, root[ll]) <= rr
	}

	getPos := func(l, r, x int) int {
		if len(bts[x]) == 0 {
			return r + 1
		}
		var ll int
		rr := len(bts[x]) - 1
		for ll < rr-1 {
			mid := (ll + rr) / 2
			if bts[x][mid] >= l {
				rr = mid
			} else {
				ll = mid
			}
		}
		if bts[x][ll] >= l {
			rr = ll
		}
		if bts[x][rr] < l {
			return r + 1
		}
		return bts[x][rr]
	}

	var find func(l, r, ll, rr int, x []byte, y int) int

	find = func(l, r, ll, rr int, x []byte, y int) int {
		if y < 0 {
			return get(1, sz, l, r, root[ll])
		}

		z := getPos(l, r, y)

		z = min(z, r+1)

		if x[y] == '1' {
			if z > l && check(l, z-1, ll, rr) {
				return find(l, z-1, ll, rr, x, y-1)
			}
			return find(z, r, ll, rr, x, y-1)
		}
		if z <= r && check(z, r, ll, rr) {
			return find(z, r, ll, rr, x, y-1)
		}
		return find(l, z-1, ll, rr, x, y-1)
	}

	var getLastPos func(l, r, x, rr int) int

	getLastPos = func(l, r, x, rr int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if t[right[x]] <= rr {
			return getLastPos(mid+1, r, right[x], rr)
		}
		return getLastPos(l, mid, left[x], rr)
	}

	var getSecondPos func(l, r, ll, rr, x, y int) int

	getSecondPos = func(l, r, ll, rr, x, y int) int {
		if l > r || ll > r || l > rr || ll > rr {
			return 0
		}
		if l >= ll && r <= rr && tt[x] >= y {
			return 0
		}
		if l == r {
			return l
		}
		mid := (l + r) >> 1
		z := getSecondPos(mid+1, r, ll, rr, x+x+1, y)
		if z > 0 {
			return z
		}
		return getSecondPos(l, mid, ll, rr, x+x, y)
	}
	ans := make([]int, Q)

	for i := 0; i < Q; i++ {
		ll, rr, S := L[i], R[i], X[i]
		buf := []byte(S)
		reverse(buf)
		l, r := 1, N
		for l < r-1 {
			mid := (l + r) / 2
			if len(bb[mid-1].s) > len(buf) {
				r = mid
			} else {
				l = mid
			}
		}

		if len(bb[l-1].s) > len(buf) {
			r = l
		}
		if len(bb[r-1].s) <= len(buf) || !check(r, N, ll, rr) {
			if len(bb[r-1].s) > len(buf) {
				r--
			}
			ans[i] = find(1, r, ll, rr, buf, len(buf)-1)
		} else {
			z := getLastPos(1, sz, root[ll], rr)
			zz := getSecondPos(1, sz, 1, z, 1, len(bb[z].s)-len(buf))
			ans[i] = find(zz, z, ll, rr, buf, len(buf)-1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type B struct {
	s string
	i int
}

type BB []B

func (bb BB) Len() int {
	return len(bb)
}

func (bb BB) Less(i, j int) bool {
	if len(bb[i].s) < len(bb[j].s) {
		return true
	}
	if len(bb[i].s) == len(bb[j].s) && bb[i].s < bb[j].s {
		return true
	}
	return false
}

func (bb BB) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}
