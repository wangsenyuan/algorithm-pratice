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
		a := readNNums(reader, n)

		res := solve(a)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(a []int) bool {
	// 从leaf开始，每个节点进入的边的weight是可以确定的
	// a[i] > a[i+1] 那么i的入边=1, else 0
	// 但是这里有个问题，就是遇到 [1, 0, 1] 的时候，
	// 直观的知道应该是可行的，但规则是什么？
	// 为什么是 1, 0-(0, 1)的组合而不是 0 -（1， 0）, 1的组合？
	// 似乎也没有问题？
	// 还是没有抓住题眼
	// 好像从最大的a[i]开始是有戏的
	// 这是因为，最大的那个a[i],和它对应的，必然是 a[i] - 1, 且只能在一个方向，如果两边是a[i] - 1 => false
	// 找到这样的最大值后， 替换成 a[i]-1，并消耗掉一个 a[i]-1
	n := len(a)

	prev := make([]int, n)
	next := make([]int, n)

	for i := 0; i < n; i++ {
		prev[i] = i - 1
		next[i] = i + 1
	}

	check := func(i int) bool {
		if (prev[i] >= 0 && a[prev[i]] == a[i]-1) || (next[i] < n && a[next[i]] == a[i]-1) {
			return true
		}
		return false
	}

	in := make([]bool, n)
	pq := make(PQ, 0, n)

	for i := 0; i < n; i++ {
		if check(i) {
			in[i] = true
			it := new(Item)
			it.id = i
			it.priority = a[i]
			heap.Push(&pq, it)
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		i := it.id
		if prev[i] >= 0 {
			next[prev[i]] = next[i]
		}
		if next[i] < n {
			prev[next[i]] = prev[i]
		}
		if prev[i] >= 0 && check(prev[i]) && !in[prev[i]] {
			in[prev[i]] = true
			tmp := new(Item)
			tmp.id = prev[i]
			tmp.priority = a[prev[i]]
			heap.Push(&pq, tmp)
		}
		if next[i] < n && check(next[i]) && !in[next[i]] {
			in[next[i]] = true
			tmp := new(Item)
			tmp.id = next[i]
			tmp.priority = a[next[i]]
			heap.Push(&pq, tmp)
		}
	}
	var bad int
	mn := n
	for i := 0; i < n; i++ {
		if !in[i] {
			bad++
		}
		mn = min(mn, a[i])
	}

	return mn == 0 && bad == 1
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
	return pq[i].priority > pq[j].priority
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
