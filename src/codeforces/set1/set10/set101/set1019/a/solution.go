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

	n, m := readTwoNums(reader)
	voters := make([][]int, n)
	for i := 0; i < n; i++ {
		voters[i] = readNNums(reader, 2)
	}

	res := solve(m, voters)

	fmt.Println(res)
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

const oo = 1 << 60

func solve(m int, voters [][]int) int {
	if m == 1 {
		// only one party, no expense
		return 0
	}

	n := len(voters)

	parties := make([][]int, m)
	for i := 0; i < m; i++ {
		parties[i] = make([]int, 0, 1)
	}
	for _, cur := range voters {
		parties[cur[0]-1] = append(parties[cur[0]-1], cur[1])
	}
	win := len(parties[0])

	parties = parties[1:]

	sum := make([][]int, m-1)
	for i := 0; i < m-1; i++ {
		sort.Ints(parties[i])
		sum[i] = make([]int, len(parties[i])+1)
		for j := 0; j < len(parties[i]); j++ {
			sum[i][j+1] = sum[i][j] + parties[i][j]
		}
	}

	check := func(expect int) int {
		var expense int
		var cnt int
		pq := make(IntHeap, 0, n)
		for i := 0; i < len(parties); i++ {
			if len(parties[i]) > expect {
				expense += sum[i][len(parties[i])-expect]
				cnt += len(parties[i]) - expect
				for j := len(parties[i]) - expect; j < len(parties[i]); j++ {
					heap.Push(&pq, parties[i][j])
				}
			} else {
				for j := 0; j < len(parties[i]); j++ {
					heap.Push(&pq, parties[i][j])
				}
			}
		}

		for win+cnt <= expect {
			expense += heap.Pop(&pq).(int)
			cnt++
		}

		return expense
	}

	best := oo

	for x := 0; x <= n/2; x++ {
		best = min(best, check(x))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
