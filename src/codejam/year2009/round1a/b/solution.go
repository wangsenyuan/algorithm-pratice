package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		N, M := readTwoNums(scanner)
		G := make([][]int, N)
		for i := 0; i < N; i++ {
			G[i] = readNNums(scanner, 3*M)
		}
		res := solve(N, M, G)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(N int, M int, G [][]int) int {
	cs := make([]*Corner, N*M*4)
	que := make(PriorityQueue, 0, N*M*4)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for k := 0; k < 4; k++ {
				c := new(Corner)
				c.i = i
				c.j = j
				c.k = k
				c.dist = math.MaxInt32
				cs[i*M*4+j*4+k] = c
				heap.Push(&que, c)
			}
		}
	}
	// the begin position
	que.update(cs[(N-1)*M*4+3], 0)

	travelNorth := func(i, j, k, cur int) {
		kk := 3 - k
		if k == 0 || k == 1 {
			if i > 0 {
				x := cs[(i-1)*4*M+4*j+kk]
				if x.dist > cur+2 {
					que.update(x, cur+2)
				}
			}
			return
		}
		S, W, T := G[i][3*j], G[i][3*j+1], G[i][3*j+2]

		gap := ((cur-T)%(S+W) + (S + W)) % (S + W)
		var tmp int
		if gap < S {
			tmp = cur + 1
		} else {
			tmp = cur + S + W - gap + 1
		}

		x := cs[i*4*M+j*4+kk]
		if x.dist > tmp {
			que.update(x, tmp)
		}
	}

	travelEast := func(i, j, k, cur int) {
		if k == 1 || k == 2 {
			if j < M-1 {
				kk := 0
				if k == 2 {
					kk = 3
				}
				x := cs[i*4*M+(j+1)*4+kk]
				if x.dist > cur+2 {
					que.update(x, cur+2)
				}
			}
			return
		}

		S, W, T := G[i][3*j], G[i][3*j+1], G[i][3*j+2]
		gap := ((cur-T)%(S+W) + (S + W)) % (S + W)
		var tmp int
		if gap < S {
			tmp = cur + S - gap + 1
		} else {
			tmp = cur + 1
		}
		kk := 1
		if k == 3 {
			kk = 2
		}

		x := cs[i*4*M+j*4+kk]
		if x.dist > tmp {
			que.update(x, tmp)
		}
	}

	travelSouth := func(i, j, k, cur int) {
		kk := 3 - k
		if k == 2 || k == 3 {
			if i < N-1 {
				x := cs[(i+1)*4*M+j*4+kk]
				if x.dist > cur+2 {
					que.update(x, cur+2)
				}
			}

			return
		}
		S, W, T := G[i][3*j], G[i][3*j+1], G[i][3*j+2]

		gap := ((cur-T)%(S+W) + (S + W)) % (S + W)
		var tmp int
		if gap < S {
			tmp = cur + 1
		} else {
			tmp = cur + S + W - gap + 1
		}

		// kk := 3 - k
		x := cs[i*4*M+j*4+kk]

		if x.dist > tmp {
			que.update(x, tmp)
		}
	}

	travelWest := func(i, j, k, cur int) {
		if k == 0 || k == 3 {
			if j > 0 {
				kk := 1
				if k == 3 {
					kk = 2
				}
				x := cs[i*4*M+(j-1)*4+kk]
				if x.dist > cur+2 {
					que.update(x, cur+2)
				}
			}

			return
		}

		S, W, T := G[i][3*j], G[i][3*j+1], G[i][3*j+2]
		gap := ((cur-T)%(S+W) + (S + W)) % (S + W)

		var tmp int
		if gap < S {
			tmp = cur + S - gap + 1
		} else {
			tmp = cur + 1
		}

		kk := 0
		if k == 2 {
			kk = 3
		}
		x := cs[i*4*M+j*4+kk]
		if x.dist > tmp {
			que.update(x, tmp)
		}
	}

	for que.Len() > 0 {
		head := heap.Pop(&que).(*Corner)
		i, j, k := head.i, head.j, head.k
		travelNorth(i, j, k, head.dist)
		travelEast(i, j, k, head.dist)
		travelSouth(i, j, k, head.dist)
		travelWest(i, j, k, head.dist)
	}
	return cs[(M-1)*4+1].dist
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func div(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return a/b - 1
}

type Corner struct {
	i, j  int // intercection
	k     int // corner index, 0 for n-w, 1 for n-e, 2 for s-e, 3 for s-w
	dist  int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Corner

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Corner)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Corner, dist int) {
	if item.index < 0 {
		return
	}
	item.dist = dist
	heap.Fix(pq, item.index)
}
