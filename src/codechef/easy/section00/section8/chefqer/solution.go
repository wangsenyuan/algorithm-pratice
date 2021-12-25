package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

	n, k := readTwoNums(reader)

	A := readNNums(reader, n)

	Q := make([][]int, k)

	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t)
		if t == 1 {
			var l, r, x int
			pos = readInt(s, pos+1, &l)
			pos = readInt(s, pos+1, &r)
			readInt(s, pos+1, &x)
			Q[i] = []int{t, l, r, x}
		} else {
			var y int
			readInt(s, pos+1, &y)
			Q[i] = []int{t, y}
		}
	}

	res := solve(n, A, Q)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int, Q [][]int) []int64 {
	// (X + i - L) ** 2
	// let y = X - L
	// y * y + 2 * y * i + i * i
	YY := NewBit(n)
	Y2 := NewBit(n)
	II := NewBit(n)

	res := make([]int64, 0, len(Q))

	for _, cur := range Q {
		if cur[0] == 1 {
			L, R, X := cur[1], cur[2], cur[3]
			L--
			R--

			YY.UpdateRange(L, R, int64(X-L)*int64(X-L))
			Y2.UpdateRange(L, R, int64(2*(X-L)))
			II.UpdateRange(L, R, 1)
		} else {
			P := cur[1]
			P--
			tmp := YY.Query(P)
			tmp += Y2.Query(P) * int64(P)
			tmp += II.Query(P) * int64(P) * int64(P)
			tmp += int64(A[P])
			res = append(res, tmp)
		}
	}
	return res
}

type Bit struct {
	arr []int64
	n   int
}

func NewBit(n int) *Bit {
	arr := make([]int64, n+1)
	return &Bit{arr, n}
}

func (this *Bit) Update(p int, v int64) {
	p++
	for p <= this.n {
		this.arr[p] += v
		p += p & -p
	}
}

func (this *Bit) UpdateRange(l, r int, v int64) {
	this.Update(l, v)
	this.Update(r+1, -v)
}

func (this *Bit) Query(p int) int64 {
	p++
	var res int64
	for p > 0 {
		res += this.arr[p]
		p -= p & -p
	}
	return res
}
