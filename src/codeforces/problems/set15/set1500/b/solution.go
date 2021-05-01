package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadBytes('\n')

	var n, m int
	pos := readInt(line, 0, &n) + 1
	pos = readInt(line, pos, &m) + 1
	var k uint64
	readUint64(line, pos, &k)

	A := readNNums(reader, n)
	B := readNNums(reader, m)
	res := solve(A, B, int64(k))

	fmt.Println(res)
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

func extGcd(a, b int, x *int64, y *int64) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var x1, y1 int64
	g := extGcd(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - int64(a/b)*y1

	return g
}

const MAX = 1000010

func solve(A, B []int, K int64) int64 {
	n := len(A)
	m := len(B)

	pa := make([]int, MAX)

	for i := 0; i < MAX; i++ {
		pa[i] = -1
	}

	for i, a := range A {
		pa[a] = i
	}
	var x, y int64

	g := extGcd(n, m, &x, &y)

	M := int64(m / g)

	all := make([]int64, 0, m)

	for i := 0; i < m; i++ {
		if pa[B[i]] < 0 || (i-pa[B[i]])%g != 0 {
			continue
		}
		kk := int64(i-pa[B[i]]) / int64(g)
		xx := x * kk
		xx = (xx%M + M) % M
		xx = xx*int64(n) + int64(pa[B[i]]) + 1
		all = append(all, xx)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})

	L := int64(n) * M
	aa := L - int64(len(all))
	day := K / aa * L
	K %= aa
	if K == 0 {
		day -= L
		K = aa
	}

	check := func(x int64) bool {
		i := sort.Search(len(all), func(i int) bool {
			return all[i] > x
		})
		return x-int64(i) >= K
	}

	var l, r int64 = 0, L

	for l < r {
		mid := (l + r) / 2

		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r + day
}
