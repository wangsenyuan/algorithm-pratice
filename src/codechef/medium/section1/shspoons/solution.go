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
		n := readNum(reader)
		A := make([]int64, n)
		var pos int
		var tmp uint64
		s, _ := reader.ReadBytes('\n')
		for i := 0; i < n; i++ {
			pos = readUint64(s, pos, &tmp) + 1
			A[i] = int64(tmp)
		}
		res := solve(n, A)

		if len(res) == 0 {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", res[i].first, res[i].second, res[i].x))
		}
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

func solve(n int, A []int64) []Result {
	ps := make(Pairs, 0, n)

	for i := 1; i < n; i++ {
		heap.Push(&ps, Pair{A[i], i + 1})
	}
	res := make([]Result, 0, n)

	var cur = A[0]

	for ps.Len() > 0 {
		for ps.Len() > 0 && ps[0].first <= cur {
			cur += ps[0].first
			res = append(res, Result{ps[0].second, 1, ps[0].first})
			heap.Pop(&ps)
		}

		if ps.Len() < 2 {
			break
		}

		tmp1 := heap.Pop(&ps).(Pair)
		tmp2 := heap.Pop(&ps).(Pair)
		res = append(res, Result{tmp1.second, tmp2.second, tmp1.first - cur})

		heap.Push(&ps, tmp2.changeFirst(tmp2.first+tmp1.first-cur))

		heap.Push(&ps, tmp1.changeFirst(cur))
	}

	if ps.Len() > 0 {
		return nil
	}
	return res
}

type Result struct {
	first, second int
	x             int64
}

type Pair struct {
	first  int64
	second int
}

func (this Pair) changeFirst(first int64) Pair {
	return Pair{first, this.second}
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps *Pairs) Push(x interface{}) {
	*ps = append(*ps, x.(Pair))
}

func (ps *Pairs) Pop() interface{} {
	n := len(*ps)
	res := (*ps)[n-1]
	*ps = (*ps)[:n-1]
	return res
}
