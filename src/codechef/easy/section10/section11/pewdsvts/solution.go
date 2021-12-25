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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
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
		firstLine := readNNums(scanner, 6)
		N, A, B, X, Y, Z := firstLine[0], firstLine[1], firstLine[2], firstLine[3], firstLine[4], firstLine[5]
		C := readNNums(scanner, N)
		res := solve(N, A, B, X, Y, Z, C)
		if res >= 0 {
			fmt.Println(res)
		} else {
			fmt.Println("RIP")
		}
	}
}

func solve(N, A, B, X, Y, Z int, C []int) int {
	n := (Z - A + X - 1) / X
	m := (Z - B + Y - 1) / Y
	if n < m {
		// win
		return 0
	}

	// need to win before m day
	Z0 := A + (m - 1) * X
	need := Z - Z0

	pq := &IntHeap{}
	heap.Init(pq)

	for _, c := range C {
		heap.Push(pq, c)
	}
	var times int
	var got int
	for pq.Len() > 0 && got < need {
		h := heap.Pop(pq).(int)
		times++
		got += h
		if h / 2 > 0 {
			heap.Push(pq, h / 2)
		}
	}
	if got >= need {
		return times
	}
	return -1
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