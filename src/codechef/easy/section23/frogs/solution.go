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
		n := readNum(reader)
		W := readNNums(reader, n)
		L := readNNums(reader, n)
		res := solve(n, W, L)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, W, L []int) int {
	items := make(Items, n)
	P := make([]*Item, n+1)
	for i := 0; i < n; i++ {
		item := new(Item)
		item.pos = i + 1
		item.jumpDist = L[i]
		item.weight = W[i]
		item.index = i
		items[i] = item
		P[W[i]] = item
	}

	sort.Sort(items)

	var res int

	for i := 1; i <= n; {
		if P[i].index == i-1 {
			i++
			continue
		}
		for j := i + 1; j <= n; j++ {
			if P[j].index <= i-1 {
				P[j].pos += P[j].jumpDist
				res++
			}
		}
		sort.Sort(items)
	}
	return res
}

type Item struct {
	pos, jumpDist, weight int
	index                 int
}

type Items []*Item

func (this Items) Len() int {
	return len(this)
}

func (this Items) Less(i, j int) bool {
	return this[i].pos < this[j].pos || this[i].pos == this[j].pos && this[i].weight > this[j].weight
}

func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}
