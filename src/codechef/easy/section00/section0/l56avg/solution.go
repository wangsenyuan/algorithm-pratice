package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
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

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
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
	t := readOneNum(scanner)

	for i := 0; i < t; i++ {
		n := readOneNum(scanner)
		A := readNNums(scanner, n)
		edges := make([][]int, n-1)
		for j := 0; j < n-1; j++ {
			edges[j] = readNNums(scanner, 2)
		}

		res := solve(n, edges, A)
		fmt.Printf("%.7f\n", res)
	}
}

const oo = 1 << 30

func solve(n int, edges [][]int, A []int) float64 {
	graph := make([]map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a][b] = true
		graph[b][a] = true
	}

	dp := make([][]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]float64, 2)
	}

	var dfs func(v, p int, x float64) bool

	dfs = func(v, p int, x float64) bool {
		dp[v][0] = -oo
		dp[v][1] = -oo

		child := make([]float64, 0, 3)
		rchild := make([]float64, 0, 3)

		for w := range graph[v] {
			if w == p {
				continue
			}
			tmp := dfs(w, v, x)
			if !tmp {
				// fail at children w, no need to process current node any more
				return false
			}
			if dp[w][0] < 0 {
				child = append(child, dp[w][1])
			} else {
				rchild = append(rchild, dp[w][1])
				sort.Sort(sort.Reverse(FloatSlice(rchild)))
				if len(rchild) > 2 {
					rchild = rchild[0:2]
				}
			}
		}

		if len(child) > 2 {
			return false
		}

		if len(child) == 2 {
			// current node has to connect child 0 & 1
			if child[0]+child[1]+float64(A[v-1])-x < 0 {
				return false
			}
			dp[v][0] = child[0] + child[1] + float64(A[v-1]) - x
			dp[v][1] = 0
			return true
		}

		if len(child) == 1 {
			dp[v][0] = float64(A[v-1]) - x + child[0]
			if len(rchild) > 0 && rchild[0] > 0 {
				// connect child 0 & rchild 0 by current
				dp[v][0] += rchild[0]
			}
			dp[v][1] = float64(A[v-1]) - x + child[0]
		}

		if len(child) == 0 {
			if dp[v][1] < float64(A[v-1])-x {
				// just current node
				dp[v][1] = float64(A[v-1]) - x
			}
			if len(rchild) > 0 && float64(A[v-1])-x+rchild[0] > dp[v][1] {
				// connect with the largest sum child
				dp[v][1] = float64(A[v-1]) - x + rchild[0]
			}
			dp[v][0] = float64(A[v-1]) - x
			if len(rchild) > 0 && rchild[0] > 0 {
				dp[v][0] += rchild[0]
			}
			if len(rchild) > 1 && rchild[1] > 0 {
				dp[v][0] += rchild[1]
			}
		}

		if dp[v][0] < dp[v][1] {
			dp[v][0] = dp[v][1]
		}

		if p == 0 {
			return dp[v][0] >= 0
		}
		return true
	}

	check := func(x float64) bool {
		v := rand.Intn(n) + 1
		return dfs(v, 0, x)
	}

	left, right := float64(1e-6), float64(1e6)

	for i := 0; i < 50; i++ {
		mid := (left + right) / 2.0

		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}

	return (left + right) / 2.0
}

type FloatSlice []float64

func (this FloatSlice) Len() int {
	return len(this)
}

func (this FloatSlice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this FloatSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
