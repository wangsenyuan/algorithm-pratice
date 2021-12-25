package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, k, A))
	}
}

const MAX_P = 30

func solve(n int, K int, A []int) int {
	items := make([]Item, MAX_P)
	for p := MAX_P - 1; p >= 0; p-- {
		var cnt int
		for i := 0; i < n; i++ {
			cnt += (A[i] >> uint(p)) & 1
		}
		// if this position is 1, then it will contributes cnt * (1 << uint(p))
		items[p] = Item{p, int64(cnt) * int64(1<<uint(p))}
	}

	sort.Sort(Items(items))

	var res int

	for i := 0; i < K && i < len(items); i++ {
		res |= 1 << uint(items[i].pos)
	}
	return res
}

type Item struct {
	pos int
	val int64
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].val > items[j].val || items[i].val == items[j].val && items[i].pos < items[j].pos
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
