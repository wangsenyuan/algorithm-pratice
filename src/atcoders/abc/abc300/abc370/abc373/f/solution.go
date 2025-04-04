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
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 2)
	}
	return solve(m, items)
}

func solve(m int, items [][]int) int {
	// by size
	pqs := make([]PQ, m+1)
	for _, cur := range items {
		w, v := cur[0], cur[1]
		heap.Push(&pqs[w], v-1)
	}

	dp := make([]int, m+1)
	// dp[i] = 使用前面装箱i个时的最大值

	fp := make([]int, m+1)

	for w := 1; w <= m; w++ {
		if len(pqs[w].IntSlice) == 0 {
			continue
		}

		// 使用j个，重w的item
		for j := 1; j*w <= m; j++ {
			v := heap.Pop(&pqs[w]).(int)
			fp[j] = fp[j-1] + v
			heap.Push(&pqs[w], v-2)
		}

		for j := m; j >= w; j-- {
			for k := 1; k*w <= j; k++ {
				dp[j] = max(dp[j], dp[j-k*w]+fp[k])
			}
		}
	}

	return dp[m]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type PQ struct {
	sort.IntSlice
}

func (pq PQ) Less(i, j int) bool {
	return pq.IntSlice[i] > pq.IntSlice[j]
}

func (pq *PQ) Push(x any) {
	pq.IntSlice = append(pq.IntSlice, x.(int))
}

func (pq *PQ) Pop() any {
	old := pq.IntSlice
	n := len(old)
	x := old[n-1]
	pq.IntSlice = old[0 : n-1]
	return x
}
