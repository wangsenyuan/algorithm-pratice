package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("possible")
	} else {
		fmt.Println("impossible")
	}
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

func process(reader *bufio.Reader) bool {
	m, n := readTwoNums(reader)
	days := make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		days[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &days[i][j]) + 1
		}
	}
	return solve(n, days)
}

func solve(n int, days [][]int) bool {
	m := len(days)

	sets := make([]*BitSet, m)
	for i := 0; i < m; i++ {
		sets[i] = NewBitSet(n)
		for j := 0; j < len(days[i]); j++ {
			sets[i].Set(days[i][j] - 1)
		}
	}

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			tmp := sets[i].Copy()
			tmp.And(sets[j])
			if tmp.IsEmpty() {
				return false
			}
		}
	}
	return true
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) IsEmpty() bool {
	for i := 0; i < bs.sz; i++ {
		if bs.arr[i] != 0 {
			return false
		}
	}
	return true
}

func (bs *BitSet) Copy() *BitSet {
	res := NewBitSet(len(bs.arr) * 64)
	copy(res.arr, bs.arr)
	return res
}

func (this *BitSet) And(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] &= that.arr[i]
	}
}
