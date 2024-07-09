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

	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	res, steps := solve(a, b)
	if res < 0 {
		fmt.Println(res)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", res))
	for i := 0; i < res; i++ {
		buf.WriteString(fmt.Sprintf("%d ", steps[i]))
	}
	buf.WriteByte('\n')
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

const inf = 1 << 50

type pair struct {
	first  int
	second int
}

func (this pair) min(that pair) pair {
	if this.first < that.first {
		return this
	}
	return that
}

func solve(a []int, b []int) (int, []int) {
	n := len(a)
	var ans int
	l, r := n-a[n-1], n
	from := make([]int, n+1)
	for l > 0 {
		ans++
		p := r
		for i := l; i <= r; i++ {
			if p > i+b[i-1]-a[i+b[i-1]-1] {
				p = i + b[i-1] - a[i+b[i-1]-1]
				from[ans] = i
			}
		}
		if p >= l {
			return -1, nil
		}
		r = l + 1
		l = p
	}
	ans++
	return ans, from[1 : ans+1]
}

func solve1(a []int, b []int) (int, []int) {
	n := len(a)
	a = shift(a)
	b = shift(b)
	// 这里有个重要的地方，就是如果位置i -> j
	// 那么从位置i+1， 即使能跳跃到j，也不应该跳跃过去
	items := make([]*Item, n+1)
	pq := make(PriorityQueue, n+1)
	for i := 0; i <= n; i++ {
		it := new(Item)
		it.id = i
		it.index = i
		it.priority = inf
		items[i] = it
		pq[i] = it
	}

	heap.Init(&pq)

	pq.update(items[n], 0)

	l := n

	jumpFrom := make([]int, n+1)
	for i := 0; i <= n; i++ {
		jumpFrom[i] = -1
	}

	flipFrom := make([]int, n+1)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority == inf {
			break
		}
		u := it.id
		nl := u - a[u]
		if nl == 0 {
			pq.update(items[0], it.priority+1)
			jumpFrom[0] = u
			break
		}

		for i := nl; i < l && i < u; i++ {
			jumpFrom[i] = u
			j := i + b[i]
			if items[j].priority > it.priority+1 {
				flipFrom[j] = i
				pq.update(items[j], it.priority+1)
			}
		}
		l = min(l, nl)
	}

	if jumpFrom[0] < 0 || items[0].priority == inf {
		return -1, nil
	}
	steps := []int{0}
	cur := jumpFrom[0]

	for cur != n && len(steps) < items[0].priority {
		cur = flipFrom[cur]
		steps = append(steps, cur)
		cur = jumpFrom[cur]
	}

	reverse(steps)

	return len(steps), steps
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func shift(a []int) []int {
	arr := make([]int, len(a)+1)
	copy(arr[1:], a)
	return arr
}
