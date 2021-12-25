package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, b := readTwoNums(reader)

	N := make([]int, n)
	for i := 0; i < n; i++ {
		N[i] = readNum(reader)
	}

	q := readNum(reader)

	Q := make([]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNum(reader)
	}

	res := solve(b, N, Q)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const MOD = 1000000007

func solve(B int, N []int, Q []int) []int {
	n := len(N)
	sum := make([]int, n)
	res := make([]int, n)
	sum[B-1] = 1
	res[B-1] = 1
	for i := B - 2; i >= 0; i-- {
		var sub int
		if i+1+N[i] < n {
			sub = sum[i+1+N[i]]
		}
		res[i] = modAdd(sum[i+1], MOD-sub)
		sum[i] = modAdd(sum[i+1], res[i])
	}
	ans := make([]int, len(Q))

	for i, cur := range Q {
		ans[i] = res[cur-1]
	}

	return ans
}

func solve1(B int, N []int, Q []int) []int {
	n := len(N)

	bit := NewBit(n)

	bit.Add(B, 1)

	F := make([]int, B+1)

	for i := B - 1; i > 0; i-- {
		r := n

		if i < n {
			r = (i + 1) + N[i]
		}

		F[i] = bit.Get(r)
		bit.Add(i, F[i])
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		ans[i] = F[cur]
	}

	return ans
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

type Bit struct {
	arr []int
	n   int
}

func NewBit(n int) Bit {
	arr := make([]int, n+1)
	return Bit{arr, n}
}

func (bit *Bit) Add(p int, v int) {

	for p <= bit.n {
		bit.arr[p] = modAdd(bit.arr[p], v)
		p += p & -p
	}

}

func (bit *Bit) Get(p int) int {
	var res int
	for p > 0 {
		res = modAdd(res, bit.arr[p])
		p -= p & -p
	}
	return res
}
