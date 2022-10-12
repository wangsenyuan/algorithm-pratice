package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, q := readTwoNums(reader)
	A := readNInt64s(reader, n)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(A, Q)

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func solve(A []int64, Q [][]int) []int {
	var mx int64
	var mn int64 = math.MaxInt64

	for _, num := range A {
		mn = min(mn, num)
		mx = max(mx, num)
	}
	flag := make([]bool, 1000001)
	sp := make([]int, 10000001)
	for i := int64(2); i*i <= mx; i++ {
		if flag[int(i)] {
			continue
		}
		for j := i * i; j*j <= mx; j += i {
			flag[int(j)] = true
		}
		st := (mn+i-1)/i*i - mn
		for j := st; j <= mx-mn; j += i {
			if sp[int(j)] == 0 {
				sp[int(j)] = int(i)
			}
		}
	}

	var ptr int

	n := len(A)

	mp := make(map[int64]int)

	B := make([]int, n+1)

	inv := make([]int, n+2)

	for i := 0; i < n; i++ {
		if sp[int(A[i]-mn)] != 0 {
			A[i] = int64(sp[int(A[i]-mn)])
		}
		if mp[A[i]] == 0 {
			ptr++
			mp[A[i]] = ptr
		}

		B[i+1] = mp[A[i]]
		inv[i+1] = inverse(i + 1)
	}

	inv[n+1] = inverse(n + 1)

	qry := make([]Query, len(Q))

	for i, cur := range Q {
		qry[i] = Query{cur[0], cur[1], i}
	}

	sort.Slice(qry, func(i, j int) bool {
		return qry[i].Less(qry[j])
	})

	ans := make([]int, len(qry))

	var l, r int = 1, 0
	var res int = 1

	um := make([]int, ptr+1)

	add := func(i int) {
		res = modMul(res, inv[um[B[i]]+1])
		um[B[i]]++
		res = modMul(res, um[B[i]]+1)
	}

	del := func(i int) {
		res = modMul(res, inv[um[B[i]]+1])
		um[B[i]]--
		res = modMul(res, um[B[i]]+1)
	}

	for _, cur := range qry {
		tl := cur.l
		tr := cur.r

		for r < tr {
			r++
			add(r)
		}
		for l > tl {
			l--
			add(l)
		}

		for l < tl {
			del(l)
			l++
		}

		for r > tr {
			del(r)
			r--
		}
		ans[cur.id] = modSub(res, 1)
	}

	return ans
}

const BLOCK_SIZE = 300

type Query struct {
	l  int
	r  int
	id int
}

func (this Query) Less(that Query) bool {
	if this.l/BLOCK_SIZE != that.l/BLOCK_SIZE {
		return this.l/BLOCK_SIZE < that.l/BLOCK_SIZE
	}
	return this.r < that.r
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a int, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}
