package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		b := readNNums(reader, n)
		a := solve(b)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", a[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(b []int) []int {
	// b[i] = floor(i / a[i])
	n := len(b)
	// 这样会出现重复的部分
	// a[i] = 1, b[i] = i
	// 对于i, 根据b[i]，可以获得a[i]的一个范围, 这个可以通过二分的方式获得
	// 0, 2, 0, 1
	// [2...4], [1, 1], [4, 4], [3, 4]
	evs := make([]Pair, n)
	for i := 0; i < n; i++ {
		evs[i] = Pair{(i+1)/(b[i]+1) + 1, i}
	}
	sort.Slice(evs, func(i, j int) bool {
		return evs[i].first < evs[j].first || evs[i].first == evs[j].first && evs[i].second < evs[j].second
	})

	pq := make(PQ, 0, n)
	ans := make([]int, n)
	var j int
	for x := 1; x <= n; x++ {
		for j < n && evs[j].first == x {
			id := evs[j].second
			j++
			if b[id] == 0 {
				heap.Push(&pq, Item{id, n})
			} else {
				heap.Push(&pq, Item{id, (id + 1) / b[id]})
			}
		}
		cur := heap.Pop(&pq).(Item)
		ans[cur.id] = x
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

type Item struct {
	id    int
	value int
}

func (this Item) Less(that Item) bool {
	return this.value < that.value || this.value == that.value && this.id < that.id
}

type PQ []Item

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PQ) Push(x interface{}) {
	it := x.(Item)
	*this = append(*this, it)
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	return res
}
