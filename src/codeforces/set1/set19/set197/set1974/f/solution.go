package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		nums := readNNums(reader, 4)
		a, b, n, m := nums[0], nums[1], nums[2], nums[3]
		clips := make([][]int, n)
		for i := 0; i < n; i++ {
			clips[i] = readNNums(reader, 2)
		}
		ops := make([]string, m)
		for i := 0; i < m; i++ {
			ops[i] = readString(reader)
		}
		res := solve(a, b, clips, ops)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}

	fmt.Print(buf.String())
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

const inf = 1 << 60

func solve(a, b int, clips [][]int, operations []string) []int {
	top := NewStore(clips, func(clip []int) int {
		return clip[0]
	})

	bot := NewStore(clips, func(clip []int) int {
		return -clip[0]
	})

	lft := NewStore(clips, func(clip []int) int {
		return clip[1]
	})

	rgt := NewStore(clips, func(clip []int) int {
		return -clip[1]
	})

	stores := []*Store{top, bot, lft, rgt}
	ids := map[byte]int{'U': 0, 'D': 1, 'L': 2, 'R': 3}

	rect := []int{1, 1, a, b}

	play := func(s string) int {
		cmd := s[0]
		var k int
		readInt([]byte(s), 2, &k)

		if cmd == 'U' {
			rect[0] += k
		} else if cmd == 'D' {
			rect[2] -= k
		} else if cmd == 'L' {
			rect[1] += k
		} else {
			rect[3] -= k
		}

		i := ids[cmd]

		var cnt int

		for !stores[i].IsEmpty() {
			first := stores[i].First()
			clip := clips[first.id]

			if clip[0] >= rect[0] && clip[0] <= rect[2] && clip[1] >= rect[1] && clip[1] <= rect[3] {
				// in the left area
				break
			}
			cnt++

			for j := 0; j < 4; j++ {
				stores[j].Remove(first.id)
			}
		}

		return cnt
	}

	res := make([]int, 2)

	for i, cur := range operations {
		res[i&1] += play(cur)
	}

	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func (pq *PQ) remove(it *Item) {
	pq.update(it, -inf)
	heap.Pop(pq)
}

type Store struct {
	arr []*Item
	pq  *PQ
}

func NewStore(clips [][]int, pf func([]int) int) *Store {
	n := len(clips)

	arr := make([]*Item, n)
	pq := make(PQ, n)

	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.index = i
		it.priority = pf(clips[i])
		arr[i] = it
		pq[i] = it
	}

	heap.Init(&pq)

	return &Store{arr, &pq}
}

func (store *Store) IsEmpty() bool {
	return store.pq.Len() == 0
}

func (store *Store) First() *Item {
	return (*(store.pq))[0]
}

func (store *Store) Remove(id int) {
	store.pq.remove(store.arr[id])
}
