package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	H := make([][]int, n)
	for i := 0; i < n; i++ {
		H[i] = readNNums(reader, 3)
	}
	V := make([][]int, m)
	for i := 0; i < m; i++ {
		V[i] = readNNums(reader, 3)
	}
	res := solve(H, V)

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

const MAX = 1000_000

func solve(H [][]int, V [][]int) int {

	res := 1

	events := make([][]Pair, MAX+1)

	addEvent := func(i int, evt Pair) {
		if len(events[i]) == 0 {
			events[i] = make([]Pair, 0, 1)
		}
		events[i] = append(events[i], evt)
	}

	for _, cur := range H {
		if cur[2] == MAX && cur[1] == 0 {
			res++
		}
		// cur[0] is y
		addEvent(cur[1], Pair{cur[0], 0})
		addEvent(cur[2], Pair{cur[0], 1})
	}

	queries := make([][]Pair, MAX+1)

	addQuery := func(i int, qry Pair) {
		if queries[i] == nil {
			queries[i] = make([]Pair, 0, 1)
		}
		queries[i] = append(queries[i], qry)
	}

	for _, cur := range V {
		// cur[1] & cur[2] are y
		addQuery(cur[0], Pair{cur[1], cur[2]})
	}

	// set 表示仍然活跃的y的值
	set := NewBit(MAX + 1)
	set.Add(1, 1)
	set.Add(MAX+1, 1)

	for i := 0; i <= MAX; i++ {
		for _, it := range events[i] {
			if it.second == 0 {
				set.Add(it.first+1, 1)
			}
		}
		for _, it := range queries[i] {
			tmp := set.Get(it.second+1) - set.Get(it.first)
			res += tmp - 1
		}

		for _, it := range events[i] {
			if it.second == 1 {
				set.Add(it.first+1, -1)
			}
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Bit struct {
	arr []int
}

func NewBit(n int) *Bit {
	arr := make([]int, n+3)
	return &Bit{arr}
}

func (b *Bit) Add(p int, v int) {
	for i := p; i < len(b.arr); i += i & -i {
		b.arr[i] += v
	}
}

func (b *Bit) Get(p int) int {
	var res int
	for p > 0 {
		res += b.arr[p]
		p -= p & -p
	}
	return res
}
