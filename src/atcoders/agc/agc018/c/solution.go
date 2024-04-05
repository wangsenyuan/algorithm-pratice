package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	x, y, z := readThreeNums(reader)
	people := make([][]int, x+y+z)
	for i := 0; i < len(people); i++ {
		people[i] = readNNums(reader, 3)
	}

	res := solve(people, x, y, z)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(people [][]int, x int, y int, z int) int64 {
	n := len(people)

	id := make([]int, n)

	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		a := people[id[i]][0] - people[id[i]][1]
		b := people[id[j]][0] - people[id[j]][1]
		return a > b
	})
	dp := make([]int64, n)
	pq := make(IntHeap, 0, x+1)
	var sum int64
	for j, i := range id {
		v := int64(people[i][0] - people[i][2])
		sum += v
		heap.Push(&pq, v)
		dp[j] = -inf
		if pq.Len() >= x {
			if pq.Len() > x {
				sum -= heap.Pop(&pq).(int64)
			}
			dp[j] = sum
		}
	}

	pq = make(IntHeap, 0, y+1)

	var best int64 = -inf

	sum = 0
	for j := n - 1; j >= x; j-- {
		i := id[j]
		v := int64(people[i][1] - people[i][2])
		sum += v
		heap.Push(&pq, v)
		if pq.Len() >= y {
			if pq.Len() > y {
				sum -= heap.Pop(&pq).(int64)
			}
			if sum+dp[j-1] > best {
				best = sum + dp[j-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		best += int64(people[i][2])
	}

	return best
}

// An IntHeap is a min-heap of ints.
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
