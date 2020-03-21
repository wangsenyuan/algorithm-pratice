package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && bs[x] == ' ' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)
		V := readNNums(scanner, N)
		scanner.Scan()
		D := scanner.Text()
		U := readNNums(scanner, M)
		scanner.Scan()
		E := scanner.Text()
		fmt.Println(solve(N, V, D, M, U, E))
	}
}

func solve(N int, V []int, D string, M int, U []int, E string) int64 {
	a := NewPlayer(N, V, D)
	b := NewPlayer(M, U, E)

	var loop func(a, b *Player) int64

	loop = func(a, b *Player) int64 {
		best := a.value - b.value

		if a.attack.Len() == 0 || b.defense.Len() == 0 {
			return best
		}

		x := heap.Pop(a.attack).(int)
		y := heap.Pop(b.defense).(int)

		b.value -= int64(y)
		heap.Push(a.defense, x)

		tmp := -1 * loop(b, a)
		return max(best, tmp)
	}

	return loop(a, b)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Player struct {
	attack  *MinHeap
	defense *MaxHeap
	value   int64
}

func NewPlayer(n int, V []int, D string) *Player {
	attack := make(MinHeap, 0, n)
	defense := make(MaxHeap, 0, n)
	var sum int64
	heap.Init(&attack)
	heap.Init(&defense)
	for i := 0; i < n; i++ {
		if D[i] == 'A' {
			heap.Push(&attack, V[i])
		} else {
			heap.Push(&defense, V[i])
		}
		sum += int64(V[i])
	}

	return &Player{&attack, &defense, sum}
}

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
