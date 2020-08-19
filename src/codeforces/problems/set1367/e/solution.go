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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		fmt.Println(solve(n, k, s[:n]))
	}
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

func solve(n, k int, s string) int {
	cnt := make(map[byte]int)

	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}

	cycles := make([]int, n)

	used := make([]bool, n)

	for l := n; l > 0; l-- {
		var p int
		for i := 0; i < l; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			var x = 1

			j := (i + k) % l

			for !used[j] {
				x++
				used[j] = true
				j = (j + k) % l
			}
			cycles[p] = x
			p++
		}

		canFill := true

		sort.Ints(cycles[:p])
		pq := make(IntHeap, 0, len(cnt))

		for _, v := range cnt {
			heap.Push(&pq, v)
		}

		for p > 0 {
			cur := cycles[p-1]
			p--

			top := heap.Pop(&pq).(int)

			if top < cur {
				canFill = false
				break
			}

			heap.Push(&pq, top-cur)
		}

		if canFill {
			return l
		}

		for i := 0; i < n; i++ {
			used[i] = false
		}
	}
	return 0
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve1(n, k int, s string) int {
	// find all divisors of k
	cnt := make(map[byte]int)

	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}

	arr := make([]int, 0, len(cnt))
	for _, v := range cnt {
		arr = append(arr, v)
	}

	sort.Ints(arr)

	check := func(l, k int) bool {
		// can we have a l length segments length, each seg has k beads
		for i := len(arr) - 1; i >= 0 && arr[i] >= l && k > 0; i-- {
			x := arr[i] / l
			k -= x
		}

		return k <= 0
	}

	find := func(k int) int {
		// one segment has k
		var left, right = 1, n + 1

		for left < right {
			mid := (left + right) / 2
			if !check(mid, k) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		right--
		return right * k
	}

	best := 1
	for x := 1; x*x <= k; x++ {
		if k%x != 0 {
			continue
		}

		best = max(best, find(x))

		if k/x != x {
			best = max(best, find(k/x))
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
