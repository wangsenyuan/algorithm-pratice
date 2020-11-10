package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, u := readTwoNums(reader)

	inst := make([][]int, u)
	for i := 0; i < u; i++ {
		inst[i] = readNNums(reader, 2)
	}

	fmt.Println(solve(n, u, inst))
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

//const MAX_N = 100010

var zero = NewHash(0)

func solve(n, u int, inst [][]int) string {
	MAX_N := n + 1
	precalc := make([]Hash, MAX_N)
	precalc[0] = zero
	precalc[1] = NewHash(1)

	base := NewHash(1337)

	for i := 2; i < MAX_N; i++ {
		precalc[i] = precalc[i-1].Mul(base).AddInt(1)
	}
	N := 15 * 18 * MAX_N
	L := make([]int, N)
	R := make([]int, N)
	P := make([]int, N)
	subs := make([]Hash, N)
	root := make([]int, u+1)
	var sz = 1
	var rt = 1

	copy := func(u *int, v int) int {
		L[sz] = L[v]
		R[sz] = R[v]
		P[sz] = P[v]
		subs[sz] = subs[v]
		*u = sz
		sz++
		return *u
	}

	makeRoot := func() int {
		copy(&root[rt], root[rt-1])
		rt++
		return root[rt-1]
	}

	segment := func(l, r int) Hash {
		return precalc[r].Sub(precalc[l])
	}

	push := func(v int, l, r int) {
		if P[v] > 0 {
			subs[v] = segment(l, r).Sub(subs[v])
			if r-l > 1 {
				P[copy(&L[v], L[v])] ^= 1
				P[copy(&R[v], R[v])] ^= 1
			}
		}
		P[v] = 0
	}

	var inv func(a, b int, v int, l, r int)

	inv = func(a, b int, v int, l, r int) {
		push(v, l, r)
		if a <= l && r <= b {
			P[v] = 1
			push(v, l, r)
			return
		}
		if r <= a || b <= l {
			return
		}
		mid := (l + r) / 2
		inv(a, b, copy(&L[v], L[v]), l, mid)
		inv(a, b, copy(&R[v], R[v]), mid, r)
		subs[v] = subs[L[v]].Add(subs[R[v]])
	}

	var get func(p, v int, l, r int) Hash

	get = func(p, v int, l, r int) Hash {
		push(v, l, r)
		if r-l == 1 {
			return NewHash(0)
		}
		mid := (l + r) / 2
		if p < mid {
			return get(p, L[v], l, mid)
		}
		push(L[v], l, mid)
		return subs[L[v]].Add(get(p, R[v], mid, r))
	}

	small := func(a, b int) bool {
		a, b = root[a], root[b]
		var l, r = 0, n
		for r-l > 1 {
			mid := (l + r) / 2
			if get(mid, a, 0, MAX_N) == get(mid, b, 0, MAX_N) {
				l = mid
			} else {
				r = mid
			}
		}
		x := get(l+1, a, 0, MAX_N) == get(l, a, 0, MAX_N)
		y := get(l+1, b, 0, MAX_N) != get(l, b, 0, MAX_N)
		return x && y
	}

	var cur int
	for i := 1; i <= u; i++ {
		l, r := inst[i-1][0], inst[i-1][1]
		root[i] = makeRoot()
		inv(l-1, r, root[i], 0, MAX_N)
		if small(cur, i) {
			cur = i
		}
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if get(i+1, root[cur], 0, MAX_N) != get(i, root[cur], 0, MAX_N) {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}
	return buf.String()
}

var MOD = [...]uint64{1000000007, 1000000009}

type Hash struct {
	h [2]uint64
}

func NewHash(x uint64) Hash {
	h := [2]uint64{x % MOD[0], x % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x uint64) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + x%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint64{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}
