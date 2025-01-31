package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		res := process(reader)
		if res < 0 {
			break
		}
		fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	s, err := reader.ReadBytes('\n')
	if err != nil {
		return -1
	}
	var n int
	pos := readInt(s, 0, &n) + 1
	goods := make([][]int, n)
	for i := 0; i < n; i++ {
		goods[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			pos = readInt(s, pos, &goods[i][j]) + 1
		}
	}
	return solve(goods)
}

func solve(goods [][]int) int {
	n := len(goods)

	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		arr[i] = pair{goods[i][0], goods[i][1]}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.second - b.second
	})

	var res int
	m := arr[n-1].second
	active := make(IntHeap, 0, n)
	for i, j := m, n-1; i > 0; i-- {
		for j >= 0 && arr[j].second == i {
			heap.Push(&active, arr[j].first)
			j--
		}
		if active.Len() > 0 {
			res += heap.Pop(&active).(int)
		}
	}
	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
