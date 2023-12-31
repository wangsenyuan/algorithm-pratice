package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)

	Q := make([][]int, k)
	for i := 0; i < k; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(a, Q)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const N = 5000010

var phi [N]int
var D [N]int
var P [N][6]int

type Node struct {
	lca, ans, mndis int
	allrt           bool
}

var tr [4 * N]*Node

func init() {
	phi[1] = 1
	var primes []int

	for i := 2; i < N; i++ {
		if phi[i] == 0 {
			phi[i] = i - 1
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= N {
				break
			}
			if i%primes[j] == 0 {
				phi[i*primes[j]] = phi[i] * primes[j]
				break
			} else {
				phi[i*primes[j]] = phi[i] * phi[primes[j]]
			}
		}
	}
	D[1] = 1
	for i := 0; i < 6; i++ {
		P[1][i] = 1
	}
	for i := 2; i < N; i++ {
		D[i] = D[phi[i]] + 1
		P[i][0] = phi[i]
		for j := 1; j < 6; j++ {
			P[i][j] = P[P[i][j-1]][j-1]
		}
	}

	for i := 1; i < len(tr); i++ {
		tr[i] = new(Node)
	}
}

func lca(u int, v int) int {
	if D[u] < D[v] {
		u, v = v, u
	}
	for i := 5; i >= 0; i-- {
		if D[u]-(1<<i) >= D[v] {
			u = P[u][i]
		}
	}
	if u == v {
		return u
	}
	for i := 5; i >= 0; i-- {
		if P[u][i] != P[v][i] {
			u = P[u][i]
			v = P[v][i]
		}
	}
	return P[u][0]
}

func pullUp(i int, l int, r int) {
	mid := (l + r) / 2
	tr[i].lca = lca(tr[2*i].lca, tr[2*i+1].lca)
	tr[i].ans = tr[2*i].ans + tr[2*i+1].ans +
		(mid-l+1)*(D[tr[2*i].lca]-D[tr[i].lca]) +
		(r-mid)*(D[tr[2*i+1].lca]-D[tr[i].lca])
	tr[i].mndis = min(tr[2*i].mndis, tr[2*i+1].mndis)
	tr[i].allrt = tr[2*i].allrt && tr[2*i+1].allrt
}

func build(i int, l int, r int, arr []int) {
	if l == r {
		tr[i].lca = arr[l]
		tr[i].ans = 0
		tr[i].mndis = D[arr[l]]
		tr[i].allrt = arr[l] == 1
		return
	}
	mid := (l + r) / 2
	build(i*2, l, mid, arr)
	build(i*2+1, mid+1, r, arr)
	pullUp(i, l, r)
}

func modify(u int, l int, r int, ql int, qr int) {
	if tr[u].allrt {
		// 全部都是1
		return
	}
	if l == r {
		tr[u].lca = P[tr[u].lca][0]
		tr[u].mndis--
		tr[u].allrt = tr[u].lca == 1
		return
	}
	mid := (l + r) / 2
	if ql <= mid {
		modify(2*u, l, mid, ql, qr)
	}
	if mid < qr {
		modify(2*u+1, mid+1, r, ql, qr)
	}
	pullUp(u, l, r)
}

func queryLca(u int, l int, r int, ql int, qr int) int {
	if ql <= l && r <= qr {
		return tr[u].lca
	}
	mid := (l + r) / 2
	if qr <= mid {
		return queryLca(u*2, l, mid, ql, qr)
	}
	if ql > mid {
		return queryLca(u*2+1, mid+1, r, ql, qr)
	}
	ans := queryLca(u*2, l, mid, ql, qr)
	if ans == 1 {
		return 1
	}
	return lca(ans, queryLca(u*2+1, mid+1, r, ql, qr))
}

func queryAns(u int, l int, r int, ql int, qr int, a int) int {
	if ql <= l && r <= qr {
		return tr[u].ans + (r-l+1)*(D[tr[u].lca]-D[a])
	}
	mid := (l + r) / 2
	var res int
	if ql <= mid {
		res += queryAns(u*2, l, mid, ql, qr, a)
	}
	if mid < qr {
		res += queryAns(u*2+1, mid+1, r, ql, qr, a)
	}
	return res
}

func solve(a []int, Q [][]int) []int {
	// [l...r] change a[i] => phi[a[i]]
	// 这个可以用 (lazy) range update to process
	// a[i] => phi[a[i]] 只会越来越小
	// 如果 a[i] != a[j] 那么 phi[a[i]] != phi[a[j]]， 这个成立吗?
	// 这个是不成立的
	n := len(a)

	w := a[0]
	for i := 0; i < n; i++ {
		w = max(w, a[i])
	}
	w++
	
	build(1, 0, n-1, a)
	var res []int

	for _, cur := range Q {
		l, r := cur[1], cur[2]
		l--
		r--
		if cur[0] == 1 {
			modify(1, 0, n-1, l, r)
		} else {
			x := queryLca(1, 0, n-1, l, r)
			y := queryAns(1, 0, n-1, l, r, x)
			res = append(res, y)
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
