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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	lines := make([][]int, n)
	for i := 0; i < n; i++ {
		lines[i] = readNNums(reader, 4)
	}
	cnt, res := solve(n, lines)

	fmt.Println(cnt)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", res[i])
	}
}

const MAX_Y = 100001

func solve(n int, lines [][]int) (int, []int) {
	res := sweep(n, lines)

	var cnt int

	for i := 0; i < n; i++ {
		cnt += res[i]
	}

	for i := 0; i < n; i++ {
		lines[i][0], lines[i][1] = lines[i][1], lines[i][0]
		lines[i][2], lines[i][3] = lines[i][3], lines[i][2]
	}

	res2 := sweep(n, lines)

	for i := 0; i < len(res2); i++ {
		if res2[i] > 0 {
			res[i] = res2[i]
		}
	}

	return cnt, res
}

func sweep(n int, lines [][]int) []int {
	cnt := make(map[Pair]int)

	items := make([]Item, 0, 2*n)
	for i := 0; i < n; i++ {
		x1, x2, y2 := lines[i][0], lines[i][2], lines[i][3]

		if x1 == x2 {
			// vertical
			items = append(items, Item{x1, i, 0})
		} else {
			// y1 == y2 horizontal
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			items = append(items, Item{x1, i, -1})
			items = append(items, Item{x2, i, 1})
			cnt[Pair{x1, y2}]++
			cnt[Pair{x2, y2}]++
		}
	}

	sort.Sort(Items(items))

	arr := make([]int, MAX_Y+1)

	update := func(pos int, v int) {
		pos++
		for pos <= MAX_Y {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		pos++
		var res int
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}

		return res
	}

	getRange := func(left int, right int) int {
		res := get(right)
		if left > 0 {
			res -= get(left - 1)
		}
		return res
	}

	res := make([]int, n)

	for _, item := range items {
		i := item.idx
		line := lines[i]

		if item.op == -1 {
			//y1 == y2
			y := line[1]
			update(y, 1)
		} else if item.op == 1 {
			y := line[1]
			update(y, -1)
		} else {
			// o
			x := line[0]
			y1, y2 := line[1], line[3]
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			res[i] = getRange(y1, y2)
			res[i] -= cnt[Pair{x, y1}]
			res[i] -= cnt[Pair{x, y2}]
		}
	}

	return res
}

type Pair struct {
	first, second int
}

type Item struct {
	x   int
	idx int
	op  int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].x < items[j].x || items[i].x == items[j].x && items[i].op < items[j].op
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
