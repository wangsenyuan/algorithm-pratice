package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, m := readTwoNums(scanner)
		Z := readNNums(scanner, n)
		L, R, K := make([]int, m), make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			L[i], R[i], K[i] = readThreeNums(scanner)
		}
		res := solve(n, m, Z, L, R, K)
		if res < 0 {
			fmt.Println("NO")
		} else {
			fmt.Printf("YES %d\n", res)
		}
		t--
	}
}

func solve(n, m int, Z []int, L []int, R []int, K []int) int {
	cranes := make(CA, m)

	for i := 0; i < m; i++ {
		cranes[i] = &Crane{L[i] - 1, R[i] - 1, K[i]}
	}

	sort.Sort(cranes)

	h := make(CQ, 0, m)
	heap.Init(&h)

	diff := make([]int, n+1)
	var ans int
	for i, j := 0, 0; i < n; {
		if Z[i]+diff[i] > 0 {
			for j < m && cranes[j].left <= i {
				heap.Push(&h, cranes[j])
				j++
			}

			if h.Len() == 0 || h[0].right < i {
				return -1
			}

			first := h[0]

			took := min(Z[i]+diff[i], first.beams)
			diff[i] -= took
			diff[first.right+1] += took
			first.beams -= took
			ans += took
			if first.beams == 0 {
				heap.Pop(&h)
			} else {
				heap.Fix(&h, 0)
			}
		}
		if Z[i]+diff[i] <= 0 {
			diff[i+1] += diff[i]
			i++
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Crane struct {
	left, right int
	beams       int
}

type CA []*Crane

func (ca CA) Len() int {
	return len(ca)
}

func (ca CA) Less(i, j int) bool {
	return ca[i].left < ca[j].left
}

func (ca CA) Swap(i, j int) {
	ca[i], ca[j] = ca[j], ca[i]
}

type CQ []*Crane

func (cq CQ) Len() int {
	return len(cq)
}

func (cq CQ) Less(i, j int) bool {
	return cq[i].right > cq[j].right
}

func (cq CQ) Swap(i, j int) {
	cq[i], cq[j] = cq[j], cq[i]
}

func (cq *CQ) Push(item interface{}) {
	crane := item.(*Crane)
	*cq = append(*cq, crane)
}

func (cq *CQ) Pop() interface{} {
	n := len(*cq)
	item := (*cq)[n-1]
	*cq = (*cq)[:n-1]
	return item
}
