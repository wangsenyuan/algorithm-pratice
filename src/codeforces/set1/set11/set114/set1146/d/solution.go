package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m, a, b := readThreeNums(reader)
	res := solve(a, b, m)
	fmt.Println(res)
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

func solve(a int, b int, m int) int {
	n := min(m, a+b-1)
	g := gcd(a, b)
	dist := make([]int, n+1)
	for i := range n + 1 {
		dist[i] = n + 1
	}
	dist[0] = 0
	pq := make(PriorityQueue, 0, n)
	heap.Push(&pq, pair{0, 0})

	sum := make([]int, n+1)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(pair)
		x := cur.first
		i := cur.second
		if dist[i] != x {
			continue
		}
		sum[x]++
		if i >= b && dist[i-b] > x {
			dist[i-b] = x
			heap.Push(&pq, pair{x, i - b})
		}
		if i+a <= n && max(x, i+a) < dist[i+a] {
			dist[i+a] = max(x, i+a)
			heap.Push(&pq, pair{dist[i+a], i + a})
		}
	}

	var ans int
	for i := 0; i <= n; i++ {
		if i > 0 {
			sum[i] += sum[i-1]
		}
		ans += sum[i]
	}
	if n != m {
		for (m+1)%g > 0 {
			ans += m/g + 1
			m--
		}
		ans += g * (f(m/g+1) - f(n/g+1))
	}
	return ans
}

func f(x int) int {
	return x * (x + 1) / 2
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type pair struct {
	first  int
	second int
}

type PriorityQueue []pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first || pq[i].first == pq[j].first && pq[i].second < pq[j].second
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(pair))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
