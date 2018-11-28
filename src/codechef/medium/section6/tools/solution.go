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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		lines := make([][]int, n)
		for i := 0; i < n; i++ {
			lines[i] = readNNums(scanner, 4)
		}
		fmt.Println(solve(n, lines))
	}
}

const INF = 1 << 30

var MAIN = []int{0, 0}

func solve(n int, lines [][]int) int {
	N := 2 * n
	all := 1 << uint(N)
	dp := make([][]int, all)
	for i := 0; i < all; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<uint(i)][i] = dist(MAIN, lines[i][2:])
	}

	for state := 0; state < all; state++ {
		for j := 0; j < N; j++ {
			if dp[state][j] >= INF {
				continue
			}
			var carry int
			for i := 0; i < n; i++ {
				carry += (state >> uint(i)) & 1
				carry -= (state >> uint(i+n)) & 1
			}
			var pos []int
			if j < n {
				pos = lines[j][2:]
			} else {
				pos = lines[j-n][:2]
			}
			if carry < 2 {
				// can go to another tool
				for i := 0; i < n; i++ {
					if state&(1<<uint(i)) == 0 {
						next := state | (1 << uint(i))
						dp[next][i] = min(dp[next][i], dp[state][j]+dist(pos, lines[i][2:]))
					}
				}
			}
			for i := n; i < N; i++ {
				if state&(1<<uint(i)) == 0 && state&(1<<uint(i-n)) > 0 {
					next := state | (1 << uint(i))
					dp[next][i] = min(dp[next][i], dp[state][j]+dist(pos, lines[i-n][:2]))
				}
			}
		}
	}

	best := INF

	for i := n; i < N; i++ {
		best = min(best, dp[all-1][i]+dist(MAIN, lines[i-n][:2]))
	}

	return best
}

func dist(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(n int, lines [][]int) int {
	chefs := make([][]int, n)
	tools := make([][]int, n)

	for i := 0; i < n; i++ {
		line := lines[i]
		chefs[i] = []int{line[0], line[1]}
		tools[i] = []int{line[2], line[3]}
	}
	all := 1 << uint(n)

	pq := new(PriorityQueue)
	heap.Init(pq)

	//dp[state1][j][state2] = at position j (including chef and tool), state1 chefs are waiting, while carrying state2 tools
	table := make([][][]*Item, all)

	for i := 0; i < all; i++ {
		table[i] = make([][]*Item, 2*n)
		for j := 0; j < 2*n; j++ {
			table[i][j] = make([]*Item, all)
			for k := 0; k < all; k++ {
				item := NewItem(i, j, k, INF)
				table[i][j][k] = item
				heap.Push(pq, item)
			}
		}
	}

	// tool is from 0 to n, and chef from n to 2 * n
	for j := 0; j < n; j++ {
		item := table[all-1][j][1<<uint(j)]
		pq.update(item, dist(tools[j], MAIN))
	}

	best := INF
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		carry := cur.carry
		j := cur.j
		mask := cur.mask
		pri := cur.priority
		if j < n {
			// at tool
			for i := 0; i < n; i++ {
				if carry&(1<<uint(i)) > 0 {
					//go to chef i
					ii := table[mask][i+n][carry]
					if ii.index >= 0 && ii.priority > pri+dist(chefs[i], tools[j]) {
						pq.update(ii, pri+dist(chefs[i], tools[j]))
					}
				}
			}

			if carry^(1<<uint(j)) == 0 {
				//can carry another one
				for i := 0; i < n; i++ {
					if mask&(1<<uint(i)) > 0 && i != j {
						// chef i is awaiting
						iii := table[mask][i][carry|(1<<uint(i))]
						tmp := pri + dist(tools[j], tools[i])
						if iii.index >= 0 && iii.priority > tmp {
							pq.update(iii, tmp)
						}
					}
				}
			}
		} else {
			// at chef
			j -= n
			mask ^= (1 << uint(j))
			if mask == 0 {
				// all served
				best = min(best, pri+dist(chefs[j], MAIN))
				continue
			}
			// tool j is dropped here
			carry ^= (1 << uint(j))
			for i := 0; i < n; i++ {
				if mask&(1<<uint(i)) > 0 && carry&(1<<uint(i)) == 0 {
					//chef i is awaiting, go to tool i
					ii := table[mask][i][carry|(1<<uint(i))]
					tmp := pri + dist(chefs[j], tools[i])
					if ii.index >= 0 && ii.priority > tmp {
						// still in the que
						pq.update(ii, tmp)
					}
				}
			}

			if carry > 0 {
				//can go another chef with the tool
				for i := 0; i < n; i++ {
					if carry&(1<<uint(i)) > 0 {
						ii := table[mask][i+n][carry]
						tmp := pri + dist(chefs[j], chefs[i])
						if ii.index >= 0 && ii.priority > tmp {
							pq.update(ii, tmp)
						}
					}
				}
			}
		}
	}

	return best
}

type Item struct {
	mask     int
	j        int
	carry    int
	priority int
	index    int
}

func NewItem(mask, j, carry, dist int) *Item {
	item := new(Item)
	item.mask = mask
	item.j = j
	item.carry = carry
	item.priority = dist
	item.index = -1
	return item
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
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
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
