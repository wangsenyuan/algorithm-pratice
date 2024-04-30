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

	n, T := readTwoNums(reader)
	people := make([][]int, n)
	for i := 0; i < n; i++ {
		people[i] = readNNums(reader, 2)
	}

	score, res := solve(T, people)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%d\n", score, len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(T int, people [][]int) (int, []int) {
	n := len(people)
	set := make([]Pair, n)
	for i := 0; i < n; i++ {
		set[i] = Pair{people[i][0], i}
	}

	sort.Slice(set, func(i, j int) bool {
		return set[i].first > set[j].first
	})
	active := make(PriorityQueue, 0, n)
	var time int
	for k, i := n, 0; k > 0; k-- {
		for i < n && set[i].first >= k {
			j := set[i].second
			heap.Push(&active, Pair{people[j][1], j})
			time += people[j][1]
			i++
		}

		for active.Len() > k {
			time -= heap.Pop(&active).(Pair).first
		}

		if active.Len() == k && time <= T {
			ans := make([]int, active.Len())
			for j := 0; j < len(active); j++ {
				ans[j] = active[j].second + 1
			}
			return k, ans
		}
	}
	return 0, nil
}
func solve1(T int, people [][]int) (int, []int) {
	n := len(people)

	ans := make([]int, n)

	check := func(k int) bool {
		// 是否可以选择k个人，使得他们的总时间 <= T 且 a[i] >= k
		active := make(PriorityQueue, 0, n)
		var time int
		for i := 0; i < n; i++ {
			if people[i][0] < k {
				continue
			}
			// a[i] >= k
			time += people[i][1]
			heap.Push(&active, Pair{people[i][1], i})

			for time > T {
				time -= heap.Pop(&active).(Pair).first
			}
			for active.Len() > k {
				time -= heap.Pop(&active).(Pair).first
			}
		}

		ok := time <= T && len(active) == k

		if ok {
			for i := 0; i < k; i++ {
				ans[i] = active[i].second + 1
			}
		}

		return ok
	}

	if check(n) {
		return n, ans
	}

	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	r--
	// do one more check to create the ans
	check(r)

	return r, ans[:r]
}

type Pair struct {
	first  int
	second int
}

type PriorityQueue []Pair

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].first > h[j].first }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Pair))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
