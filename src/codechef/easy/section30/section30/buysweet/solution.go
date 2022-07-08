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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, R := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(R, A, B)
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
		if s[i] == '\n' {
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

func solve(R int, A []int, B []int) int {
	n := len(A)

	sweets := make([]Sweet, n)

	for i := 0; i < n; i++ {
		sweets[i] = Sweet{A[i], B[i]}
	}

	sort.Sort(Sweets(sweets))
	var res int
	for i := 0; i < n; i++ {
		if sweets[i].price > R {
			continue
		}
		cost := sweets[i].price - sweets[i].back
		x := (R - sweets[i].price) / cost
		if R-x*cost >= sweets[i].price {
			x++
		}
		res += x
		R -= x * cost
	}

	return res
}

type Sweet struct {
	price int
	back  int
}

type Sweets []Sweet

func (this Sweets) Len() int {
	return len(this)
}

func (this Sweets) Less(i, j int) bool {
	return this[i].price-this[i].back < this[j].price-this[j].back
}

func (this Sweets) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
