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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)

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

func solve(n int, A []int) int {
	cnt := make([]int, n+1)
	var mx int
	for i := 0; i < n; i++ {
		cnt[A[i]]++
		if cnt[A[i]] > mx {
			mx = cnt[A[i]]
		}
	}

	var y, k int

	for i := 1; i <= n; i++ {
		if cnt[i] == mx {
			k++
		} else {
			y += cnt[i]
		}
	}

	return y/(mx-1) + k - 1
}

const N = 100010

var B [N]int

func solve1(n int, A []int) int {

	check := func(dist int) bool {
		if dist == 1 {
			return true
		}
		cnt := make(map[int]int)
		for i := 0; i < n; i++ {
			cnt[A[i]]++
		}

		ps := make(Pairs, 0, len(cnt))

		for k, v := range cnt {
			p := Pair{v, k, len(ps) - 1}
			heap.Push(&ps, &p)
		}

		for i := 0; i < n; i++ {
			if i >= dist && cnt[B[i-dist]] > 0 {
				p := Pair{cnt[B[i-dist]], B[i-dist], len(ps) - 1}
				heap.Push(&ps, &p)
			}
			if len(ps) == 0 {
				return false
			}
			cur := heap.Pop(&ps).(*Pair)
			B[i] = cur.second
			cnt[cur.second]--
		}

		return true
	}

	var left = 2
	var right = n
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 2
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first, second int
	index         int
}

type Pairs []*Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first || ps[i].first == ps[j].first && ps[i].second > ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (pq *Pairs) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Pair)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Pairs) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
