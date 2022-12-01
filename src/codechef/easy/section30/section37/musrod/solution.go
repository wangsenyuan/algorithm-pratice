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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int, B []int) int64 {
	n := len(A)
	rods := make([]Rod, n)

	for i := 0; i < n; i++ {
		rods[i] = Rod{A[i], B[i]}
	}

	sort.Slice(rods, func(i, j int) bool {
		return rods[i].Less(rods[j])
	})

	var res int64
	var cur int64
	for i := 0; i < n; i++ {
		res += cur * int64(rods[i].b)
		cur += int64(rods[i].a)
	}

	return res
}

type Rod struct {
	a int
	b int
}

// 假设在排列中 i在j的前面 它们的贡献是 xi * bi + xj * bj
// 假设交换后
// xi * bj + (xj - ai + aj) * bi
// 下式减去上式
// xi * (bj - bi) + xj * (bi - bj) + (aj - ai) * bi
// 使 (xj - xi) * (bi - bj) + (aj - ai) * bi > 0
// xj - xi = ai
// ai * bi - ai * bj + aj * bi - ai * bi > 0
// aj * bi > ai * bj
func (this Rod) Less(that Rod) bool {
	x := int64(this.a) * int64(that.b)
	y := int64(that.a) * int64(this.b)
	return x >= y
}
