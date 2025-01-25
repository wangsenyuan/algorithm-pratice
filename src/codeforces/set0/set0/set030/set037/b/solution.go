package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, time, res := process(reader)
	if time < 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n%d %d\n", time, len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (hp int, reg int, scrolls [][]int, time int, res [][]int) {
	n, hp, reg := readThreeNums(reader)
	scrolls = make([][]int, n)
	for i := range n {
		scrolls[i] = readNNums(reader, 2)
	}
	time, res = solve(hp, reg, scrolls)
	return
}

func solve(hp int, reg int, scrolls [][]int) (time int, res [][]int) {

	n := len(scrolls)
	items := make([]item, n)
	for i := range n {
		items[i] = item{scrolls[i][0], scrolls[i][1], i}
	}

	slices.SortFunc(items, func(a, b item) int {
		if a.pow != b.pow {
			return b.pow - a.pow
		}
		if a.dmg != b.dmg {
			return b.dmg - a.dmg
		}
		return a.id - b.id
	})

	if items[0].pow < 100 {
		return -1, nil
	}
	res = append(res, []int{0, items[0].id + 1})
	sum := items[0].dmg
	i := 1
	cur := hp

	active := make(PQ, 0, n)

	for {
		time++
		old := cur
		cur -= sum
		cur = min(hp, cur+reg)
		if cur <= 0 {
			return
		}
		for i < n && items[i].pow*hp >= cur*100 {
			heap.Push(&active, items[i])
			i++
		}

		if old <= cur && active.Len() == 0 {
			break
		}

		if active.Len() > 0 {
			it := heap.Pop(&active).(item)
			res = append(res, []int{time, it.id + 1})
			sum += it.dmg
		}
	}
	return -1, nil
}

type item struct {
	pow int
	dmg int
	id  int
}

type PQ []item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].dmg > h[j].dmg }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *PQ) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
