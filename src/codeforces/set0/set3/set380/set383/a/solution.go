package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	a := readNNums(reader, n)

	res := solve(a)

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

func solve(a []int) int {
	var sum int
	var res int
	for _, x := range a {
		if x == 0 {
			res += sum
		} else {
			sum++
		}
	}
	return res
}

func solve1(a []int) int {
	n := len(a)
	var zeros []int
	var ones []int
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			zeros = append(zeros, i)
		} else {
			ones = append(ones, i)
		}
	}

	reverse(zeros)

	id := zeros
	id = append(id, ones...)

	tr := NewSegTree(n)
	var res int
	for _, i := range id {
		if a[i] == 1 {
			res += tr.Get(i, n)
		} else {
			res += tr.Get(0, i)
		}
		tr.Set(i, 1)
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)

	return &SegTree{arr, n}
}

func (t *SegTree) Set(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.n
	r += t.n

	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
