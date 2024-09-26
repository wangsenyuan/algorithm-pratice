package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	h, t, R := readThreeNums(reader)
	n := readNum(reader)
	op1 := make([][]int, n)
	for i := 0; i < n; i++ {
		op1[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	op2 := make([][]int, m)
	for i := 0; i < m; i++ {
		op2[i] = readNNums(reader, 2)
	}

	winner, cnt := solve(h, t, R, op1, op2)

	fmt.Println(winner)

	if winner != "Draw" {
		fmt.Println(cnt)
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

const inf = 1 << 60

func solve(head int, tail int, R int, op1 [][]int, op2 [][]int) (string, int) {
	cnt := checkWin(head, tail, R, op1, op2)

	if cnt >= 0 && cnt < inf {
		return "Ivan", cnt
	}

	draw := checkDraw(head, tail, R, op1, op2)
	if draw {
		return "Draw", 0
	}

	cnt = checkLose(head, tail, R, op1, op2)
	return "Zmey", cnt
}

func checkLose(head int, tail int, R int, op1 [][]int, op2 [][]int) int {
	// 这里所有的操作，肯定都是在尽增加的

	dp := make([][]int, R+1)
	for i := 0; i <= R; i++ {
		dp[i] = make([]int, R+1)
		for j := 0; j <= R; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(x int, y int) int
	dfs = func(x int, y int) int {
		if x+y > R {
			return 0
		}
		if dp[x][y] >= 0 {
			return dp[x][y]
		}
		var ans int
		for i, cur := range op1 {
			if i+1 > x {
				break
			}
			nx, ny := x-(i+1)+cur[0], y+cur[1]
			ans = max(ans, 1+dfs(nx, ny))
		}
		for i, cur := range op2 {
			if i+1 > y {
				break
			}
			nx, ny := x+cur[0], y-(i+1)+cur[1]
			ans = max(ans, 1+dfs(nx, ny))
		}

		// ans 肯定是存在的，因为只有(0, 0)没法继续
		dp[x][y] = ans
		return ans
	}

	return dfs(head, tail)
}

func checkWin(head int, tail int, R int, op1 [][]int, op2 [][]int) int {
	X := R + 1
	dp := make([][]*Item, X)
	pq := make(PriorityQueue, X*X)

	for i := 0; i < X; i++ {
		dp[i] = make([]*Item, X)
		for j := 0; j < X; j++ {
			it := new(Item)
			it.value = i*X + j
			it.priority = inf
			it.index = i*X + j
			dp[i][j] = it
			pq[i*X+j] = it
		}
	}

	heap.Init(&pq)
	pq.update(dp[head][tail], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		x, y := it.value/X, it.value%X
		if x == 0 && y == 0 {
			return it.priority
		}

		for i, cur := range op1 {
			if i+1 > x {
				break
			}
			h, t := cur[0], cur[1]
			nh := x - (i + 1) + h
			nt := y + t
			if nh+nt <= R && dp[nh][nt].priority >= it.priority+1 && dp[nh][nt].index >= 0 {
				// ok to proceed
				pq.update(dp[nh][nt], it.priority+1)
			}
		}

		for i, cur := range op2 {
			if i+1 > y {
				break
			}
			h, t := cur[0], cur[1]
			nh := x + h
			nt := y - (i + 1) + t
			if nh+nt <= R && dp[nh][nt].priority >= it.priority+1 && dp[nh][nt].index >= 0 {
				pq.update(dp[nh][nt], it.priority+1)
			}
		}
	}

	return -1
}

func checkDraw(head int, tail int, R int, op1 [][]int, op2 [][]int) bool {
	X := R + 1
	vis := make([][]int, X)
	for i := 0; i < X; i++ {
		vis[i] = make([]int, X)
	}

	var dfs func(x int, y int) bool
	dfs = func(x int, y int) bool {
		// x = 0 and y = 0, 不可能到达
		vis[x][y]++

		for i, cur := range op1 {
			if i+1 > x {
				break
			}
			h, t := cur[0], cur[1]
			nh := x - (i + 1) + h
			nt := y + t
			if nh+nt > R || vis[nh][nt] == 2 {
				continue
			}
			if vis[nh][nt] == 1 || dfs(nh, nt) {
				return true
			}
		}

		for i, cur := range op2 {
			if i+1 > y {
				break
			}
			h, t := cur[0], cur[1]
			nh := x + h
			nt := y - (i + 1) + t
			if nh+nt > R || vis[nh][nt] == 2 {
				continue
			}
			if vis[nh][nt] == 1 || dfs(nh, nt) {
				return true
			}
		}

		vis[x][y]++

		return false
	}

	return dfs(head, tail)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
