package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readUInt64(bytes []byte, from int, val *uint64) int {
	i := from
	tmp := uint64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var low, high int64

	scanner.Scan()
	pos := readInt64(scanner.Bytes(), 0, &low)
	readInt64(scanner.Bytes(), pos+1, &high)
	fmt.Println(solve(low, high))
}

const MAX_N = 1000005

func solve(low, high int64) int64 {
	set := make([]bool, MAX_N)

	divisors := make([][]int, MAX_N)

	for i := 0; i < MAX_N; i++ {
		// about max log(max_n) = 20
		divisors[i] = make([]int, 0, 10)
	}

	divCeil := func(i int) int64 {
		I := int64(i)
		return (low + I - 1) / I * I
	}

	for i := 2; i < MAX_N; i++ {
		if !set[i] {
			for j := 2 * i; j < MAX_N; j += i {
				set[j] = true
			}

			for x := divCeil(i); x <= high; x += int64(i) {
				divisors[x-low] = append(divisors[x-low], i)
			}
		}
	}

	var ans int64

	que := make(IntHeap, 0, 20)

	for x := low; x <= high; x++ {
		tmp := x
		for _, p := range divisors[x-low] {
			P := int64(p)
			var cnt int
			for tmp%P == 0 {
				cnt++
				tmp /= P
			}
			heap.Push(&que, cnt)
		}

		if tmp > 1 {
			// tmp is is prime
			heap.Push(&que, 1)
		}

		for que.Len() > 0 {
			here := int64(1)
			for _, a := range que {
				here *= int64(a + 1)
			}
			ans += here
			b := heap.Pop(&que).(int)
			if b > 1 {
				heap.Push(&que, b-1)
			}
		}
	}

	return ans
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
